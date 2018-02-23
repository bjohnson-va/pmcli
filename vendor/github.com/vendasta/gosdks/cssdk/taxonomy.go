package cssdk

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

// Taxonomy is the definition for a business category as it's understood by Core Services
type Taxonomy struct {
	// Canonical Identifier (the one you should be using and storing)
	ID string `json:"taxId"`

	// Display Name (for UIs)
	Name string `json:"name"`

	// Legacy identifier from Rep Intel. Use `ID` instead if possible.
	LegacyID string `json:"legacyCategoryId"`

	// Child taxonomies.
	Children []*Taxonomy `json:"children"`
}

func (t *Taxonomy) isLeaf() bool {
	return len(t.Children) == 0
}

func (t *Taxonomy) Equals(ID string) bool {
	return ID == t.ID
}

func (t *Taxonomy) IsChild(ID string) bool {
	if !strings.HasPrefix(ID, t.ID) {
		return false
	}
	if ID == t.ID {
		return false
	}
	return true
}

func (t *Taxonomy) GetChild(ID string) *Taxonomy {
	if !strings.HasPrefix(ID, t.ID) {
		return nil
	}
	if t.ID == ID {
		return t
	}
	if t.isLeaf() {
		return nil
	}
	for _, child := range t.Children {
		c := child.GetChild(ID)
		if c != nil {
			return c
		}
	}
	return nil
}

// taxonomiesFromResponse converts an http response from core services to a list of taxonomies
func taxonomiesFromResponse(r io.Reader) ([]*Taxonomy, error) {
	type Response struct {
		Taxonomies []*Taxonomy `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		reason := "Failed to convert response to Taxonomies: " + err.Error()
		return nil, errors.New(reason)
	}

	return res.Taxonomies, nil
}

type TaxonomySet []*Taxonomy

func (ts TaxonomySet) IsValid(taxID string) bool {
	for _, t := range ts {
		if t.Equals(taxID) {
			return true
		} else if t.IsChild(taxID) {
			return TaxonomySet(t.Children).IsValid(taxID)
		}
	}
	return false
}
