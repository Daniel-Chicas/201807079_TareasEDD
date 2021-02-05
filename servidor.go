package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "INICIO")

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicio).Methods("Get")
	log.Fatal(http.ListenAndServe(":3000", router))
}
