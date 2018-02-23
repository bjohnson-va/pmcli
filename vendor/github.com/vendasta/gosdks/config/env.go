package config

import (
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/oauth2/v2"
)

// Getenvironment returns the environment for this instance
func Getenvironment() string {
	return os.Getenv("ENVIRONMENT")
}

// GetGkeNamespace returns the GKE Namespace for this instance
func GetGkeNamespace() string {
	return os.Getenv("GKE_NAMESPACE")
}

//GetGkePodName reutrns the GKE PodName for this instance
func GetGkePodName() string {
	return os.Getenv("GKE_NAMESPACE")
}

// IsLocal returns true if this instance is running locally
func IsLocal() bool {
	return Getenvironment() == "local" || Getenvironment() == ""
}

// IsProd returns true if this instance is running on production
func IsProd() bool {
	return Getenvironment() == "production"
}

// cache email
var email = ""

// GetCurrentLocalUserName returns the local user's name based on their google oauth email
func GetCurrentLocalUserName() string {
	if email != "" {
		return email
	}
	client, err := google.DefaultClient(context.Background(), oauth2.UserinfoEmailScope)
	service, err := oauth2.New(client)
	if err != nil {
		panic(err)
	}
	us := oauth2.NewUserinfoService(service)
	ui, err := us.Get().Do()
	if err != nil {
		panic(err)
	}
	email = ui.Email
	return email
}

// GetServiceAccount returns the service account to use for this instance
func GetServiceAccount() string {
	if IsLocal() {
		return GetCurrentLocalUserName()
	}
	return os.Getenv("SERVICE_ACCOUNT")
}
