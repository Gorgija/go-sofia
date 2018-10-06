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

	blPort := os.Getenv("APP_PORT")

	if len(blPort) == 0 {
		log.Fatal("The application port should be set")
	}

	diagPort := os.Getenv("DIAG_PORT")

	if len(diagPort) == 0 {
		log.Fatal("The diagnostics port should be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	possibleErrors := make(chan error, 2)

	go func() {
		log.Print("Application server is listyening ....")

		server := &http.Server{
			Addr:    ":" + blPort,
			Handler: router,
		}

		err := server.ListenAndServe()

		if err != nil {
			possibleErrors <- err
		}

	}()

	go func() {

		diagnostics := diagnostics.NewDiagnostics()
		log.Print("Diagnostics server is listyening ....")
		err := http.ListenAndServe(":"+diagPort, diagnostics)
		if err != nil {
			log.Fatal(err)
		}

	}()

	select {
	case err := <-possibleErrors:
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
