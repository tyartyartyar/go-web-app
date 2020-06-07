package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Tyar")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
