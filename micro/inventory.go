package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Inventory struct{
	Id int `json:"ID"`
	ProductName string `json:"product_name"`
	Quantity int `json:"product_quantity"`
}

var db *gorm.DB

func main () {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=micro port=9000 sslmode=disable"
	db, err  = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faild to connect to database")
	}
   inventorydb,_:=db.DB()
	   fmt.Println("Connection opened to database")
	   
   db.AutoMigrate(&Inventory{})

    router := mux.NewRouter()

    router.HandleFunc("/healthcheck",HealthCheck).Methods("GET")
    router.HandleFunc("/inventory", inventory) .Methods("GET")
    router.HandleFunc("/addinventory",Addinventory) .Methods("POST")
    router.HandleFunc("/getinventory/{id}",Getinventory) .Methods("GET")
    router.HandleFunc("/deleteinventory/{id}",Delinventory) .Methods("DELETE")

    http.ListenAndServe(":8003", router)

    defer inventorydb.Close()
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Area")
}

func inventory(w http.ResponseWriter, r *http.Request) {

    var inventorys []Inventory
	db.Find(&inventorys)

	json.NewEncoder(w).Encode(inventorys)

}

func Addinventory(w http.ResponseWriter, r *http.Request) {

	var newinventory Inventory

	reqBody, _:= ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &newinventory)
	db.Create(&newinventory)

	json.NewEncoder(w).Encode(newinventory)

}

func Getinventory(w http.ResponseWriter, r *http.Request) {
	var inventory Inventory
	inventoryID := mux.Vars(r)["id"]

		 db.First(&inventory, inventoryID) 
			json.NewEncoder(w).Encode(inventory)
		
}

func Delinventory(w http.ResponseWriter, r *http.Request) {
	var delinventory Inventory
	inventoryID := mux.Vars(r)["id"]

		 db.Delete(&delinventory, inventoryID) 
			json.NewEncoder(w).Encode(delinventory)
		
}

