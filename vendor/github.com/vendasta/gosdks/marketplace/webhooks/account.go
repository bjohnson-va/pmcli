package marketplacewebhooks

import (
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
)

// MarketplaceAccount is NAP data for an account from marketplace
type MarketplaceAccount struct {
	Deleted          bool     `json:"deleted_flag"`
	ContactFirstName string   `json:"contact_first_name"`
	CommonName       []string `json:"common_name"`
	WorkNumber       []string `json:"work_number"`
	PartnerID        string   `json:"partner_id"`
	AccountGroupID   string   `json:"id"`
	City             string   `json:"city"`
	Zip              string   `json:"zip"`
	MarketID         string   `json:"market_id"`
	State            string   `json:"state"`
	CompanyName      string   `json:"company_name"`
	Latitude         float64  `json:"latitude"`
	Longitude        float64  `json:"longitude"`
	FoursquareURL    string   `json:"foursquare_url"`
	TaxonomyIDs      []string `json:"tax_ids"`
	Website          string   `json:"website"`
	RssURL           string   `json:"rss_url"`
	Updated          string   `json:"updated"`
	TwitterURL       string   `json:"twitter_url"`
	FacebookURL      string   `json:"facebook_url"`
	FaxNumber        string   `json:"fax_number"`
	ContactEmail     string   `json:"contact_email"`
	Address          string   `json:"address"`
	YoutubeURL       string   `json:"youtube_url"`
	InstagramURL     string   `json:"instagram_url"`
	KeyPerson        []string `json:"key_person"`
	PinterestURL     string   `json:"pinterest_url"`
	ContactLastName  string   `json:"contact_last_name"`
	CellNumber       string   `json:"cell_number"`
	Country          string   `json:"country"`
	Created          string   `json:"created"`
	LinkedInURL      string   `json:"linkedin_url"`
	SalespersonID    string   `json:"sales_person_id"`
	GooglePlusURL    string   `json:"googleplus_url"`
}

// AccountWebhookPayload represents the payload of an account webhook
type AccountWebhookPayload struct {
	Account   MarketplaceAccount `json:"account"`
	MarketID  string             `json:"market_id"`
	WebhookID string             `json:"webhook_id"`
	Action    string             `json:"action"`
	PartnerID string             `json:"partner_id"`
}

type accountWebhookClaims struct {
	jwt.StandardClaims
	WebhookPayload AccountWebhookPayload `json:"vendasta.com/marketplace/webhook"`
}

// handles an account webhook by wrapping the passed handler
func (s *server) handleAccountWebhook(r *http.Request, handler accountWebhookHandler) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logging.Warningf(s.ctx, "Error reading request body: %s", err.Error())
		return util.Error(util.Internal, "Error reading request body: %s", err.Error())
	}

	payload, err := s.parseAccountPayload(string(reqBody))
	logging.Infof(s.ctx, "Account webhook payload is %+v", payload)
	if err != nil {
		logging.Errorf(s.ctx, "Jwt auth failed: %s. Payload body: %s", err.Error(), string(reqBody))
		return util.Error(util.Unauthenticated, "Jwt auth failed: %s", err.Error())
	}

	err = handler(payload, r.URL)
	if err != nil {
		logging.Warningf(s.ctx, "Error handling account webhook payload: %s", err.Error())
		return util.Error(util.Internal, err.Error())
	}

	return nil
}

func (s *server) parseAccountPayload(jwtPayload string) (*AccountWebhookPayload, error) {
	jwtClaims := &accountWebhookClaims{}
	_, err := jwt.ParseWithClaims(jwtPayload, jwtClaims,
		func(tok *jwt.Token) (interface{}, error) {
			return s.marketplacePublicKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	payload := &jwtClaims.WebhookPayload
	return payload, nil
}
