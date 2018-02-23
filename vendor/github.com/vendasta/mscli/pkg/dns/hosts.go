package dns

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/vendasta/mscli/pkg/dns/endpoint"
	"github.com/vendasta/mscli/pkg/dns/provider"
	mscliio "github.com/vendasta/mscli/pkg/io"
	"github.com/vendasta/mscli/pkg/k8s"
	"github.com/vendasta/mscli/pkg/kubernetes"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	VendastaInternalDomain = "vendasta-internal.com"
	WhiteLabeledGRPCDomain = "apigateway.co"
)

//HostsTable Represents all entries in /etc/hosts
type HostsTable struct {
	Entries []HostsEntry
}

//HostsEntry Represents an entry in /etc/hosts
type HostsEntry struct {
	IPAddress     string
	CanonicalName string
	Alias         []string
	Comment       string
}

func pop(str string) (string, string) {
	re := regexp.MustCompile(`^\s*(\S+)(.*)`)
	match := re.FindStringSubmatch(str)
	if match == nil {
		return "", ""
	}
	return match[1], strings.Trim(match[2], " \t")
}

func parseEntry(line string) HostsEntry {
	rv := HostsEntry{}
	//Strip off the comment
	tokens := strings.Split(line, "#")
	if len(tokens) > 1 {
		rv.Comment = "#" + strings.Join(tokens[1:], "#")
	}
	data := strings.Trim(tokens[0], " \t")

	if data != "" {
		rv.IPAddress, data = pop(data)
	}
	if data != "" {
		rv.CanonicalName, data = pop(data)
	}
	var alias string
	for rv.Alias = []string{}; data != ""; {
		alias, data = pop(data)
		rv.Alias = append(rv.Alias, alias)
	}
	if len(rv.Alias) == 0 {
		rv.Alias = nil
	}

	return rv
}

//String converts an entry to a string (how it's represented in /etc/hosts)
func (m HostsEntry) String() string {
	var rv string
	if m.IPAddress == "" {
		if m.Comment == "" {
			return ""
		}
		return m.Comment
	}

	if m.CanonicalName == "" {
		rv = m.IPAddress
	} else if m.Alias == nil {
		rv = fmt.Sprintf("%s %s", m.IPAddress, m.CanonicalName)
	} else {
		rv = fmt.Sprintf("%s %s %s", m.IPAddress, m.CanonicalName, strings.Join(m.Alias, " "))
	}
	if m.Comment == "" {
		return rv
	} else if rv == "" {
		return m.Comment
	}
	return fmt.Sprintf("%s %s", rv, m.Comment)
}

//NewHostsTableFromFile parses a file (usually /etc/hosts)
func NewHostsTableFromFile(fileName string) (*HostsTable, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return NewHostsTable(f)
}

//NewHostsTable parses a stream (easier to test than a file)
func NewHostsTable(r io.Reader) (*HostsTable, error) {
	var err error
	var line string
	br := bufio.NewReader(r)
	rv := HostsTable{Entries: []HostsEntry{}}
	for {
		line, err = br.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err == nil || err == io.EOF {
			rv.Entries = append(rv.Entries, parseEntry(line))
			if err == io.EOF {
				break
			}
		} else {
			return nil, err
		}
	}
	return &rv, nil
}

//Serialize converts a full set of HostsEntries into the /etc/hosts format
func (m *HostsTable) Serialize(w io.Writer) error {
	if m == nil {
		return nil
	}
	for _, e := range m.Entries {
		line := e.String()
		if _, err := w.Write([]byte(line + "\n")); err != nil {
			return err
		}
	}
	return nil
}

//GetEntryByHost returns a host entry that corresponds to a hostname (nil if not found)
func (m *HostsTable) GetEntryByHost(hostName string) *HostsEntry {
	if m == nil {
		return nil
	}
	for _, e := range m.Entries {
		if e.CanonicalName == hostName {
			return &e
		}
	}
	return nil
}

func verifyARecord(hostname, ipAddress string) (bool, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return false, fmt.Errorf("error performing DNS lookup on hostname %s: %s", hostname, err.Error())
	} else if len(ips) > 1 {
		msg := fmt.Sprintf("You have more than a single A-record setup for hostname %s: ", hostname)
		for _, ip := range ips {
			msg += fmt.Sprintf(" - %s\n", ip.String())
		}
		return false, errors.New(msg)
	}

	ip := ips[0]
	if ip.String() != ipAddress {
		return false, fmt.Errorf("your A-record for %s points to %s, but the LoadBalancerIP is set to %s",
			hostname, ip.String(), ipAddress)
	}
	return true, nil
}

