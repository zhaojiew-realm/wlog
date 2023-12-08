package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", logRequest)
	err := http.ListenAndServe(":8108", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body content:", err)
	} else {
		fmt.Printf("%s\n", body)
	}
	fmt.Fprintln(w, string(body))
}
