package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"strconv"
	//"context"
	//"time"

	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
	//what is the _ here
  
)
//TODO Create a model directory and move structs into there
//TODO Create a controller to move all the methods there
//TODO creeate an env file to load in consts

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "user"
	dbname   = "routing_information"
)
  
type ASNData struct {
	Name string	`dbname:"name"`
	Country string `dbname:"country"`
	ASN string		`dbname:"asn"`
}
func getASNs(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")
	// we get params with mux.
	//var params = mux.Vars(r)
	var asns []ASNData
	//id, _ := primitive.ObjectIDFromHex(params["id"])
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		  "password=%s dbname=%s sslmode=disable",
		  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
  		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT asn, name, country FROM as_names;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows.Columns)
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var asn ASNData
		// & character returns the memory address of the following variable.
		err := rows.Scan(&asn.ASN, &asn.Name, &asn.Country) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		asns = append(asns, asn)
	}
	json.NewEncoder(w).Encode(asns) // encode similar to serialize process.
}

func getASN(w http.ResponseWriter, r *http.Request){
		// set header.
		w.Header().Set("Content-Type", "application/json")
		// we get params with mux.
		var params = mux.Vars(r)
		var asn ASNData
		id, _ := strconv.Atoi(params["asn"])
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			  "password=%s dbname=%s sslmode=disable",
			  host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			  panic(err)
		}
		defer db.Close()
		sqlStatement := `SELECT asn, name, country FROM as_names WHERE asn=$1;`
		row := db.QueryRow(sqlStatement, id)
		switch err := row.Scan(&asn.ASN, &asn.Name, &asn.Country); err {
		case sql.ErrNoRows:
		  fmt.Println("No rows were returned!")
		case nil:
		  fmt.Println(id)
		default:
		  panic(err)
		}		
		json.NewEncoder(w).Encode(asn) // encode similar to serialize process.
}
func getCountry(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")
	// we get params with mux.
	var params = mux.Vars(r)
	var asns []ASNData
	country, _ := params["country"]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		  "password=%s dbname=%s sslmode=disable",
		  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
  		panic(err)
	}
	defer db.Close()
	sqlStatement := `SELECT asn, name, country FROM as_names WHERE country=$1;`
	rows, err := db.Query(sqlStatement, country)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var asn ASNData
		// & character returns the memory address of the following variable.
		err := rows.Scan(&asn.ASN, &asn.Name, &asn.Country) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}
		// add item our array
		asns = append(asns, asn)
	}
	json.NewEncoder(w).Encode(asns) // encode similar to serialize process.
}
func getName(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")
	// we get params with mux.
	var params = mux.Vars(r)
	var asns []ASNData
	country, _ := params["name"]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		  "password=%s dbname=%s sslmode=disable",
		  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
  		panic(err)
	}
	defer db.Close()
	sqlStatement := `SELECT asn, name, country FROM as_names WHERE name LIKE $1;`
	rows, err := db.Query(sqlStatement, country)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var asn ASNData
		// & character returns the memory address of the following variable.
		err := rows.Scan(&asn.ASN, &asn.Name, &asn.Country) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}
		// add item our array
		asns = append(asns, asn)
	}
	json.NewEncoder(w).Encode(asns) // encode similar to serialize process.
}


// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/api/asns", getASNs).Methods("GET")
	r.HandleFunc("/api/as_name/{name}", getName).Methods("GET")
	r.HandleFunc("/api/asn/{asn}", getASN).Methods("GET")
	r.HandleFunc("/api/country/{country}", getCountry).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}