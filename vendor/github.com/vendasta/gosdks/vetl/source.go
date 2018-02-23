package vetl

import (
	"github.com/vendasta/gosdks/pb/vetl/v1"
	"github.com/vendasta/gosdks/vstore"
)

type sourceOption struct {
	source *vetl_v1.DataSource
}

type Source func(r *sourceOption)

//DataSourceFromVStoreModel returns a vetl Source, complete with schema and provider definitions, given a reference to a vstore.Model
func DataSourceFromVStoreModel(model vstore.Model, pubsubIndexID string) (Source, error) {
	v := model.Schema()
	vp := vStoreProvider(v.Namespace, v.Kind, v.PrimaryKey, pubsubIndexID)

	props, err := getPropertiesFromVStoreProperties(v.Properties)
	if err != nil {
		return nil, err
	}

	return func(r *sourceOption) {
		r.source = &vetl_v1.DataSource{
			Provider: &vetl_v1.DataSource_Vstore{
				Vstore: vp,
			},
			Schema: &vetl_v1.Schema{
				Properties: props,
			},
		}
	}, nil
}
