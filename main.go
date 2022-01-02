package main

import (
	"encoding/json"
	"fmt"

	kin "github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v2"
)

var testJSON = `{
    "id": 1,
    "name": "My Awesome Campaign",
    "slug": "my-awesome-campaign",
    "url": "/@username/my-awesome-campaign",
    "startsAt": 1493355600000,
    "endsAt": 1496206800000,
    "description": "My awesome weekend campaign.",
    "avatar": {
      "src": "https://asdf.cloudfront.net/asdf.jpg",
      "alt": "",
      "width": 200,
      "height": 200
    },
    "causeId": 17,
    "fundraisingEventId": 39,
    "fundraiserGoalAmount": 10000,
    "originalGoalAmount": 5000,
    "amountRaised": 3402.00,
    "supportingAmountRaised": 8923.00,
    "totalAmountRaised": 12325.00,
    "supportable": true,
    "status": "published",
    "user": {
      "id": 1,
      "username": "UserName",
      "slug": "username",
      "url": "/@username",
      "avatar": {
        "src": "https://asdf.cloudfront.net/asdf.jpg",
        "alt": "",
        "width": 200,
        "height": 200
      }
    },
    "team": {
      "id": 1,
      "username": "Team Name",
      "slug": "teamslug",
      "url": "/+teamslug",
      "avatar": {
        "src": "https://asdf.cloudfront.net/asdf.jpg",
        "alt": "",
        "width": 200,
        "height": 200
      }
    },
    "livestream": {
      "type": "twitch",
      "channel": "tiltify"
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
