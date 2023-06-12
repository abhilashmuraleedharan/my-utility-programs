# JSON Transformer

This is a Go program designed to parse and transform JSON records from a specified input file and 
write the transformed records to a specified output file.

The program is built to transform JSON records of the following input format:

json

[
    {
        "key1": {
            "int": 1
        },
        "key2": {
            "string": "hello"
        }
    },
    // ...additional records...
]

Into JSON records of the following output format:

json

[
    {
        "key1": 1,
        "key2": "hello"
    },
    // ...additional records...
]

## Usage


go

go build transform.go

Then run the program with your input and output files as arguments:

css

./transform input.json output.json

## How It Works

The program reads the JSON file specified by the first argument (input.json in the example above). 
It parses each JSON record into a Go map structure, where each key is a string and each value is 
another map with string keys and interface{} values.

Then it transforms each parsed record by iterating over the key-value pairs in the map. 
It creates a new map (representing the transformed record) where each key corresponds directly to a single value.

Finally, the program converts the transformed records back into JSON format (with pretty-print formatting), 
and writes the JSON to the file specified by the second argument (output.json in the example above).

## Error Handling

The program includes error handling for the following scenarios:

    1. The input file cannot be read.
    2. The input file cannot be parsed as JSON.
    3. The transformed records cannot be converted back into JSON.
    4. The output file cannot be written.

In each case, the program will log an error message and exit.
