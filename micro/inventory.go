package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
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
	router.HandleFunc("/getinventoryqty/{id}/{Quantity}",GetinventoryQuantity) .Methods("GET")
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
type Response struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}	


func GetinventoryQuantity(w http.ResponseWriter, r *http.Request) {
	var inventory Inventory
	//var inventoryQty Inventory
	//func getFieldInteger(e *Inventory, field string) int {
		//r := reflect.ValueOf(e)
		//f := reflect.Indirect(r).FieldByName(field)
		//return int(f.Int())
		
	//}
	inventoryqty := mux.Vars(r)["Quantity"]

	qty, _ := strconv.Atoi(inventoryqty)


	inventoryID := mux.Vars(r)["id"]

	    db.First(&inventory, inventoryID) 
        invent:=inventory 
	fmt.Println("body is printing ", (invent))

	//p:=invent

		//if inventoryqty==(p.Quantity ){
				//	fmt.Println("body is nyc")p:=inventory{}
	p:=invent
	
	inty, _ := strconv.Atoi(inventoryID)

	
	if inty==(p.Id) {
		

		if qty<=(p.Quantity ){

			resp := Response{Result: true, Message: "stock is available "}
			jsonResponse, _ := json.Marshal(resp)
			w.Write(jsonResponse)
					//fmt.Println("stock is available ")

		}else{

			resp := Response{Result: false, Message: "Out of Stock Only haves"}
			jsonResponse, _ := json.Marshal(resp)
			w.Write(jsonResponse)
       	//fmt.Println("Out of Stock Only haves ",(p.Quantity ))
        }
	}else{

		resp := Response{Result: false, Message: "Out of Stock SORRY"}
		jsonResponse, _ := json.Marshal(resp)
	    w.Write(jsonResponse)

		//fmt.Println("Out of Stock SORRY")

	}

}	