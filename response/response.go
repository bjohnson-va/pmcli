package response

import (
	"github.com/bjohnson-va/pmcli/random"
	"github.com/bjohnson-va/pmcli/parse"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
	"context"
	"github.com/emicklei/proto"
	"strings"
	"fmt"
)

func GenerateForMessage(ctx context.Context, p *random.FieldProvider, breadcrumb BreadCrumb, message proto.Message,
	t *parse.FieldTypes, c config.InputsProvider) interface{} {
	obj := make(map[string]interface{})
	fieldz := parse.Fields(message.Elements)
	for _, f := range fieldz {
		fBreadcrumb := breadcrumb.AndField(f.Name)
		if c.GetFieldExclusion(fBreadcrumb.ToString()) {
			logging.Debugf(ctx, "%s is excluded via config file", fBreadcrumb)
			continue;
		}
		var value interface{}
		if f.Repeated {
			// json unmarshal defaults to float64
			instruction := c.GetFieldInstruction(fBreadcrumb.ToString(), "num", 1.0)
			instFlt64, ok := instruction.(float64)
			if !ok {
				logging.Errorf(ctx, "Unexpected value for num: %#v", instruction)
			}
			length := int(instFlt64)
			var list []interface{}
			for x := 0; x < length; x++ {
				z, err := randomFieldValue(ctx, *p, fBreadcrumb.Indexed(x), *f.Field, t, c)
				if err != nil {
					value = err.Error()
					break
				}
				list = append(list, z)
			}
			value = list
		} else {
			var err error
			value, err = randomFieldValue(ctx, *p, fBreadcrumb.Dive(), *f.Field, t, c)
			if err != nil {
				value = err.Error() // Expose the error to the user of the API
			}
		}
		obj[util.ToCamelCase(f.Name)] = value
	}
	return obj
}

func randomFieldValue(ctx context.Context, p random.FieldProvider, crumb BreadCrumb, field proto.Field, t *parse.FieldTypes, c config.InputsProvider) (interface{}, error) {
	override := c.GetFieldOverride(crumb.ToString(), nil)
	if override != nil {
		logging.Infof(ctx, "Using override for %s: %v", crumb, override)
		return override, nil
	}
	supercrumb := crumb.AsSuperCrumb()
	if field.Type == "string" || field.Type == "bytes" {
		return p.NewString(supercrumb), nil
	}
	if strings.Contains(field.Type, "int") {
		return p.NewInt32(supercrumb), nil
	}
	if strings.Contains(field.Type, "float") {
		return p.NewFloat32(supercrumb), nil
	}
	if strings.Contains(field.Type, "double") {
		return p.NewFloat64(supercrumb), nil
	}
	if strings.Contains(field.Type, "bool") {
		return p.NewBool(supercrumb), nil
	}
	if strings.Contains(field.Type, "google.protobuf.Timestamp") {
		return p.NewTimestamp(supercrumb), nil // TODO: Use correct format
	}

	var isEnum bool
	fieldType := field.Type
	if strings.Contains(field.Type, ".") {
		// Probably an enum.  Eg: CampaignStatus.Status
		parts := strings.Split(field.Type, ".")
		fieldType = parts[0]
		isEnum = true;
	}

	for _, e := range t.Enums {
		if fieldType == e.Name {
			return p.NewEnumValue(supercrumb, e), nil
		}
	}

	m := t.GetMatchingMessage(fieldType)
	if m == nil {
		return "", fmt.Errorf("unexpected field type %s", field.Type)
	}

	if isEnum {
		for _, e := range parse.Enums(m.Elements) {
			return p.NewEnumValue(supercrumb, e), nil
		}
	}
	return GenerateForMessage(ctx, &p, crumb.Dive(), *m, t, c), nil

}
