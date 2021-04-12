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

func ReadFile(mappingFile string) []string {
	content, err := os.ReadFile(mappingFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {
	jsonData := Mappings{
		Mappings: Properties{
			Properties: Type{
				Type: "ip",
			},
		},
	}
	mappingFile, _ := json.MarshalIndent(jsonData, "", "    ")
	fmt.Printf("%+v\n", string(mappingFile))
	for _, line := range ReadFile("mapping.txt") {
		fmt.Print(line + string('\n'))
	}
	//_ = ioutil.WriteFile("mappings.json", mappingFile, 0644)
}
