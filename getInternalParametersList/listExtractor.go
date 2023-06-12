package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
)

type Parameter struct {
	ColumnNumber       int      `json:"ColumnNumber"`
	Grouping           bool     `json:"Grouping"`
	ColumnName         string   `json:"ColumnName"`
	ColumnType         string   `json:"ColumnType"`
	Mandatory          bool     `json:"Mandatory"`
	InternalParameter  string   `json:"InternalParameter"`
	Size               string   `json:"Size"`
	Default            string   `json:"Default"`
	ColumnDelimiter    string   `json:"ColumnDelimiter"`
	MandatoryForList   []string `json:"MandatoryForList"`
	EncloseWith        string   `json:"EncloseWith"`
}

type Body struct {
	Parameters []Parameter `json:"Parameters"`
}

func main() {
	// Check command line arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run <program-name>.go <input-file> <output-file> <field-name>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	fieldName := os.Args[3]

	// Read the JSON file
	jsonFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	// Read the opened JSON file as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize the Body struct
	var body Body

	// Unmarshal the byte value into the Body struct
	json.Unmarshal(byteValue, &body)

	// Extract and sort the "InternalParameter" values
	var fieldValues []string
	for _, param := range body.Parameters {
		r := reflect.ValueOf(param)
		f := reflect.Indirect(r).FieldByName(fieldName)
		if f.IsValid() {
			fieldValues = append(fieldValues, f.String())
		}
	}
	sort.Strings(fieldValues)

	// Prepare the output map for JSON
	outputMap := make(map[string]string)
	for _, value := range fieldValues {
		outputMap[value] = value
	}

	// Convert map to JSON
	jsonOutput, err := json.MarshalIndent(outputMap, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	// Write the output JSON to a file
	err = ioutil.WriteFile(outputFile, jsonOutput, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
