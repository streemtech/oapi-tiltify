package main

import (
	"encoding/json"
	"fmt"

	kin "github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v2"
)

var testJSON = `{
	"id": 387,
	"name": "DerpSquad",
	"slug": "derpsquad",
	"url": "https://tiltify.com/+derpsquad",
	"bio": "",
	"inviteOnly": true,
	"disbanded": false,
	"social": {
	  "facebook": "tradechatonyoutube",
	  "twitch": "tradechat",
	  "twitter": "tradechat",
	  "website": "",
	  "youtube": "tradechat"
	},
	"avatar": {
	  "src": "https://asdf.cloudfront.net/asdf.jpg",
	  "alt": "",
	  "width": 200,
	  "height": 200
	}
  }`

//convert json data to an openapi schema.
func main() {

	var dest map[string]any
	err := json.Unmarshal([]byte(testJSON), &dest)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", dest)

	s := kin.NewSchema()
	s.Properties = mapToSchemas(dest)

	d, err := s.MarshalJSON()
	if err != nil {
		panic(err)
	}

	dest = make(map[string]any)
	err = json.Unmarshal(d, &dest)
	if err != nil {
		panic(err)
	}

	res, err := yaml.Marshal(&dest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", res)
}

func mapToSchemas(m map[string]any) kin.Schemas {

	out := kin.Schemas{}

	for k, v := range m {
		out[k] = anyToSchema(v)
	}
	return out
}

func anyToSchema(v any) *kin.SchemaRef {
	out := &kin.SchemaRef{}
	switch x := v.(type) {
	case float64:

		if isIntegral(x) {
			out.Value = &kin.Schema{
				Type: kin.TypeInteger,
			}
		} else {
			out.Value = &kin.Schema{
				Type: kin.TypeNumber,
			}
		}
	case string:
		out.Value = &kin.Schema{
			Type: kin.TypeString,
		}
	case bool:
		out.Value = &kin.Schema{
			Type: kin.TypeBoolean,
		}
	case map[string]any:
		out.Value = &kin.Schema{
			Type:       kin.TypeObject,
			Properties: mapToSchemas(x),
		}
	case []any:
		out.Value = &kin.Schema{
			Type: kin.TypeArray,
		}
		if len(x) > 0 {
			out.Value.Items = anyToSchema(x[0])
		}
	default:
		//assume nil.
		out.Value = &kin.Schema{
			Type: kin.TypeObject,
		}
	}

	return out
}

func isIntegral(val float64) bool {
	return val == float64(int64(val))
}
