package main

import (
	"log"
	"net/http"
	"context"
	"time"
	"os"
	"os/signal"

	//"github.com/gorilla/mux"
	//what is the _ here
	"github.com/alex53856/RouteControlMapREST/handlers"
  
)
//TODO Create a model directory and move structs into there
//TODO Create a controller to move all the methods there
//TODO creeate an env file to load in consts
// Main function
func main() {
	l := log.New(os.Stdout, "Hail-Blazer ", log.LstdFlags)
	asNamesHandler := handlers.NewAS(l)

	sm := http.NewServeMux()
	sm.Handle("/", asNamesHandler)

	// Init router
	//r := mux.NewRouter()
	
	// Route handles & endpoints
	//r.HandleFunc("/api/asns", asNamesHandler).Methods("GET")
	//r.HandleFunc("/api/as_name/{name}", getName).Methods("GET")
	//r.HandleFunc("/api/asn/{asn}", getASN).Methods("GET")
	//r.HandleFunc("/api/country/{country}", getCountry).Methods("GET")

	// Configure Server Shizzle
	s := &http.Server{
		Addr:":9090",
		Handler: sm,
		IdleTimeout: 20*time.Second,
		ReadTimeout: 3*time.Second,
		WriteTimeout: 30*time.Second,
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