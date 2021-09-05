package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello world")); err != nil {
		log.Println(err)
	}
}