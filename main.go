package main

import (
	"fmt"
	"flag"
	"errors"
	"log"
	"net/http"
)

// These values represent the default values that the flags should get
// if not specified when running the program.
const (
	PORTDEFAULT      = 8080 // Default port number if none is specified.
	FORMALITYDEFAULT = 0    // Defaults to most formal greeting if not specified.
	LANGUAGEDEFAULT = "english" // Defaults to english if not specified.
)

var (
	var nonValidPort = errors.New("Port Number not Valid")
	var nonValidFormality = errors.New("Formaility value not valid")
	var nonValidLanguage = errors.New("Non-Valid Language")
)

var pPort *int64
var pFormality *int64

func init() {
	pPort = flag.Int64("port", PORTDEFAULT, "Specify the port the server should be listening on. 8080 by default.")
	pFormality = flag.Int64("formality", FORMALITYDEFAULT, "Formal Greeting==0, Semi-Formal Greeting==1, Informal Greeting==2")
	pLanguage = flag.String("lang", LANGUAGEDEFAULT, "Options for Greeting: 'english', 'spanish', 'russian'. English by default.")
	flag.Parse()

	err := checkFlags(pPort, pFormality, pLanguage)
	errorCheck(err)
}

func main() {
	fmt.Println("Howdy")
	fmt.Println("English Greeting listening on port 8081")
	fmt.Println("Spanish Greeting listens on port 8082")
	fmt.Println("Russian Greeting listens on port 8083")
}

func checkFlags(port *int64, formality *int64, language *string) error {
	if *port < 0 {
		return nonValidPort
	}
	if *formality < 0 || *formality > 2 {
		return nonValidFormality
	}
	if *language != "english" || *language != "spanish" || *language != "russian" {
		return nonValidLanguage
	}
}

func errorCheck(e error) {
	if e != nil {
		log.Fatalf("Error: %w", e)
	}
}