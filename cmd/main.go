package main

import (
	"flag"
	routes "github.com/b-entangled/alphabet-api/pkg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	var (
		serviceAddr string
		servicePort string
		apiEndpoint string
	)

	// Initialize Flags
	flag.StringVar(&serviceAddr, "ServiceAddr", "0.0.0.0", "Service Address")
	flag.StringVar(&servicePort, "ServicePort", "8080", "Service Port")
	flag.StringVar(&apiEndpoint, "ApiEndpoint", "/alphabet", "Alphabet API")
	flag.Parse()

	// Initialize Router
	mainRouter := mux.NewRouter()
	// Create NewAlphabetRouter and attach it to main router
	alphabetRouter := routes.NewAlphabetRouter(mainRouter, apiEndpoint)
	mainRouter.Handle(apiEndpoint, alphabetRouter)

	// Create Service
	srv := &http.Server{
		Addr:    serviceAddr + ":" + servicePort,
		Handler: mainRouter,
	}
	// Run Service
	go func() {
		log.Println("Running Server at Port : ", servicePort)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for end signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("shutting down server")
	os.Exit(0)
}
