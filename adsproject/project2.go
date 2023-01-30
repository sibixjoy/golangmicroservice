package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Product struct {
	gorm.Model
	Id int `json:"ID"`
	Product string `json:"name"`
	Price int `json:"price"`
}
type Producthandler struct {
	DB *gorm.DB
}
func main() {
	var producthandlerObj Producthandler
	producthandlerObj.Connect("localhost", "postgres", "1234", "order", "9000")

	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/product", producthandlerObj.product).Methods("GET")
	router.HandleFunc("/addproduct", producthandlerObj.Addproduct).Methods("POST")
	router.HandleFunc("/getproduct", producthandlerObj.Getproduct).Methods("GET")

	http.ListenAndServe(":8002", router)
	dbinstance,_ := producthandlerObj.DB.DB()
	defer dbinstance.Close()
}
func (producthandler *Producthandler) Connect(host, user, password, dbname, port string) {
	var err error

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	producthandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")

	producthandler.DB.AutoMigrate(&Product{})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Service is up and running")
}
func (producthandler *Producthandler) product(w http.ResponseWriter, r *http.Request) {
	var products []Product
	producthandler.DB.Find(&products)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (producthandler *Producthandler) Addproduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &product)
	
	producthandler.DB.Create(&product)
	json.NewEncoder(w).Encode(product)

}
func (producthandler *Producthandler) Getproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	params := mux.Vars(r)

	producthandler.DB.First(&product, params["id"])

	json.NewEncoder(w).Encode(product)
}