package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// These values represent the default values that the flags should get
// if not specified when running the program.
const (
	PORTDEFAULT      = 8080      // Default port number if none is specified.
	FORMALITYDEFAULT = 0         // Defaults to most formal greeting if not specified.
	LANGUAGEDEFAULT  = "english" // Defaults to english if not specified.
	ENGLISHPORT      = ":8081"   // Port number for english server
	SPANISHPORT      = ":8082"   // Port number for spanish server
	RUSSIANPORT      = ":8083"   // Port number for russian server
	FORMAL           = 0         // Retrieves back most formal response from greeting map.
	SEMIFORMAL       = 1         // Retrieves a moderately formal response from greeting map.
	INFORMAL         = 2         // Retrieves an informal response from greeting map.
)

var mEnglish = make(map[int64]string)
var mSpanish = make(map[int64]string)
var mRussian = make(map[int64]string)

func init() {
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
	spanishMux := http.NewServeMux()
	russianMux := http.NewServeMux()

	// English Handlers
	englishMux.HandleFunc("/english", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formalVal, err := strconv.ParseInt(r.Form.Get("formal"), 10, 64)
		errorCheck(err)

		if formalVal == 0 {
			formalVal = FORMALITYDEFAULT
		}

		greeting := mEnglish[formalVal] + "\n"

		n, err := w.Write([]byte(greeting))
		errorCheck(err)
		fmt.Printf("Wrote %d bytes\n", n)
	})
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
	}()
	fmt.Println("Spanish Greeting listens on port 8082")

	err := http.ListenAndServe(RUSSIANPORT, russianMux)
	errorCheck(err)
	fmt.Println("Russian Greeting listens on port 8083")
}

func errorCheck(e error) {
	if e != nil {
		log.Fatalf("Error: %w", e)
	}
}
