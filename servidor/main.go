package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var results []string

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test Webhook")
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body",
					http.StatusInternalServerError)
			}
			results = append(results, string(body))

			fmt.Fprint(w, "POST done")
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})

	http.ListenAndServe(":8181", nil)

}
