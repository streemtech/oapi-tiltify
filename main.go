package main

import (
	"fmt"
	"io/ioutil"
	"os"

	kin "github.com/getkin/kin-openapi/openapi3"
)

var file = "tiltify.openapi.yml"
var out = "tiltify.openapi.output.yml"

func main() {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	dat, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	//var inYaml kin.Schema{}
	l := kin.NewLoader()
	l.IsExternalRefsAllowed = true

	t, err := l.LoadFromData(dat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v/n", t)
}
