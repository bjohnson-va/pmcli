package domain

import (
	"net/url"

	"github.com/vendasta/gosdks/util"
)

// ReplaceDomain attempts to replace the in URL's domain with the new domain.
func ReplaceDomain(originalURL string, newDomain *Domain) (string, error) {
	if newDomain == nil {
		return "", util.Error(util.InvalidArgument, "invalid whitelabel name")
	}

	u, err := url.Parse(originalURL)
	if err != nil {
		return "", err
	}
	u.Host = newDomain.Name
	u.Scheme = "http"
	if newDomain.Secure {
		u.Scheme = "https"
	}

	// This URL will be going inside of a query param, so doing this prevents double URL encoding
	res, err := url.PathUnescape(u.String())
	if err != nil {
		return "", err
	}

	return res, nil
}
