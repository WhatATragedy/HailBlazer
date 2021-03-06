package handlers

import (
	"HailBlazer/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Tal struct {
	l *log.Logger
}

func NewTal(l *log.Logger) *Tal {
	return &Tal{l}
}

// getProducts returns the products from the data store
func (tal *Tal) GetTals(rw http.ResponseWriter, r *http.Request) {
	tal.l.Println("Handle GET All Tals")

	// fetch the products from the datastore
	tals, err := models.GetTals()
	for _, i := range tals {
		fmt.Println(i.Prefix, i.AutonomousSystem, i.ValidFrom, i.SourceRIR, i.SourceDate)

	}
	if err != nil {
		tal.l.Println("Error Getting List of Tals")
	}

	// serialize the list to JSON and write it to the response writer
	err = tals.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (tal *Tal) GetTalIP(rw http.ResponseWriter, r *http.Request) {
	tal.l.Println("Handle GET Tal Prefix")
	vars := mux.Vars(r)
	IP := vars["IP"]
	tal.l.Printf("Handle GET Tals For Prefix/Address %v\n", IP)
	tals, err := models.GetTalIP(IP)
	if err != nil {
		tal.l.Println("Error Gettng Tal Prefix")
	}
	// serialize the list to JSON and write it to the response writer
	err = tals.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (tal *Tal) GetTalASN(rw http.ResponseWriter, r *http.Request) {
	tal.l.Println("Handle GET Tal ASN")
	vars := mux.Vars(r)
	ASNStr := vars["ASN"]
	ASN, err := strconv.ParseInt(ASNStr, 10, 64)
	if err != nil {
		fmt.Println(ASNStr, "is not an integer.")
		http.Error(rw, "Argument is not an integer", http.StatusInternalServerError)
	}
	tal.l.Printf("Handle GET Tals For ASN %v\n", ASN)
	tals, err := models.GetTalASN(ASN)
	if err != nil {
		tal.l.Println("Error Gettng Tal Prefix")
	}
	// serialize the list to JSON and write it to the response writer
	err = tals.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (tal *Tal) GetTalRIR(rw http.ResponseWriter, r *http.Request) {
	tal.l.Println("Handle GET Tal RIR")
	vars := mux.Vars(r)
	RIR := vars["RIR"]
	tal.l.Printf("Handle GET Tals For RIR %v\n", RIR)
	tals, err := models.GetTalRIR(RIR)
	if err != nil {
		tal.l.Println("Error Gettng Tal Prefix")
	}
	// serialize the list to JSON and write it to the response writer
	err = tals.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
