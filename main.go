package main

import (
	"encoding/json"
	"flag"
	"os"

	m "github.com/leandrojmp/es-mapping-creator/mappings"
)

func main() {
	inputFilePtr := flag.String("in", "mapping.txt", "txt input file with the fields and types")
	outputFilePtr := flag.String("out", "es-mappings.json", "json output file with the mappings")

	flag.Parse()

	jsonData := m.Template{Template: m.Mappings{Mappings: m.Properties{Properties: m.CreateMapping(*inputFilePtr)}}}
	mappingFile, _ := json.MarshalIndent(jsonData, "", "    ")
	_ = os.WriteFile(*outputFilePtr, mappingFile, 0644)
	// fmt.Printf("%+v\n", string(mappingFile))
}
