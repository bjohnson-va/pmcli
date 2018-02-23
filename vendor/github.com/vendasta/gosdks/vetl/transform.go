package vetl

import "github.com/vendasta/gosdks/pb/vetl/v1"

type transformOption struct {
	transform *vetl_v1.Transform
}

type Transform func(r *transformOption)

func KeepPropertiesTransform(names []string) Transform {
	return func(r *transformOption) {
		r.transform = &vetl_v1.Transform{
			Transform: &vetl_v1.Transform_KeepProperties{
				KeepProperties: &vetl_v1.KeepProperties{
					Names: names,
				},
			},
		}
	}
}