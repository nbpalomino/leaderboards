package main

import (
	"fmt"
	_ "log"
	"net/http"
)

type Persona struct{
	Nombres string
	Apellidos string
	Dni string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bienvenidos</h1><p>A mi web.</p>")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}