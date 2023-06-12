package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type InputRecord map[string]map[string]interface{}

type OutputRecord map[string]interface{}

func main() {
	// Check command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program <input.json> <output.json>")
		os.Exit(1)
	}

	// Read the JSON file
	inputFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// Parse the JSON file
	var input []InputRecord
	err = json.Unmarshal(inputFile, &input)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %s", err)
	}

	// Transform the JSON
	var output []OutputRecord
	for _, item := range input {
		newItem := make(OutputRecord)

		for key, value := range item {
			for _, v := range value {
				newItem[key] = v
			}
		}

		output = append(output, newItem)
	}

	// Convert the output to JSON
	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("Failed to create JSON: %s", err)
	}

	// Write the JSON to a file
	err = ioutil.WriteFile(os.Args[2], outputJSON, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}
}
