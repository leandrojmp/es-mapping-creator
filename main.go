package main

import (
	"encoding/json"
	"io/ioutil"
)

type Mappings struct {
	Mappings Properties `json:"mappings"`
}

type Properties struct {
	Properties string `json:"properties"`
}

func main() {
	jsonData := Mappings{
		Mappings: Properties{
			Properties: "teste",
		},
	}
	mappingFile, _ := json.MarshalIndent(jsonData, "", "    ")
	_ = ioutil.WriteFile("mappings.json", mappingFile, 0644)
}
