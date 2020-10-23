package handlers
import (
	"net/http"
	"log"
	"HailBlazer/models"
	"github.com/gorilla/mux"
	"strconv"
)
type ASNames struct {
	l *log.Logger
}

func NewAS(l *log.Logger) *ASNames {
	return &ASNames{l}
}
// getProducts returns the products from the data store
func (asn *ASNames) GetAutonomousSystems(rw http.ResponseWriter, r *http.Request) {
	asn.l.Println("Handle GET All Autonomous Systems")

	// fetch the products from the datastore
	autonomousSystems, err := models.GetAutonomousSystems()
	if err != nil {
		asn.l.Println("Error Getting List of Autonomous Systems")
	}

	// serialize the list to JSON and write it to the response writer
	err = autonomousSystems.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (asn *ASNames) GetAutonomousSystemName(rw http.ResponseWriter, r *http.Request) {
	asn.l.Println("Handle GET Autonomous System Name")
	vars := mux.Vars(r)
	name := vars["name"]
	asn.l.Printf("Handle GET Autonomous System Country For Name %v\n", name)
	autonomousSystem, err := models.GetAutonomousSystemName(name)
	if err != nil {
		asn.l.Println("Error Getting List of Autonomous Systems from name query")
		return
	}
	// serialize the list to JSON and write it to the response writer
	err = autonomousSystem.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (asn *ASNames) GetAutonomousSystemCountry(rw http.ResponseWriter, r *http.Request) {
	asn.l.Println("Handle GET Autonomous System Country")
	vars := mux.Vars(r)
	country := vars["country"]
	asn.l.Printf("Handle GET Autonomous System Country For Country %v\n", country)
	autonomousSystem, err := models.GetAutonomousSystemCountry(country)
	if err != nil {
		asn.l.Println("Error Getting List of Autonomous Systems from country query")
		return
	}
	// serialize the list to JSON and write it to the response writer
	err = autonomousSystem.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	
}
func (asn *ASNames) GetAutonomousSystemNumber(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	asNumber, err := strconv.Atoi(vars["asn"])
	if err != nil {
		asn.l.Println("Error Converting ASN to Int")
		return
	}
	asn.l.Printf("Handle GET Autonomous System Number For AS Number %v\n", asNumber)
	autonomousSystem, err := models.GetAutonomousSystemNumber(asNumber)
	if err != nil {
		asn.l.Println("Error Getting List of Autonomous Systems from number query")
		return
	}
	// serialize the list to JSON and write it to the response writer
	err = autonomousSystem.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}
func (asn *ASNames) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	asn.l.Println("Serving All ASNs")
	if r.Method == http.MethodGet {
		rw.Header().Set("Content-Type", "application/json")
		asn.GetAutonomousSystems(rw, r)
	} else {
		asn.l.Println("Method Other than GET called on ASN..")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
	
}