package main

import (
	"encoding/json"
	"os"

	m "github.com/leandrojmp/es-mapping-creator/mappings"
)

func main() {
	jsonData := m.Mappings{Mappings: m.Properties{Properties: m.CreateMapping("fields.txt")}}
	mappingFile, _ := json.MarshalIndent(jsonData, "", "    ")
	_ = os.WriteFile("es-mappings.json", mappingFile, 0644)
	// fmt.Printf("%+v\n", string(mappingFile))
}
