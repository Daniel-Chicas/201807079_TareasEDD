package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"./listas"
	"github.com/gorilla/mux"
	"log"
	"net/http"

)

var ins listas.Lista
var ms listas.Datos
var nodo listas.Nodo

func inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "INICIO")
}

func insertar(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Fprint(w, "Error al insertar")
	}
	json.Unmarshal(reqBody, &ms)
	for i := 0; i < len(ms.General); i++ {
		a := ms.General[i]
		nodo.OrigenM = a.Origen
		nodo.DestinoM = a.Destino
		for j := 0; j < len(a.Msg); j++ {
			b := a.Msg[j]
			nodo.CuerpoM = b.Texto
			nuevo := listas.Nodo{OrigenM: a.Origen, DestinoM: a.Destino, CuerpoM: b.Texto}
			fmt.Fprint(w, ins.Insertar(&nuevo))
			fmt.Fprint(w, nodo.RetornarNodo())
		}
	}
}

type Mensaje struct{
	Origen string `json:"Origen"`
	Destino string `json:"Destino"`
	Cuerpo string `json:"Texto"`
}

func mensajes(w http.ResponseWriter, r *http.Request){
	var aux string
	for i := 0; i < len(ms.General); i++ {
		a := ms.General[i]
		for j := 0; j < len(a.Msg); j++ {
			b := a.Msg[j]
			aux = b.Texto
			a := Mensaje{a.Origen,a.Destino, aux}
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(a)
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicio).Methods("Get")
	router.HandleFunc("/insertar", insertar).Methods("Post")
	router.HandleFunc("/mensajes", mensajes).Methods("Get")
	log.Fatal(http.ListenAndServe(":3000", router))
}
