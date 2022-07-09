package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Fprintf(w, "Hello world")
	})
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello golang!"))
	})
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
