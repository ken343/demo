package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	mSpanish[FORMAL] = "¡Buenos días!"
	mSpanish[SEMIFORMAL] = "¡Hola!"
	mSpanish[INFORMAL] = "¿Qué pasa?"

	mRussian[FORMAL] = "Добрый день!"
	mRussian[SEMIFORMAL] = "Здравствуйте!"
	mRussian[INFORMAL] = "Привет!"

}

func main() {
	fmt.Println("Howdy Revature!")

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
		fmt.Printf("Wrote %d bytes in English\n", n)
	})
	// Spanish Handlers
	spanishMux.HandleFunc("/spanish", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formalVal, err := strconv.ParseInt(r.Form.Get("formal"), 10, 64)
		errorCheck(err)

		if formalVal == 0 {
			formalVal = FORMALITYDEFAULT
		}

		greeting := mSpanish[formalVal] + "\n"

		n, err := w.Write([]byte(greeting))
		errorCheck(err)
		fmt.Printf("Wrote %d bytes in Spanish\n", n)
	})

	// Russian Handlers
	russianMux.HandleFunc("/russian", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formalVal, err := strconv.ParseInt(r.Form.Get("formal"), 10, 64)
		errorCheck(err)

		if formalVal == 0 {
			formalVal = FORMALITYDEFAULT
		}

		greeting := mRussian[formalVal] + "\n"

		n, err := w.Write([]byte(greeting))
		errorCheck(err)
		fmt.Printf("Wrote %d bytes in Russian\n", n)
	})

	go func() {
		err := http.ListenAndServe(ENGLISHPORT, englishMux)
		errorCheck(err)
	}()
	fmt.Println("English Greeting listening on port 8081...")

	go func() {
		err := http.ListenAndServe(SPANISHPORT, spanishMux)
		errorCheck(err)
	}()
	fmt.Println("Spanish Greeting listens on port 8082...")

	go func() {
		err := http.ListenAndServe(RUSSIANPORT, russianMux)
		errorCheck(err)
	}()
	fmt.Println("Russian Greeting listens on port 8083...")

	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, syscall.SIGINT)

	for {
		select {
		case sig := <-shutDown:
			fmt.Println("\nShutting down server cluster:", sig)
			os.Exit(0)
		}
	}
}

func errorCheck(e error) {
	if e != nil {
		log.Fatalf("Error: %w", e)
	}
}
