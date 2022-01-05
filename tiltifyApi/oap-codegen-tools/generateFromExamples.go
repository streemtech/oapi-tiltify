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

	jsonReg := regexp.MustCompile(`.*\.json`)
	filepath.WalkDir("./examples", func(path string, d fs.DirEntry, err error) error {

		if !jsonReg.MatchString(path) {
			return nil
		}
		fmt.Printf("Walking %s\n", path)

		f, err := os.Open(path)
		if err != nil {
			fmt.Printf("ERROR OPEING FILE: %s", err.Error())
		}
		data, err := ioutil.ReadAll(f)
		f.Close()
		if err != nil {
			fmt.Printf("ERROR READING FILE: %s", err.Error())
		}
		res, err := jsonExampleToYamlProperties(data)
		if err != nil {
			fmt.Printf("ERROR CONVERTING FILE: %s", err.Error())
		}
		dest := strings.ReplaceAll(path, ".json", ".yaml")
		dest = strings.ReplaceAll(dest, "examples", "data")
		fmt.Printf("Creating %s\n", dest)

		f, err = os.Create(dest)
		if err != nil {
			fmt.Printf("ERROR CREATING FILE: %s", err.Error())
		}
		_, err = f.Write(res)
		if err != nil {
			fmt.Printf("ERROR WRITING FILE: %s", err.Error())
		}

		return nil
	})
}

func jsonExampleToYamlProperties(data []byte) (out []byte, err error) {

	var dest any

	err = json.Unmarshal(data, &dest)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to unmarshal JSON: %w", err)
	}

	// fmt.Printf("%+v\n", dest)

	s := anyToSchema(dest)

	d, err := s.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to marshal results to JSON: %w", err)
	}

	dest = *new(any)
	err = json.Unmarshal(d, &dest)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to Unmarshal output json to map: %w", err)
	}

	res, err := yaml.Marshal(&dest)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to marshal to YAML: %w", err)
	}
	return res, nil
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
