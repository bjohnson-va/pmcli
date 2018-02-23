package cssdk

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

var taxonomyResponse = `{
  "version": "1.0",
  "data": [
    {
      "legacyCategoryId": "REC",
      "taxId": "active",
      "children": [
        {
          "taxId": "active:amateursportsteams",
          "name": "Amateur Sports Teams"
        },
        {
          "taxId": "active:amusementparks",
          "name": "Amusement Parks"
        },
        {
          "taxId": "active:aquariums",
          "name": "Aquariums"
        },
        {
          "taxId": "active:archery",
          "name": "Archery"
        },
        {
          "taxId": "active:badminton",
          "name": "Badminton"
        },
        {
          "taxId": "active:basketballcourts",
          "name": "Basketball Courts"
        },
        {
          "taxId": "active:beaches",
          "name": "Beaches"
        },
        {
          "taxId": "active:bikerentals",
          "name": "Bike Rentals"
        },
        {
          "taxId": "active:boating",
          "name": "Boating"
        },
        {
          "taxId": "active:bowling",
          "name": "Bowling"
        },
        {
          "taxId": "active:climbing",
          "name": "Climbing"
        },
        {
          "taxId": "active:discgolf",
          "name": "Disc Golf"
        },
        {
          "taxId": "active:diving",
          "children": [
            {
              "taxId": "active:diving:freediving",
              "name": "Free Diving"
            },
            {
              "taxId": "active:diving:scuba",
              "name": "Scuba Diving"
            }
          ],
          "name": "Diving"
        },
        {
          "taxId": "active:fishing",
          "name": "Fishing"
        },
        {
          "taxId": "active:fitness",
          "children": [
            {
              "taxId": "active:fitness:barreclasses",
              "name": "Barre Classes"
            },
            {
              "taxId": "active:fitness:bootcamps",
              "name": "Boot Camps"
            },
            {
              "taxId": "active:fitness:boxing",
              "name": "Boxing"
            },
            {
              "taxId": "active:fitness:dancestudio",
              "name": "Dance Studios"
            },
            {
              "taxId": "active:fitness:gyms",
              "name": "Gyms"
            },
            {
              "taxId": "active:fitness:martialarts",
              "name": "Martial Arts"
            },
            {
              "taxId": "active:fitness:pilates",
              "name": "Pilates"
            },
            {
              "taxId": "active:fitness:swimminglessons",
              "name": "Swimming Lessons/Schools"
            },
            {
              "taxId": "active:fitness:taichi",
              "name": "Tai Chi"
            },
            {
              "taxId": "active:fitness:healthtrainers",
              "name": "Trainers"
            },
            {
              "taxId": "active:fitness:yoga",
              "name": "Yoga"
            }
          ],
          "name": "Fitness & Instruction"
        },
        {
          "taxId": "active:gokarts",
          "name": "Go Karts"
        },
        {
          "taxId": "active:golf",
          "name": "Golf"
        },
        {
          "taxId": "active:gun_ranges",
          "name": "Gun/Rifle Ranges"
        },
        {
          "taxId": "active:gymnastics",
          "name": "Gymnastics"
        },
        {
          "taxId": "active:hanggliding",
          "name": "Hang Gliding"
        },
        {
          "taxId": "active:hiking",
          "name": "Hiking"
        },
        {
          "taxId": "active:horseracing",
          "name": "Horse Racing"
        },
        {
          "taxId": "active:horsebackriding",
          "name": "Horseback Riding"
        },
        {
          "taxId": "active:hot_air_balloons",
          "name": "Hot Air Balloons"
        },
        {
          "taxId": "active:kiteboarding",
          "name": "Kiteboarding"
        },
        {
          "taxId": "active:lakes",
          "name": "Lakes"
        },
        {
          "taxId": "active:lasertag",
          "name": "Laser Tag"
        },
        {
          "taxId": "active:leisure_centers",
          "name": "Leisure Centers"
        },
        {
          "taxId": "active:mini_golf",
          "name": "Mini Golf"
        },
        {
          "taxId": "active:mountainbiking",
          "name": "Mountain Biking"
        },
        {
          "taxId": "active:paddleboarding",
          "name": "Paddleboarding"
        },
        {
          "taxId": "active:paintball",
          "name": "Paintball"
        },
        {
          "taxId": "active:parks",
          "children": [
            {
              "taxId": "active:parks:dog_parks",
              "name": "Dog Parks"
            },
            {
              "taxId": "active:parks:skate_parks",
              "name": "Skate Parks"
            }
          ],
          "name": "Parks"
        },
        {
          "taxId": "active:playgrounds",
          "name": "Playgrounds"
        },
        {
          "taxId": "active:rafting",
          "name": "Rafting/Kayaking"
        },
        {
          "taxId": "active:recreation",
          "name": "Recreation Centers"
        },
        {
          "taxId": "active:rock_climbing",
          "name": "Rock Climbing"
        },
        {
          "taxId": "active:skatingrinks",
          "name": "Skating Rinks"
        },
        {
          "taxId": "active:skydiving",
          "name": "Skydiving"
        },
        {
          "taxId": "active:football",
          "name": "Soccer"
        },
        {
          "taxId": "active:spinclasses",
          "name": "Spin Classes"
        },
        {
          "taxId": "active:sports_clubs",
          "name": "Sports Clubs"
        },
        {
          "taxId": "active:squash",
          "name": "Squash"
        },
        {
          "taxId": "active:summer_camps",
          "name": "Summer Camps"
        },
        {
          "taxId": "active:surfing",
          "name": "Surfing"
        },
        {
          "taxId": "active:swimmingpools",
          "name": "Swimming Pools"
        },
        {
          "taxId": "active:tennis",
          "name": "Tennis"
        },
        {
          "taxId": "active:trampoline",
          "name": "Trampoline Parks"
        },
        {
          "taxId": "active:tubing",
          "name": "Tubing"
        },
        {
          "taxId": "active:zoos",
          "name": "Zoos"
        }
      ],
      "name": "Active Life"
    },
    {
      "legacyCategoryId": "ARTS",
      "taxId": "arts",
      "children": [
        {
          "taxId": "arts:arcades",
          "name": "Arcades"
        },
        {
          "taxId": "arts:galleries",
          "name": "Art Galleries"
        },
        {
          "taxId": "arts:gardens",
          "name": "Botanical Gardens"
        },
        {
          "taxId": "arts:casinos",
          "name": "Casinos"
        },
        {
          "taxId": "arts:movietheaters",
          "name": "Cinema"
        },
        {
          "taxId": "arts:culturalcenter",
          "name": "Cultural Center"
        },
        {
          "taxId": "arts:festivals",
          "name": "Festivals"
        },
        {
          "taxId": "arts:jazzandblues",
          "name": "Jazz & Blues"
        },
        {
          "taxId": "arts:museums",
          "name": "Museums"
        },
        {
          "taxId": "arts:musicvenues",
          "name": "Music Venues"
        },
        {
          "taxId": "arts:opera",
          "name": "Opera & Ballet"
        },
        {
          "taxId": "arts:theater",
          "name": "Performing Arts"
        },
        {
          "taxId": "arts:sportsteams",
          "name": "Professional Sports Teams"
        },
        {
          "taxId": "arts:psychic_astrology",
          "name": "Psychics & Astrologers"
        },
        {
          "taxId": "arts:racetracks",
          "name": "Race Tracks"
        },
        {
          "taxId": "arts:social_clubs",
          "name": "Social Clubs"
        },
        {
          "taxId": "arts:stadiumsarenas",
          "name": "Stadiums & Arenas"
        },
        {
          "taxId": "arts:ticketsales",
          "name": "Ticket Sales"
        },
        {
          "taxId": "arts:wineries",
          "name": "Wineries"
        }
      ],
      "name": "Arts & Entertainment"
    },
    {
      "legacyCategoryId": "AUTO",
      "taxId": "auto",
      "children": [
        {
          "taxId": "auto:auto_detailing",
          "name": "Auto Detailing"
        },
        {
          "taxId": "auto:autoglass",
          "name": "Auto Glass Services"
        },
        {
          "taxId": "auto:autoloanproviders",
          "name": "Auto Loan Providers"
        },
        {
          "taxId": "auto:autopartssupplies",
          "name": "Auto Parts & Supplies"
        },
        {
          "taxId": "auto:autorepair",
          "name": "Auto Repair"
        },
        {
          "taxId": "auto:boatdealers",
          "name": "Boat Dealers"
        },
        {
          "taxId": "auto:bodyshops",
          "name": "Body Shops"
        },
        {
          "taxId": "auto:car_dealers",
          "name": "Car Dealers"
        },
        {
          "taxId": "auto:stereo_installation",
          "name": "Car Stereo Installation"
        },
        {
          "taxId": "auto:carwash",
          "name": "Car Wash"
        },
        {
          "taxId": "auto:servicestations",
          "name": "Gas & Service Stations"
        },
        {
          "taxId": "auto:motorcycledealers",
          "name": "Motorcycle Dealers"
        },
        {
          "taxId": "auto:motorcyclerepair",
          "name": "Motorcycle Repair"
        },
        {
          "taxId": "auto:oilchange",
          "name": "Oil Change Stations"
        },
        {
          "taxId": "auto:parking",
          "name": "Parking"
        },
        {
          "taxId": "auto:rv_dealers",
          "name": "RV Dealers"
        },
        {
          "taxId": "auto:smog_check_stations",
          "name": "Smog Check Stations"
        },
        {
          "taxId": "auto:tires",
          "name": "Tires"
        },
        {
          "taxId": "auto:towing",
          "name": "Towing"
        },
        {
          "taxId": "auto:truck_rental",
          "name": "Truck Rental"
        },
        {
          "taxId": "auto:windshieldinstallrepair",
          "name": "Windshield Installation & Repair"
        }
      ],
      "name": "Automotive"
    },
    {
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
    }
  ],
  "requestId": "5949a26800ff0a8453f1116a270001737e726570636f72652d70726f6400016170693a3630312d73706c696e7465722d646174616c616b652d6c697374696e672d6d6574726963730001014f",
  "responseTime": 10,
  "statusCode": 200
}`

func Test_ListReturnListOfTaxonomiesOn200(t *testing.T) {
	client := &TaxonomyClient{SDKClient: &basesdk.BaseClientMock{JSONBody: taxonomyResponse}}

	result, err := client.List(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	expectedTopLevelIDs := []string{"active", "arts", "auto", "beautysvc"}
	var topLevelIDs []string

	for _, tax := range result {
		topLevelIDs = append(topLevelIDs, tax.ID)
	}

	for _, ID := range expectedTopLevelIDs {
		assert.Contains(t, topLevelIDs, ID)
	}
}

func Test_ListReturnsErrorWhenCoreReturnsError(t *testing.T) {
	expectedError := errors.New("New error")
	client := &TaxonomyClient{SDKClient: &basesdk.BaseClientMock{JSONBody: taxonomyResponse, Error: expectedError}}

	_, err := client.List(context.Background())

	assert.Equal(t, expectedError, err)
}

func Test_ListReturnsErrorWhenInflatingResponseHasError(t *testing.T) {
	client := &TaxonomyClient{SDKClient: &basesdk.BaseClientMock{JSONBody: `{"data":"garbage"}`}}

	_, err := client.List(context.Background())

	assert.Error(t, err)
}
