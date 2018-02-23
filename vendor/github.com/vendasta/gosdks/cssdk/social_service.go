package cssdk

import (
	"encoding/json"
	"errors"
	"net/http"
)

//ServiceType is the type of service (facebook user, facebook page, twitter user, etc)
type ServiceType string

const (
	//ServiceTypeTwitterUser indicates a twitter user service
	ServiceTypeTwitterUser = ServiceType("TW_USER")
	//ServiceTypeFacebookUser indicates a facebook user service
	ServiceTypeFacebookUser = ServiceType("FB_USER")
	//ServiceTypeFacebookPage indicates a facebook page service
	ServiceTypeFacebookPage = ServiceType("FB_PAGE")
	//ServiceTypeGooglePlusUser indicates a google plus user service
	ServiceTypeGooglePlusUser = ServiceType("GP_USER")
	//ServiceTypeGooglePlusPage indicates a google plus page service
	ServiceTypeGooglePlusPage = ServiceType("GP_PAGE")
	//ServiceTypeLinkedinUser indicates a linkedin service
	ServiceTypeLinkedinUser = ServiceType("LI_USER")
	//ServiceTypeLinkedinCompany indicates a linkedin company page service
	ServiceTypeLinkedinCompany = ServiceType("LI_COMPANY")
)

//SocialService represents a cs social service
type SocialService struct {
	SSID            string `json:"ssid"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	AccountGroupID  string `json:"agid"`
	IsAuthenticated bool   `json:"isAuthenticated"`
	ProfileURL      string `json:"profileUrl"`
	ProfileImageURL string `json:"profileImageUrl"`
	ServiceType     string `json:"serviceType"`
	SocialProfileID string `json:"spid"`
}

type SocialServicesByType struct {
	TwitterUsers      []*SocialService `json:"TW_USER"`
	FacebookUsers     []*SocialService `json:"FB_USER"`
	FacebookPages     []*SocialService `json:"FB_PAGE"`
	GooglePlusUsers   []*SocialService `json:"GP_USER"`
	GooglePlusPages   []*SocialService `json:"GP_PAGE"`
	LinkedinUsers     []*SocialService `json:"LI_USER"`
	LinkedinCompanies []*SocialService `json:"LI_COMPANY"`
}

// socialServiceFromResponse converts an http response from core services to a list of social services
func socialServicesFromResponse(r *http.Response) ([]*SocialService, error) {
	defer r.Body.Close()
	type Response struct {
		SocialServicesByType *SocialServicesByType `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to SocialServices: " + err.Error()
		return nil, errors.New(reason)
	}
	var socialServices []*SocialService

	socialServices = append(socialServices, res.SocialServicesByType.TwitterUsers...)
	socialServices = append(socialServices, res.SocialServicesByType.FacebookUsers...)
	socialServices = append(socialServices, res.SocialServicesByType.FacebookPages...)
	socialServices = append(socialServices, res.SocialServicesByType.GooglePlusUsers...)
	socialServices = append(socialServices, res.SocialServicesByType.GooglePlusPages...)
	socialServices = append(socialServices, res.SocialServicesByType.LinkedinUsers...)
	socialServices = append(socialServices, res.SocialServicesByType.LinkedinCompanies...)

	return socialServices, nil
}

//SocialServiceIDToServiceType converts a social service ID to it's corresponding ServiceType based on the first 3 characters of the id
func SocialServiceIDToServiceType(ssid string) ServiceType {
	switch ssid[:3] {
	case "TWU":
		return ServiceTypeTwitterUser
	case "FBP":
		return ServiceTypeFacebookPage
	case "FBU":
		return ServiceTypeFacebookUser
	case "LIU":
		return ServiceTypeLinkedinUser
	case "LIC":
		return ServiceTypeLinkedinCompany
	case "GPP":
		return ServiceTypeGooglePlusPage
	case "GPU":
		return ServiceTypeGooglePlusUser
	default:
		return ""
	}
}
