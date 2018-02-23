package spec

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/vendasta/mscli/pkg/utils"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const endpointsImageTag = "1.12.0"

// Endpoints sets up endpoints
func Endpoints(environment utils.Environment, host string, endpointsVersion string, secondarySSLConfig *SecondarySSLConfig) v1.Container {
	cpuLimits := resource.MustParse("500m")
	cpuRequests := resource.MustParse("100m")
	memoryLimits := resource.MustParse("256Mi")
	memoryRequests := resource.MustParse("128Mi")

	env := []v1.EnvVar{}
	vm := []v1.VolumeMount{}

	// Configure Volume Mounts
	if environment == utils.Local {
		vm = append(vm,
			v1.VolumeMount{Name: "local-auth", MountPath: "/etc/local-auth/"},
			v1.VolumeMount{Name: "local-app-creds", MountPath: "/etc/local-app-creds/"},
			v1.VolumeMount{Name: "local-app-creds", MountPath: "/etc/nginx/ssl/vendasta-internal-com"},
		)
		env = append(env, v1.EnvVar{
			Name:  "GOOGLE_APPLICATION_CREDENTIALS",
			Value: "/etc/local-auth/application_default_credentials.json",
		})
	} else {
		// always include the vendasta-internal ssl cert
		vm = append(vm, v1.VolumeMount{
			Name:      "vendasta-internal",
			MountPath: "/etc/nginx/ssl/vendasta-internal-com",
			ReadOnly:  true,
		},
			v1.VolumeMount{
				Name:      "endpoints-nginx-conf",
				MountPath: "/etc/nginx/mscli",
				ReadOnly:  true,
			})
	}

	// mount secondary ssl cert if provided
	if secondarySSLConfig != nil && secondarySSLConfig.Name != "" {
		vm = append(vm, v1.VolumeMount{
			Name:      secondarySSLConfig.Name,
			MountPath: fmt.Sprintf("/etc/nginx/ssl/%s", secondarySSLConfig.Name),
			ReadOnly:  true,
		})
	}

	// Endpoints Container Args
	args := []string{
		"-s", host,
		"-v", endpointsVersion,
		"-a", "grpc://127.0.0.1:11000",
		"-p", "11003",
		"-P", "11005",
		"-S", "11006",
		"-z", "healthz",
		"-n", "/etc/nginx/mscli/nginx.conf",
	}
	if environment == utils.Local {
		args = append(args, "-k", "/etc/local-app-creds/key.json")
	}

	probe := &v1.Probe{
		Handler: v1.Handler{
			HTTPGet: &v1.HTTPGetAction{
				Path: "/healthz",
				Port: intstr.FromInt(11003),
			},
		},
	}

	return v1.Container{
		Name:            "endpoints",
		Image:           fmt.Sprintf("gcr.io/endpoints-release/endpoints-runtime:%s", endpointsImageTag),
		Args:            args,
		ImagePullPolicy: v1.PullIfNotPresent,
		LivenessProbe:   probe,
		ReadinessProbe:  probe,
		Ports: []v1.ContainerPort{
			{
				ContainerPort: 11003, //< Serve GRPC + REST
			},
			{
				ContainerPort: 10005, //< Serve Debug info
			},
			{
				ContainerPort: 11006, //< Optionally terminate SSL for GRPC + REST
			},
		},
		Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				v1.ResourceCPU:    cpuLimits,
				v1.ResourceMemory: memoryLimits,
			},
			Requests: v1.ResourceList{
				v1.ResourceCPU:    cpuRequests,
				v1.ResourceMemory: memoryRequests,
			},
		},
		VolumeMounts: vm,
		Env:          env,
	}
}

//EndpointsConfigMap creates an endpoints configmap
func EndpointsConfigMap(namespace string, host string, secondarySSLConfig *SecondarySSLConfig) v1.ConfigMap {
	tmpl, err := template.
		New("nginxConfig").
		Parse(nginxConfigMap)
	if err != nil {
		fmt.Printf("Error parsing nginxConfig template: %s\n", err.Error())
		os.Exit(1)
	}
	buf := bytes.NewBufferString("")
	data := map[string]string{
		"Host": host,
	}
	if secondarySSLConfig != nil && secondarySSLConfig.Name != "" {
		data["SecondaryHost"] = secondarySSLConfig.Host
		data["SecondaryCertPath"] = secondarySSLConfig.Name
	}
	if err = tmpl.Execute(buf, data); err != nil {
		fmt.Printf("Error generating nginxConfigMap: %s\n", err.Error())
		os.Exit(1)
	}

	return v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "endpoints-nginx-conf",
			Namespace: namespace,
		},
		Data: map[string]string{
			"nginx.conf": buf.String(),
		},
	}
}

