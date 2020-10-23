package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	//what is the _ here
	"HailBlazer/handlers"
)

//TODO Create a model directory and move structs into there
//TODO Create a controller to move all the methods there
//TODO creeate an env file to load in consts
// Main function
func main() {
	l := log.New(os.Stdout, "Hail-Blazer ", log.LstdFlags)
	asNamesHandler := handlers.NewAS(l)
	talHandler := handlers.NewTal(l)

	// Init router
	r := mux.NewRouter()
	asNamesRouter := r.PathPrefix("/api/as_whois").Subrouter()
	asNamesRouter.HandleFunc("/asns", asNamesHandler.GetAutonomousSystems)
	asNamesRouter.HandleFunc("/as_name/{name}", asNamesHandler.GetAutonomousSystemName)
	asNamesRouter.HandleFunc("/country/{country}", asNamesHandler.GetAutonomousSystemCountry)
	asNamesRouter.HandleFunc("/asn/{asn:[0-9]+}", asNamesHandler.GetAutonomousSystemNumber)
	talRouter := r.PathPrefix("/api/tals").Subrouter()
	talRouter.HandleFunc("/", talHandler.GetTals)
	talRouter.HandleFunc("/prefix/{IP}", talHandler.GetTalIP)

	// Route handles & endpoints
	//r.HandleFunc("/api/asns", asNamesHandler).Methods("GET")
	//r.HandleFunc("/api/as_name/{name}", getName).Methods("GET")
	//r.HandleFunc("/api/asn/{asn}", getASN).Methods("GET")
	//r.HandleFunc("/api/country/{country}", getCountry).Methods("GET")

	// Configure Server Shizzle
	s := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		IdleTimeout:  20 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// Start server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
