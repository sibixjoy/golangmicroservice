package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	
	Id int `json:"ID"`
	OrderName string `json:"order_name"`
}

var db *gorm.DB 

func main() {
 var err error

 dsn := "host=localhost user=postgres password=1234 dbname=order port=9000 sslmode=disable"
 db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
 if err != nil {
  panic("failed to connect to database")
 }
orderdb,_:=db.DB()
	fmt.Println("Connection opened to database")

db.AutoMigrate(&Order{})

 router := mux.NewRouter()

 router.HandleFunc("/health", HealthCheck).Methods("GET")
 router.HandleFunc("/order", GetOrder).Methods("GET")
 router.HandleFunc("/addorder", AddOrder).Methods("POST")

 http.ListenAndServe(":8001", router)
 fmt.Println("server is running on 8001 port")

 defer orderdb.Close() 
}	

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Area")
   }

func GetOrder(w http.ResponseWriter, r *http.Request) {

	var orders []Order
	db.Find(&orders)

	json.NewEncoder(w).Encode(orders)
}   
   
func AddOrder(w http.ResponseWriter, r *http.Request) {

	var neworder Order

	reqBody, _:= ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &neworder)
	db.Create(&neworder)

	json.NewEncoder(w).Encode(neworder)

}

