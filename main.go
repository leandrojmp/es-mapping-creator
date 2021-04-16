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

func CreateMapping(mappingFile string) interface{} {
	// read file and create list with the content
	content, err := os.ReadFile(mappingFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	lines := strings.Split(string(content), "\n")
	// for each line add item as a key in a dict with the type
	mappingFields := make(map[string]interface{})
	for _, line := range lines {
		typeField := make(map[string]interface{})
		typeField["type"] = strings.Split(line, ": ")[1]
		mappingFields[strings.Split(line, ": ")[0]] = typeField
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
}
