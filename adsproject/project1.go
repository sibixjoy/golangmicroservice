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

type Order struct {
	gorm.Model
	Id int `json:"ID"`
	OrderName string `json:"order_name"`
}

type OrdHandler struct {
	DB *gorm.DB
}

func main() {
	var ordhandlerObj OrdHandler
	ordhandlerObj.Connect("localhost", "postgres", "1234", "order", "9000")

	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/order", ordhandlerObj.GetOrder).Methods("GET")
	router.HandleFunc("/addorder", ordhandlerObj.AddOrder).Methods("POST")

	http.ListenAndServe(":8001", router)
	dbinstance,_ := ordhandlerObj.DB.DB()
	defer dbinstance.Close()
}

func (ordhandler *OrdHandler) Connect(host, user, password, dbname, port string) {
	var err error

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	ordhandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")

	ordhandler.DB.AutoMigrate(&Order{})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Service is up and running")
}

func (ordhandler *OrdHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	ordhandler.DB.Find(&orders)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (ordhandler *OrdHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &order)
	
	ordhandler.DB.Create(&order)
	json.NewEncoder(w).Encode(order)

}

