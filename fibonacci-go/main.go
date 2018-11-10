package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func fibonacciSequence(terms int) []int {
	var termsList []int
	termsList = append(termsList, 0)
	var nextTerm = 1

	for i := 0; i < (terms - 1); i++ {
		termsList = append(termsList, nextTerm)
		nextTerm = (nextTerm + termsList[i])
	}

	return termsList
}

func converter(input string) {
	// would have extracted error logic from FibbonacciAPI and outputted either
	// the sequence or the error message. This way i could test the output.
	// i need to learn how to return what i want from a function with conditional in go, kept getting errors
}

func FibonacciAPI(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var input int
	value, err1 := strconv.ParseFloat(params["num"], 32)
	if err1 == nil {
		// no error in integer conversion
		if value > -1 {
			// positive integer
			if value < 94 {
				// low enought to calc!
				input = int(value)
				var finalSeq = fibonacciSequence(input)
				json.NewEncoder(w).Encode(finalSeq)
			} else {
				// too high to calc
				var errorMsg = "Error: This number is higher than we can calculate. Please go lower"
				json.NewEncoder(w).Encode(errorMsg)
			}
		} else {
			// negative intenger or zero
			var errorMsg = "Error: No negative values please"
			json.NewEncoder(w).Encode(errorMsg)
		}
	} else {
		// input was a something that could not be successfully converted to integer
		var errorMsg = "Error: input was not converted to integer"
		json.NewEncoder(w).Encode(errorMsg)
	}

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/fibonacci/{num}", FibonacciAPI).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