// Endpoints NGINX configuration override that allows CORS requests.
// The only customization is between the comments # Begin CORS support and # End CORS support
// The rest is pulled from the Endpoints NGINX configuration.
// For more info, see: https://groups.google.com/forum/#!searchin/google-cloud-endpoints/cors%7Csort:relevance/google-cloud-endpoints/THvCfetfzW8/luyH3tOUBgAJ
const nginxConfigMap = `
# Auto-generated
daemon off;

user nginx nginx;

pid /var/run/nginx.pid;

worker_processes 1;

# Logging to stderr enables better integration with Docker and GKE/Kubernetes.
error_log stderr warn;

events { worker_connections 4096; }

http {
  include /etc/nginx/mime.types;
  server_tokens off;
  client_max_body_size 32m;
  client_body_buffer_size 128k;
  server_names_hash_bucket_size 128;

  # HTTP subrequests
  endpoints_resolver 8.8.8.8;
  endpoints_certificates /etc/nginx/trusted-ca-certificates.crt;

  server {
    server_name "{{.Host}}";

    listen 11003;
    listen 11005 http2;
    listen 11006 ssl http2;
    ssl_certificate /etc/nginx/ssl/vendasta-internal-com/nginx.crt;
    ssl_certificate_key /etc/nginx/ssl/vendasta-internal-com/nginx.key;

    access_log /dev/stdout;

    location = /healthz {
      return 200;
      access_log off;
    }

    location / {
      # Begin CORS support
        if ($request_method = 'OPTIONS') {
          # Tell client that this pre-flight info is valid for 20 days
          add_header 'Access-Control-Max-Age' 1728000;
          add_header 'Content-Type' 'text/plain charset=UTF-8';
          add_header 'Content-Length' 0;
          add_header 'Access-Control-Allow-Origin' $http_origin;
          add_header 'Access-Control-Allow-Credentials' 'true';
          add_header 'Access-Control-Allow-Methods' 'POST';
          add_header 'Access-Control-Allow-Headers' 'Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With';
          return 204;
	    }

        add_header 'Access-Control-Allow-Origin' $http_origin always;
        add_header 'Access-Control-Allow-Methods' 'POST' always;
        add_header 'Access-Control-Allow-Headers' 'Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With' always;
 		add_header 'Access-Control-Allow-Credentials' 'true' always;
      # End CORS support
      # Begin Endpoints v2 Support
      endpoints {
        on;
        server_config /etc/nginx/server_config.pb.txt;
        api /etc/nginx/endpoints/service.json;
        metadata_server http://169.254.169.254;
      }
      # End Endpoints v2 Support

      # WARNING: only first backend is used
      grpc_pass 127.0.0.1:11000 override;
    }

    include /var/lib/nginx/extra/*.conf;
  }
{{if (and (.SecondaryHost) (.SecondaryCertPath))}}
  server {
    server_name "{{.SecondaryHost}}";

    listen 11003;
    listen 11005 http2;
    listen 11006 ssl http2;

    ssl_certificate /etc/nginx/ssl/{{.SecondaryCertPath}}/nginx.crt;
    ssl_certificate_key /etc/nginx/ssl/{{.SecondaryCertPath}}/nginx.key;

    access_log /dev/stdout;

    location = /healthz {
      return 200;
      access_log off;
    }

    location / {
      # Begin CORS support
        if ($request_method = 'OPTIONS') {
          # Tell client that this pre-flight info is valid for 20 days
          add_header 'Access-Control-Max-Age' 1728000;
          add_header 'Content-Type' 'text/plain charset=UTF-8';
          add_header 'Content-Length' 0;
          add_header 'Access-Control-Allow-Origin' $http_origin;
          add_header 'Access-Control-Allow-Credentials' 'true';
          add_header 'Access-Control-Allow-Methods' 'POST';
          add_header 'Access-Control-Allow-Headers' 'Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With';
          return 204;
	    }

        add_header 'Access-Control-Allow-Origin' $http_origin always;
        add_header 'Access-Control-Allow-Methods' 'POST' always;
        add_header 'Access-Control-Allow-Headers' 'Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With' always;
 		add_header 'Access-Control-Allow-Credentials' 'true' always;
      # End CORS support
      # Begin Endpoints v2 Support
      endpoints {
        on;
        server_config /etc/nginx/server_config.pb.txt;
        api /etc/nginx/endpoints/service.json;
        metadata_server http://169.254.169.254;
      }
      # End Endpoints v2 Support

      # WARNING: only first backend is used
      grpc_pass 127.0.0.1:11000 override;
    }

    include /var/lib/nginx/extra/*.conf;
  }
{{end}}
  server {
    # expose /nginx_status and /endpoints_status but on a different port to
    # avoid external visibility / conflicts with the app.
    listen 8090;
    location /nginx_status {
      stub_status on;
      access_log off;
    }
    location /endpoints_status {
      endpoints_status;
      access_log off;
    }
    location /healthz {
      return 200;
      access_log off;
    }
    location / {
      root /dev/null;
    }
  }
}`
