package vetl

import (
	"github.com/vendasta/gosdks/pb/vetl/v1"
	"github.com/vendasta/gosdks/vstore"
	"github.com/vendasta/gosdks/util"
)

func vStoreProvider(namespace, kind string, primaryKey []string, pubsubIndexID string) *vetl_v1.VStoreSource {
	return &vetl_v1.VStoreSource{
		Namespace: namespace,
		Kind: kind,
		PrimaryKey: primaryKey,
		PubsubIndexId: pubsubIndexID,
	}
}

func getPropertiesFromVStoreProperties(vp []*vstore.Property) ([]*vetl_v1.Property, error) {
	results := make([]*vetl_v1.Property, len(vp))
	for i, p := range vp {
		s := &vetl_v1.Property{
			Name: p.Name,
			Repeated: p.IsRepeated,
			Required: p.IsRequired,
		}

		switch p.FType {
		case vstore.StringType:
			s.Type = vetl_v1.Property_STRING
		case vstore.GeoPointType:
			s.Type = vetl_v1.Property_GEOPOINT
		case vstore.BoolType:
			s.Type = vetl_v1.Property_BOOL
		case vstore.IntType:
			s.Type = vetl_v1.Property_INT64
		case vstore.FloatType:
			s.Type = vetl_v1.Property_DOUBLE
		case vstore.TimeType:
			s.Type = vetl_v1.Property_TIMESTAMP
		case vstore.StructType:
			s.Type = vetl_v1.Property_STRUCT
			if len(p.Properties) > 0 {
				sub, err := getPropertiesFromVStoreProperties(p.Properties)
				if err != nil {
					return nil, err
				}
				s.Properties = sub
			}
		default:
			return nil, util.Error(util.InvalidArgument, "Unsupported VStore type provided")
		}
		results[i] = s
	}
	return results, nil
}
