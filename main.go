package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the root")
	})
	http.HandleFunc("/123", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the root123")
	})
	http.HandleFunc("/api/math24", func(w http.ResponseWriter, r *http.Request) {
		type Input struct {
			Num1 int `json:"num1"`
			Num2 int `json:"num2"`
			Num3 int `json:"num3"`
			Num4 int `json:"num4"`
		}
		var input Input

		// Decode json object
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&input)
		// Calculate a sum
		sum := input.Num1 + input.Num2 + input.Num3 + input.Num4
		fmt.Printf("Sum: %v", sum)

		// Declare answer object
		type Answer struct {
			Sum int
		}
		var answer Answer
		// Set calculated sum into answer.Sum
		answer.Sum = sum

		// Marshal answer object for returning
		js, _ := json.Marshal(answer)

		// Send back json answer
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", nil)
}
