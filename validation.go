package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

func validate(scheme_file, data_file string) {
	// Read the JSON schema file
	schemaFile, err := ioutil.ReadFile(scheme_file)
	if err != nil {
		fmt.Println("Error reading schema file:", err)
		os.Exit(1)
	}

	// Parse the JSON schema
	schemaLoader := gojsonschema.NewStringLoader(string(schemaFile))

	// Read the JSON data file
	dataFile, err := ioutil.ReadFile(data_file)
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
		panic(err)
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
		panic("Cross check the data with Scheme!")
	}
}

func main() {
	fmt.Println("Validating Country Information!")
	validate("country_info_scheme.json", "country_info_data.json")

	fmt.Println("\n Validating Country Language Information!")
	validate("country_language_scheme.json", "country_language_data.json")

	fmt.Println("\n Validating Country Currency Information!")
	validate("country_currency_scheme.json", "country_currency_data.json")

	fmt.Println("\n Validating Country tele Information!")
	validate("country_tele_scheme.json", "country_tele_data.json")

	fmt.Println("\n Validating Country Timezones Information!")
	validate("country_timezones_scheme.json", "country_timezones_data.json")

}
