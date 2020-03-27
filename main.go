package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var jsonInputFilename string
var templateFilename string
var outputFilename string
var kvpArgs []string

func main() {
	if len(os.Args) < 5 {
		log.Fatalf("at least 3 arguments expected: jsonInputFile, templateFile, outputFilename")
	}

	jsonInputFilename = os.Args[1]
	templateFilename = os.Args[2]
	outputFilename = os.Args[3]
	kvpArgs = os.Args[4:]

	var inputJson map[string]interface{}
	//var inputJson = Atom{}
	readJson(&inputJson)
	log.Infof("read json:\n %+v", inputJson)
	addKvpArgumentsToMap(kvpArgs, inputJson)
	log.Infof("added argumnts :\n %+v", inputJson)
	templ := template.Must(template.ParseFiles(templateFilename))

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("create file: %v", err)
		return
	}

	err = templ.Execute(outputFile, inputJson)
	if err != nil {
		log.Fatalf("templating: %v", err)
	}
}

func readJson(inputJson *map[string]interface{}) {
	jsonFile, err := ioutil.ReadFile(jsonInputFilename)
	if err != nil {
		log.Fatalf("could not read file %s; %v ", jsonInputFilename, err)
	}
	err = json.Unmarshal([]byte(jsonFile), inputJson)
	if err != nil {
		log.Fatalf("could not unmarshal %s; %v", jsonInputFilename, err)
	}
}

func addKvpArgumentsToMap(args []string, jmap map[string]interface{}) {
	for _, arg := range args {
		kvp := strings.Split(arg, "=")
		if len(kvp) != 2 {
			log.Fatalf("kvp argument %s may contain only one '=' sign and should contain a key and a value", arg)
		}
		keyPath := strings.Split(kvp[0], ".")
		addValueToMap(keyPath, kvp[1], jmap)
	}
}

func addValueToMap(keyPath []string, value string, kvMap map[string]interface{}) {
	pointer := kvMap
	for index, key := range keyPath {
		isLastKeyInPath := len(keyPath)-1 == index
		if isLastKeyInPath {
			// set the value
			pointer[key] = value
		} else {
			// traverse down the tree
			if pointer[key] == nil {
				pointer[key] = make(map[string]interface{})
			}
			pointer = pointer[key].(map[string]interface{})
		}
	}
}
