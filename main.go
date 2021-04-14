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

func ReadFile(mappingFile string) []string {
	content, err := os.ReadFile(mappingFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func CreateMapping(lines []string) interface{} {
	mappingFields := make(map[string]interface{})
	for _, line := range lines {
		typeField := make(map[string]interface{})
		typeField["type"] = strings.Split(line, ": ")[1]
		mappingFields[strings.Split(line, ": ")[0]] = typeField
	}
	return mappingFields
}

func main() {
	CreateMapping(ReadFile("mapping.txt"))
	jsonData := Mappings{
		Mappings: Properties{
			CreateMapping(ReadFile("mapping.txt")),
		},
	}
	mappingFile, _ := json.MarshalIndent(jsonData, "", "    ")
	fmt.Printf("%+v\n", string(mappingFile))
}
