package main

import (
	"encoding/json"
	"fmt"

	kin "github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v2"
)

var testJSON = `{
    "id": 35,
    "name": "Good Cause",
    "legalName": "Good Cause Inc",
    "slug": "goodcause",
    "currency": "USD",
    "about": "We do good stuff",
    "video": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
    "image": {
      "src": "https://asdf.cloudfront.net/asdf.jpg",
      "alt": "",
      "width": 640,
      "height": 360
    },
    "avatar": {
      "src": "https://asdf.cloudfront.net/asdf.jpg",
      "alt": "",
      "width": 64,
      "height": 64
    },
    "logo": {
      "src": "https://asdf.cloudfront.net/asdf.jpg",
      "alt": "",
      "width": 100,
      "height": 64
    },
    "banner": {
      "src": "https://asdf.cloudfront.net/asdf.jpg",
      "alt": "",
      "width": 530,
      "height": 1440
    },
    "contactEmail": "asdf@example.org",
    "paypalEmail": "asdf@example.org",
    "paypalCurrencyCode": "USD",
    "social": {
      "twitter": "stjude",
      "youtube": "",
      "facebook": "stjude",
      "instagram": "stjude"
    },
    "settings": {
      "colors": {
        "background": "#fff",
        "highlight": "#ce2f40"
      },
      "headerIntro": "Start a fundraising campaign to help us do good things.",
      "headerTitle": "Good Cause",
      "footerCopyright": "© Copyright 2017. Good Cause Inc",
      "findOutMoreLink": "https://www.example.com/find-out-more"
    },
    "status": "published",
    "stripeConnected": true,
    "mailchimpConnected": false,
    "address": {
      "addressLine1": "123 Some St",
      "addressLine2": "",
      "city": "Houston",
      "region": "TX",
      "postalCode": "77008",
      "country": "US"
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
