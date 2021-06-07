package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/apex/gateway"
	"github.com/go-chi/chi"
)

const defaultPort = "7010"

var router *chi.Mux
var err error

func init() {
	router, err = MakeRouter()
	if err != nil {
		panic("Router Error")
	}
}

func main() {
	isRunningAtLambda := strings.Contains(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda_")

	if isRunningAtLambda {
		log.Fatal(gateway.ListenAndServe(":3000", router))
	} else {
		port := os.Getenv("PORT")

		if port == "" {
			port = defaultPort
		}
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}
