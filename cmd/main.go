package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Gorgija/go-sofia/internal/diagnostics"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("Server runing ...")

	blPort := os.Getenv("PORT")

	if len(blPort) > 0 {
		log.Fatal("The application port should be set")
	}

	diagPort := os.Getenv("DIAG_PORT")

	if len(diagPort) > 0 {
		log.Fatal("The diagnostics port should be set")
	}

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
