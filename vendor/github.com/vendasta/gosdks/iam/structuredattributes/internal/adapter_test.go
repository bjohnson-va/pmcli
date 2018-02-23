package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/pb/iam/v1"
	"github.com/vendasta/gosdks/iam/attribute"
	"time"
)

func Test_StructuredAttributeToModel(t *testing.T) {
	type InnerAttributeModel struct {
		Version  int    `attribute:"version"`
		LastName string `attribute:"last_name"`
	}
	type StructuredAttributeModel struct {
		Name            string                    `attribute:"name"`
		Version         int                       `attribute:"version"`
		Deleted         bool                      `attribute:"deleted"`
		Height          float64                   `attribute:"height"`
		InnerAttribute  *InnerAttributeModel      `attribute:"inner_attribute"`
		AssociatedUsers []string                  `attribute:"associated_users"`
		Updated         time.Time                 `attribute:"updated"`
		Location        *iam_v1.GeoPointAttribute `attribute:"location"`
	}

	testDate := time.Now().UTC()

	cases := []struct {
		name   string
		input  *iam_v1.StructAttribute
		output StructuredAttributeModel
		error  error
	}{
		{
			name: "Extracts string attribute",
			input: attribute.NewBuilder().
				String("name", "Hank").
				Build(),
			output: StructuredAttributeModel{
				Name: "Hank",
			},
		},
		{
			name: "Extracts int attribute",
			input: attribute.NewBuilder().
				Int("version", 64).
				Build(),
			output: StructuredAttributeModel{
				Version: 64,
			},
		},
		{
			name: "Extracts bool attribute",
			input: attribute.NewBuilder().
				Bool("deleted", true).
				Build(),
			output: StructuredAttributeModel{
				Deleted: true,
			},
		},
		{
			name: "Extracts double attribute",
			input: attribute.NewBuilder().
				Float("height", 75.5).
				Build(),
			output: StructuredAttributeModel{
				Height: 75.5,
			},
		},
		{
			name: "Extracts struct attribute",
			input: attribute.NewBuilder().
				Struct("inner_attribute", attribute.NewBuilder().
					Int("version", 75).
					String("last_name", "Hill"),
				).
				Build(),
			output: StructuredAttributeModel{
				InnerAttribute: &InnerAttributeModel{
					Version:  75,
					LastName: "Hill",
				},
			},
		},
		{
			name: "Extracts list attribute",
			input: attribute.NewBuilder().
				Strings("associated_users", []string{"UID-123", "UID-abc"}).
				Build(),
			output: StructuredAttributeModel{
				AssociatedUsers: []string{"UID-123", "UID-abc"},
			},
		},
		{
			name: "Extracts time attribute",
			input: attribute.NewBuilder().
				Time("updated", testDate).
				Build(),
			output: StructuredAttributeModel{
				Updated: testDate,
			},
		},
		{
			name: "Extracts GeoPoint attribute",
			input: attribute.NewBuilder().
				GeoPoint("location", 64, 10).
				Build(),
			output: StructuredAttributeModel{
				Location: &iam_v1.GeoPointAttribute{
					Latitude: 64,
					Longitude: 10,
				},
			},
		},
		{
			name: "Complex structure extractor",
			input: attribute.NewBuilder().
				String("name", "Hank").
				Int("version", 64).
				Bool("deleted", true).
				Float("height", 75.5).
				Struct("inner_attribute", attribute.NewBuilder().
					Int("version", 75).
					String("last_name", "Hill"),
				).
				Strings("associated_users", []string{"UID-123", "UID-abc"}).
				Time("updated", testDate).
				GeoPoint("location", 64, 10).
				Build(),
			output: StructuredAttributeModel{
				Name: "Hank",
				Version: 64,
				Deleted: true,
				Height: 75.5,
				InnerAttribute: &InnerAttributeModel{
					Version: 75,
					LastName: "Hill",
				},
				AssociatedUsers: []string{"UID-123", "UID-abc"},
				Updated: testDate,
				Location: &iam_v1.GeoPointAttribute{
					Latitude: 64,
					Longitude: 10,
				},
			},
		},
		{
			name: "Missing field is ignored",
			input: attribute.NewBuilder().
				String("name", "Hank").
				Int("version", 64).
				String("missing", "missing").
				Build(),
			output: StructuredAttributeModel{
				Name: "Hank",
				Version: 64,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := StructuredAttributeModel{}
			err := StructuredAttributeToModel(&actual, c.input)
			assert.Equal(t, c.error, err)
			assert.Equal(t, c.output, actual)
		})
	}
}
