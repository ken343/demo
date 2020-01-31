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
	ENGLISHPORT = ":8081" // Port number for english server
	SPANISHPORT = ":8082" // Port number for spanish server
	RUSSIANPORT = ":8083" // Port number for russian server
	FORMAL = 0 // Retrieves back most formal response from greeting map.
	SEMIFORMAL = 1 // Retrieves a moderately formal response from greeting map.
	INFORMAL = 2 // Retrieves an informal response from greeting map.
)

var (
	var nonValidPort = errors.New("Port Number not Valid")
	var nonValidFormality = errors.New("Formaility value not valid")
	var nonValidLanguage = errors.New("Non-Valid Language")
)

var pPort *int64
var pFormality *int64

var mEnglish = make(map[int64]string)
var mSpanish = make(map[int64]string)
var mRussian = make(map[int64]string)

func init() {
	pPort = flag.Int64("port", PORTDEFAULT, "Specify the port the server should be listening on. 8080 by default.")
	pFormality = flag.Int64("formality", FORMALITYDEFAULT, "Formal Greeting==0, Semi-Formal Greeting==1, Informal Greeting==2")
	pLanguage = flag.String("lang", LANGUAGEDEFAULT, "Options for Greeting: 'english', 'spanish', 'russian'. English by default.")
	flag.Parse()

	err := checkFlags(pPort, pFormality, pLanguage)
	errorCheck(err)

	mEnglish[FORMAL] = "Good Day!"
	mEnglish[SEMIFORMAL] = "Hello!"
	mEnglish[INFORMAL] = "Hey!"

	mSpanish[FORMAL] = "Buenos Dias!"
	mSpanish[SEMIFORMAL] = "Hola!"
	mSpanish[INFORMAL] = "Que tal?"

	mRussian[FORMAL] = "Dobre Den"
	mRussian[SEMIFORMAL] = "Zdrastvutye"
	mRussian[INFORMAL] = "Preevet"
	
}

func main() {
	fmt.Println("Howdy")

	// Setup Server Multiplexers
	englishMux := http.NewServeMux()
	spansihMux := http.NewServeMux()
	russianMux := http.NewServeMux()

	// English Handlers

	// Spanish Handlers
	
	// Russian Handlers
	
	go func() {
		err := http.ListenAndServe(ENGLISHPORT, englishMux)
		errorCheck(err)
	}()
	fmt.Println("English Greeting listening on port 8081")

	go func() {
		err := http.ListenAndServe(SPANISHPORT, spanishMux)
		errorCheck(err)
	}
	fmt.Println("Spanish Greeting listens on port 8082")

	err := http.ListenAndServe(RUSSIANPORT, russianMux)
	errorCheck(err)
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