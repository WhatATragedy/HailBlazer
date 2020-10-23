package models

import (
	"encoding/json"
	"fmt"
	"io"

	//what is the _ here
	_ "github.com/lib/pq"
)

type Tal struct {
	Prefix           string `dbname:"prefix"`
	AutonomousSystem string `dbname:"asn"`
	ValidFrom        string `dbname:"ValidFrom"`
	SourceRIR        string `dbname:"SourceRIR"`
	SourceDate       string `dbname:"SourceDate"`
}

// AutonomousSystems is a collection of ASData
type Tals []*Tal

func (tal *Tal) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(tal)
}

// GetAutonomousSystems returns a list of Autonomous Systems
func GetTals() (Tals, error) {
	var talList []*Tal
	//TODO move the select * logic here
	db, err := connectDB()
	defer db.Close()
	rows, err := db.Query("SELECT prefix, asn, ValidFrom, SourceRIR, SourceDate FROM tals;")
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var tal Tal
		// & character returns the memory address of the following variable.
		err := rows.Scan(&tal.Prefix, &tal.AutonomousSystem, &tal.ValidFrom, &tal.SourceRIR, &tal.SourceDate) // decode similar to deserialize process.
		if err != nil {
			panic(err)
			return nil, err
		}
		fmt.Println(tal.Prefix, tal.AutonomousSystem, tal.ValidFrom, tal.SourceRIR, tal.SourceDate)
		// add item our array
		talList = append(talList, &tal)
	}

	return talList, nil
}
func GetTalIP(IP string) (Tals, error) {
	var talList []*Tal
	//TODO move the select * logic here
	db, err := connectDB()
	defer db.Close()
	sqlStatement := "SELECT prefix, asn, ValidFrom, SourceRIR, SourceDate FROM tals WHERE prefix >>= $1;"
	rows, err := db.Query(sqlStatement, IP)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// create a value into which the single document can be decoded
		var tal Tal
		// & character returns the memory address of the following variable.
		err := rows.Scan(&tal.Prefix, &tal.AutonomousSystem, &tal.ValidFrom, &tal.SourceRIR, &tal.SourceDate) // decode similar to deserialize process.
		if err != nil {
			panic(err)
			return nil, err
		}
		// add item our array
		talList = append(talList, &tal)
	}
	return talList, nil
}
func (tals *Tals) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(tals)
}
