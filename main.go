package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Mappings struct {
	Mappings Properties `json:"mappings"`
}

type Properties struct {
	Properties interface{} `json:"properties"`
}

type Type struct {
	Type string `json:"type"`
}

func CreateMapping(mappingFile string) interface{} {
	// read file and create list with the content
	content, err := os.ReadFile(mappingFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	lines := strings.Split(string(content), "\n")
	// for each line add item as a key in a dict with the type
	mappingFields := make(map[string]interface{})
	nestedFields := make(map[string]map[string]interface{})
	topFields := make(map[string]string)
	var nestedList []string
	// check for nested, works only for two level in the format first.second: type
	for _, line := range lines {
		fieldName := strings.Split(line, ": ")[0]
		fieldType := strings.Split(line, ": ")[1]
		if strings.Contains(fieldName, ".") {
			nestedList = append(nestedList, strings.Replace(line, ": ", "-", 1))
		} else {
			mappingFields[fieldName] = Type{fieldType}
		}
	}
	// loop nested fields list
	for _, field := range nestedList {
		fieldTop := strings.Split(strings.Split(field, "-")[0], ".")[0]
		topFields[fieldTop] = fieldTop
		if nestedFields[fieldTop] == nil {
			nestedFields[fieldTop] = make(map[string]interface{})
		}
		fieldName := strings.Split(strings.Split(field, "-")[0], ".")[1]
		fieldType := strings.Split(field, "-")[1]
		nestedFields[fieldTop][fieldName] = Type{fieldType}
	}
	// create mapping for top level nested fields
	for _, top := range topFields {
		mappingFields[top] = Properties{nestedFields[top]}
	}
	return mappingFields
}

func main() {
	jsonData := Mappings{
		Mappings: Properties{
			CreateMapping("mapping.txt"),
		},
	}
	mappingFile, _ := json.MarshalIndent(jsonData, "", "    ")
	_ = os.WriteFile("es-mappings.json", mappingFile, 0644)
	// fmt.Printf("%+v\n", string(mappingFile))
}
