//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi home, I love %s!", r.URL.Path[1:])
}

func query(){

}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/home/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}