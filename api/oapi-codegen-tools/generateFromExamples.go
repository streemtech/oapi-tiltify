package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	kin "github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v2"
)

func main() {
	generate()
}
func generate() {

	t := &kin.T{
		Components: kin.Components{
			Schemas: kin.Schemas{},
		},
	}

	jsonReg := regexp.MustCompile(`.*\.json`)
	err := filepath.WalkDir("./examples", func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}
		if !jsonReg.MatchString(path) {
			return nil
		}
		// fmt.Printf("Walking %s\n", path)

		//open file and read in data.
		f, err := os.Open(path)
		if err != nil {
			fmt.Printf("ERROR OPEING FILE: %s", err.Error())
		}
		data, err := ioutil.ReadAll(f)
		f.Close()
		if err != nil {
			fmt.Printf("ERROR READING FILE: %s", err.Error())
		}

		res, err := jsonToSchema(data)
		if err != nil {
			fmt.Printf("ERROR CONVERTING FILE: %s", err.Error())
		}

		name := generateName(path)

		_, ok := t.Components.Schemas[name]
		if ok {
			fmt.Printf("Duplicate Name %s", name)
		}
		t.Components.Schemas[name] = res

		return nil
	})

	if err != nil {
		fmt.Printf("ERROR WALKING: %s", err.Error())
		return
	}

	dat, err := toYamlData(t)
	if err != nil {
		fmt.Printf("ERROR CREATING FILE: %s", err.Error())
		return
	}

	f, err := os.Create("./schemas.yaml")
	if err != nil {
		fmt.Printf("ERROR CREATING FILE: %s", err.Error())
		return
	}
	_, err = f.Write(dat)
	if err != nil {
		fmt.Printf("ERROR WRITING FILE: %s", err.Error())
	}
}

func toYamlData(s interface{ MarshalJSON() ([]byte, error) }) ([]byte, error) {
	d, err := s.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to marshal results to JSON: %w", err)
	}

	dest := *new(any)
	err = json.Unmarshal(d, &dest)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to Unmarshal output json to map: %w", err)
	}

	res, err := yaml.Marshal(&dest)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to marshal to YAML: %w", err)
	}
	return res, nil
}

func jsonToSchema(data []byte) (*kin.SchemaRef, error) {

	var dest any

	err := json.Unmarshal(data, &dest)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON: %w", err)
	}

	// fmt.Printf("%+v\n", dest)

	return anyToSchema(dest), nil

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
		//TODO3 merge items in array
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

func generateName(path string) (name string) {
	dest := strings.ReplaceAll(path, ".json", "")
	dest = strings.ReplaceAll(dest, "examples/", "")
	dest = kebabToCamelCase(dest)
	fmt.Printf("%s->%s\n", path, dest)
	return dest
}

func kebabToCamelCase(kebab string) (camelCase string) {
	isToUpper := true
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}
