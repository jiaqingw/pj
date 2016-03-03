package main

import "encoding/json"
import "io/ioutil"
import "os"
import "log"

func main() {
	inputFile := ""

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	var jsonStream []byte
	var err error

	// Either read from the input file if it's provided or read
	// from the standard input
	if inputFile == "" {
		jsonStream, err = ioutil.ReadAll(os.Stdin)
	} else {
		jsonStream, err = ioutil.ReadFile(inputFile)
	}

	//unmarshall into a json value
	var jsonValue interface{}
	if err == nil {
		json.Unmarshal(jsonStream, &jsonValue)
		prettyJson, err := json.MarshalIndent(jsonValue, "", "  ")

		if err == nil {
			os.Stdout.Write(prettyJson)
		} else {
			log.Fatalln("Errors:", err)
		}

	} else {
		log.Fatalln("Error:", err)
	}

}
