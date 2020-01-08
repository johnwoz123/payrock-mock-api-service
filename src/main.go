package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type ResidualVolume struct {
	CardType         string `json:"card_type"`
	InterchangeLevel string `json:"interchange_level"`
	InterchangeRate  string `json:"interchange_rate"`
	TransactionCount string `json:"transaction_count"`
	Volume           string `json:"volume"`
	MerchantRate     string `json:"merchant_rate"`
	GrossVolume      string `json:"gross_volume"`
	Cost             string `json:"cost"`
	Split            string `json:"split"`
	Residual         string `json:"residual"`
}

type ResidualBillables struct {
	Billable    string `json:"billable"`
	Revenue     string `json:"revenue"`
	BuyRate     string `json:"buyrate"`
	Count       string `json:"count"`
	BuyRateCost string `json:"buy_rate_cost"`
	Split       string `json:"split"`
	Residual    string `json:"residual"`
}

type Residual struct {
	RepNumber              string              `json:"rep_number"`
	RepName                string              `json:"rep_name"`
	Mid                    string              `json:"mid"`
	TotalResidualAmount    string              `json:"total_residual_amount"`
	TotalResidualVolume    string              `json:"total_residual_volume"`
	TotalResidualBillables string              `json:"total_residual_billables"`
	ResidualBillables      []ResidualBillables `json:"residual_billables"`
	ResidualVolume         []ResidualVolume    `json:"residual_volume"`
}

func checkError(e error) {
	if e != nil {
		fmt.Printf("ERROR ---- %v", e)
		panic(e)
	}
}

// ReadFile - wrapper around the ioutil ReadFile function
func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	checkError(err)
	return data, nil
}

func GetResidualsHandlers(w http.ResponseWriter, r *http.Request) {

	// read in json file
	data, err := ReadFile("utility/residual.json")
	checkError(err)

	var residual []Residual

	// marshal to struct
	err = json.Unmarshal(data, &residual)
	checkError(err)
	// respond with json
	json.NewEncoder(w).Encode(residual)

}

func main() {
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "application/json"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origns := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/residuals", GetResidualsHandlers).Methods("GET")

	fmt.Println("STARTING SERVER........")
	log.Fatal(http.ListenAndServe(":3005", handlers.CORS(headers, methods, origns)(r)))
}
