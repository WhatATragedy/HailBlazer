package handlers
import (
	"net/http"
	"log"
	"github.com/alex53856/RouteControlMapREST/models"
)
type ASNames struct {
	l *log.Logger
}

func NewAS(l *log.Logger) *ASNames {
	return &ASNames{l}
}
// getProducts returns the products from the data store
func (asn *ASNames) getAutonomousSystems(rw http.ResponseWriter, r *http.Request) {
	asn.l.Println("Handle GET Autonomous Systems")

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

func (asn *ASNames) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	asn.l.Println("Serving All ASNs")
	if r.Method == http.MethodGet {
		rw.Header().Set("Content-Type", "application/json")
		asn.getAutonomousSystems(rw, r)
	} else {
		asn.l.Println("Method Other than GET called on ASN..")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
	
}