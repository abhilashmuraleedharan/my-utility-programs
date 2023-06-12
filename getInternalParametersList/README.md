# Go JSON Parser

This is a Go program that parses a JSON file to extract and sort values from a specific field, 
then writes the output to another JSON file.

## Description

The program reads a JSON file that contains an array of objects. 
Each object has a field specified by the user. 
The program extracts all of these field values, sorts them in ascending order, and writes them to an output JSON file. 
The keys and values of the output JSON file are the same and correspond to the sorted field values.

The program takes three command-line arguments: 
the path to the input JSON file, 
the path to the output JSON file, and 
the name of the field to be extracted.

### Executing program

- Run the script with the command `go run listExtractor.go <input-file> <output-file> <field-name>`. 
Replace, 
`<input-file>` with the path to your input JSON file, 
`<output-file>` with the path where you want to write the output JSON file, and 
`<field-name>` with the name of the field you want to extract.

For e.g.
go run listExtractor.go sample_input.json output.json InternalParameter
