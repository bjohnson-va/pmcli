package cssdk

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxonomy_GetChild(t *testing.T) {
	raw := `{
      "legacyCategoryId": "CARE",
      "taxId": "beautysvc",
      "children": [
        {
          "taxId": "beautysvc:barbers",
          "name": "Barbers"
        },
        {
          "taxId": "beautysvc:cosmetics",
          "name": "Cosmetics & Beauty Supply"
        },
        {
          "taxId": "beautysvc:spas",
          "name": "Day Spas"
        },
        {
          "taxId": "beautysvc:eyelashservice",
          "name": "Eyelash Service"
        },
        {
          "taxId": "beautysvc:hair_extensions",
          "name": "Hair Extensions"
        },
        {
          "taxId": "beautysvc:hairremoval",
          "children": [
            {
              "taxId": "beautysvc:hairremoval:laser_hair_removal",
              "name": "Laser Hair Removal"
            }
          ],
          "name": "Hair Removal"
        },
        {
          "taxId": "beautysvc:hair",
          "children": [
            {
              "taxId": "beautysvc:hair:blowoutservices",
              "name": "Blow Dry/Out Services"
            },
            {
              "taxId": "beautysvc:hair:hair_extensions",
              "name": "Hair Extensions"
            },
            {
              "taxId": "beautysvc:hair:hairstylists",
              "name": "Hair Stylists"
            },
            {
              "taxId": "beautysvc:hair:menshair",
              "name": "Men's Hair Salons"
            }
          ],
          "name": "Hair Salons"
        },
        {
          "taxId": "beautysvc:makeupartists",
          "name": "Makeup Artists"
        },
        {
          "taxId": "beautysvc:massage",
          "name": "Massage"
        },
        {
          "taxId": "beautysvc:medicalspa",
          "name": "Medical Spas"
        },
        {
          "taxId": "beautysvc:othersalons",
          "name": "Nail Salons"
        },
        {
          "taxId": "beautysvc:permanentmakeup",
          "name": "Permanent Makeup"
        },
        {
          "taxId": "beautysvc:piercing",
          "name": "Piercing"
        },
        {
          "taxId": "beautysvc:rolfing",
          "name": "Rolfing"
        },
        {
          "taxId": "beautysvc:skincare",
          "name": "Skin Care"
        },
        {
          "taxId": "beautysvc:tanning",
          "children": [
            {
              "taxId": "beautysvc:tanning:spraytanning",
              "name": "Spray Tanning"
            },
            {
              "taxId": "beautysvc:tanning:tanningbeds",
              "name": "Tanning Beds"
            }
          ],
          "name": "Tanning"
        },
        {
          "taxId": "beautysvc:tattoo",
          "name": "Tattoo"
        }
      ],
      "name": "Beauty & Spas"
    }`

	tax := &Taxonomy{}
	json.Unmarshal([]byte(raw), tax)

	child := tax.GetChild("beautysvc:barbers")
	assert.Equal(t, child.ID, "beautysvc:barbers", "First level child")

	child = tax.GetChild("beautysvc:tanning:spraytanning")
	assert.Equal(t, child.ID, "beautysvc:tanning:spraytanning", "Second level child")

	child = tax.GetChild("beautysvc:tanning:potatoes")
	assert.Nil(t, child, "Child not in subtree.")
}
