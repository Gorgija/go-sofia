package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gorgija/go-sofia/internal/diagnostics"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("Server runing ...")
	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatal(err)
		}

	}()
	diagnostics := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":8585", diagnostics)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
