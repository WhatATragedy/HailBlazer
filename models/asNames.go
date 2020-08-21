package models
import (
	"io"
	"encoding/json"
	"fmt"
	"database/sql"
	//what is the _ here
	_ "github.com/lib/pq"
)

type ASData struct {
	Name string	`dbname:"name"`
	Country string `dbname:"country"`
	ASN string		`dbname:"asn"`
}
// AutonomousSystems is a collection of ASData
type AutonomousSystems []*ASData

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "user"
	dbname   = "routing_information"
)

func (as *ASData) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(as)
}

func connectDB() (*sql.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		  "password=%s dbname=%s sslmode=disable",
		  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
		return nil, err
	}
	return db, nil
}
// GetAutonomousSystems returns a list of Autonomous Systems
func GetAutonomousSystems() (AutonomousSystems, error) {
	var AutonomousSystemsList []*ASData
	//TODO move the select * logic here
	db, err := connectDB()
	defer db.Close()
	rows, err := db.Query("SELECT asn, name, country FROM as_names;")
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var asn ASData
		// & character returns the memory address of the following variable.
		err := rows.Scan(&asn.ASN, &asn.Name, &asn.Country) // decode similar to deserialize process.
		if err != nil {
			panic(err)
			return nil, err
		}
		// add item our array
		AutonomousSystemsList = append(AutonomousSystemsList, &asn)
	}
	return AutonomousSystemsList, nil
}
func GetAutonomousSystemName(name string) (AutonomousSystems, error) {
	fmt.Println("Handle GET Autonomous System Name")
	var AutonomousSystemsList []*ASData
	//TODO move the select * logic here
	db, err := connectDB()
	defer db.Close()
	name = "%" + name + "%"
	sqlStatement := "SELECT asn, name, country FROM as_names WHERE name LIKE $1;"
	rows, err := db.Query(sqlStatement, name)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var asn ASData
		// & character returns the memory address of the following variable.
		err := rows.Scan(&asn.ASN, &asn.Name, &asn.Country) // decode similar to deserialize process.
		if err != nil {
			panic(err)
			return nil, err
		}
		// add item our array
		AutonomousSystemsList = append(AutonomousSystemsList, &asn)
	}
	return AutonomousSystemsList, nil
}
func GetAutonomousSystemCountry(country string) (AutonomousSystems, error) {
	fmt.Println("Handle GET Autonomous System Country")
	var AutonomousSystemsList []*ASData
	//TODO move the select * logic here
	db, err := connectDB()
	defer db.Close()
	sqlStatement := `SELECT asn, name, country FROM as_names WHERE country=$1;`
	rows, err := db.Query(sqlStatement, country)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var asn ASData
		// & character returns the memory address of the following variable.
		err := rows.Scan(&asn.ASN, &asn.Name, &asn.Country) // decode similar to deserialize process.
		if err != nil {
			panic(err)
			return nil, err
		}
		// add item our array
		AutonomousSystemsList = append(AutonomousSystemsList, &asn)
	}
	return AutonomousSystemsList, nil
}
func GetAutonomousSystemNumber(number int) (ASData, error) {
	var asn ASData
	db, err := connectDB()
	if err != nil {
		panic(err)
		return asn, err
	}
	defer db.Close()
	sqlStatement := `SELECT asn, name, country FROM as_names WHERE asn=$1;`
	row := db.QueryRow(sqlStatement, number)
	switch err := row.Scan(&asn.ASN, &asn.Name, &asn.Country); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return asn, err
	case nil:
		return asn, err
	default:
		panic(err)
	}
	return asn, nil
}



// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (asns *AutonomousSystems) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(asns)
}
func (asn *ASData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(asn)
}

var ErrASNNotFound = fmt.Errorf("Autonomous System not found")

func findASN(id int) (*ASData, int, error) {
	//TODO Move select asn logic here
	return nil, -1, ErrASNNotFound
}
