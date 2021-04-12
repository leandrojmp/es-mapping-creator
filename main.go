package main

import (
	"encoding/json"
	"fmt"
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
	fmt.Printf("%+v\n", string(mappingFile))
	//_ = ioutil.WriteFile("mappings.json", mappingFile, 0644)
}