//TestDNS handles the command "test-dns" from the commandline
func TestDNS(specFile spec.MicroserviceFile, env utils.Environment) error {
	se, err := specFile.Microservice.GetEnv(env)
	if err != nil {
		return err
	}
	if env == utils.Local {
		ht, err := NewHostsTableFromFile("/etc/hosts")
		if err != nil {
			return fmt.Errorf("error reading /etc/hosts file: %s", err.Error())
		}
		e := ht.GetEntryByHost(se.GRPCHost)
		if e == nil {
			return fmt.Errorf("no host entry found for %s", se.GRPCHost)
		}
		if e.IPAddress == "" {
			fmt.Printf("Your /etc/hosts file does not contain an entry like the following:\n")
			he := HostsEntry{
				IPAddress:     "127.0.0.1",
				CanonicalName: se.GRPCHost,
				Comment:       fmt.Sprintf("#Static NS Lookup for mscli service %s", specFile.Microservice.Name)}
			fmt.Printf("%s\n", he.String())
		}

		fmt.Printf("The entry you have may appears correct:\n%s\n\n", e.String())
		//Get the minikube IP
		ip, err := mscliio.RunCommand("minikube", "ip").Execute()
		if err != nil {
			fmt.Printf("   minikube appears to not be running or installed (%s), so IP can't be verified\n", err.Error())
		} else {
			if e.IPAddress != ip {
				fmt.Printf("   but the IP do not match minikube's IP (%s)\n", ip)
			} else {
				fmt.Printf("The entry you have looks correct (%s)\n", e.String())
				fmt.Printf("   the IP matches the minikube IP (%s)\n", ip)
			}
		}
	} else {
		// Ensure GRPC
		grpcLoadBalancerIP := se.Network.GetGRPCLoadBalancerIP()
		ok, err := verifyARecord(se.GRPCHost, grpcLoadBalancerIP)
		if ok {
			fmt.Printf("A-record set correctly %s -> %s\n", se.GRPCHost, grpcLoadBalancerIP)
		} else {
			fmt.Printf("A-record set incorrectly: %s\n", err.Error())
			fmt.Printf("Expected it to be set to: %s -> %s\n", se.GRPCHost, grpcLoadBalancerIP)
		}
		if se.SecondarySSLConfig != nil && se.SecondarySSLConfig.Host != "" {
			// Ensure Whitelabelled GRPC
			ok, err = verifyARecord(se.GRPCHost, grpcLoadBalancerIP)
			if ok {
				fmt.Printf("A-record set correctly %s -> %s\n", se.SecondarySSLConfig.Host, grpcLoadBalancerIP)
			} else {
				fmt.Printf("A-record set incorrectly: %s\n", err.Error())
				fmt.Printf("Expected it to be set to: %s -> %s\n", se.SecondarySSLConfig.Host, grpcLoadBalancerIP)
			}
		}

		// Ensure HTTPS
		ok, err = verifyARecord(se.HTTPSHost, se.Network.HTTPSLoadBalancerIP)
		if ok {
			fmt.Printf("A-record set correctly %s -> %s\n", se.HTTPSHost, se.Network.HTTPSLoadBalancerIP)
		} else {
			fmt.Printf("A-record set incorrectly: %s\n", err.Error())
			fmt.Printf("Expected it to be set to: %s -> %s\n", se.HTTPSHost, se.Network.HTTPSLoadBalancerIP)
		}
	}
	return nil
}

// ConfigureDNS configures the external IP's and DNS records for the microservice's services (load balancers)
// Returns grpcIP, httpIP, error
func ConfigureDNS(specFile spec.MicroserviceFile, env utils.Environment) (string, string, error) {
	k8sAPI, err := kubernetes.GetK8sClientSet(specFile, env)
	if err != nil {
		return "", "", err
	}
	se, err := specFile.Microservice.GetEnv(env)
	if err != nil {
		return "", "", err
	}
	objects, err := specFile.Microservice.K8S(env, "", nil)
	if err != nil {
		return "", "", err
	}
	for i := 0; i < len(objects); i++ {
		switch objects[i].(type) {
		case k8s.Service:
			service := objects[i].(k8s.Service).Service()
			s, err := k8sAPI.CoreV1().Services(service.Namespace).Get(service.Name, metav1.GetOptions{})
			if err != nil {
				if !k8serrors.IsNotFound(err) {
					return "", "", fmt.Errorf("error getting service: %s", err.Error())
				}
				// Does not exist
			}
			ip := s.Status.LoadBalancer.Ingress[0].IP
			fmt.Printf("External IP for %s: %s\n", service.Name, ip)

			ReserveStaticIP(fmt.Sprintf("%s-%s", service.Name, se.Name), ip)
			if strings.HasSuffix(service.Name, "grpc-svc") {
				CreateDNSRecord(se.GRPCHost, ip, VendastaInternalDomain)

				if se.SecondarySSLConfig != nil && se.SecondarySSLConfig.Host != "" {
					CreateDNSRecord(se.SecondarySSLConfig.Host, ip, WhiteLabeledGRPCDomain)
				}
				return ip, "", nil
			} else {
				CreateDNSRecord(se.HTTPSHost, ip, VendastaInternalDomain)
				return "", ip, nil
			}
		}
	}
	return "", "", nil
}

// ReserveStaticIP upgrades the IP assigned to the service from ephemeral to static
func ReserveStaticIP(name string, ip string) {
	log.Println("Upgrading the IP from ephemeral to static (if necessary)...")
	cmd := exec.Command("gcloud", "compute", "addresses", "create", name, "--addresses", ip, "--region", "us-central1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

// CreateDNSRecord creates the A record to point to the service IP
func CreateDNSRecord(name string, ip string, domain string) {
	p, err := provider.NewGoogleProvider("repcore-prod", provider.NewDomainFilter([]string{domain}), false)
	if err != nil {
		log.Fatalf("error creating provider: %s", err)
	}
	e := endpoint.NewEndpoint(name, ip, "A")
	p.CreateRecords([]*endpoint.Endpoint{e})
	if err != nil {
		log.Fatalf("error creating DNS record: %s", err)
	}
}
