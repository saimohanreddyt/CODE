


package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "whoa, Go is neat!")	// Fprint is gonna format whatever you specified which we;re onna say is the writer itself and then it's got output whatever you asked
}
func about_handler(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "WELCOME TO THE OUR WEBSITE")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
