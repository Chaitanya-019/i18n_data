package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// Read the JSON schema file
	schemaFile, err := ioutil.ReadFile("country_tele_scheme.json")
	if err != nil {
		fmt.Println("Error reading schema file:", err)
		os.Exit(1)
	}

	// Parse the JSON schema
	schemaLoader := gojsonschema.NewStringLoader(string(schemaFile))

	// Read the JSON data file
	dataFile, err := ioutil.ReadFile("country_tele_data.json")
	if err != nil {
		fmt.Println("Error reading data file:", err)
		os.Exit(1)
	}

	// Parse the JSON data
	dataLoader := gojsonschema.NewStringLoader(string(dataFile))

	// Load schema and data
	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Check validation result
	if result.Valid() {
		fmt.Println("The JSON data is valid.")
	} else {
		fmt.Println("The JSON data is not valid.")
		fmt.Println("Errors:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
