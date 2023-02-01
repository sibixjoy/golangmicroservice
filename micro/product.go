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
type Product struct {

	Id int `json:"ID"`
	Product string `json:"name"`
	Price int `json:"price"`
}
var db *gorm.DB 

func main() {
	var err error
   
	dsn := "host=localhost user=postgres password=1234 dbname=micro port=9000 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	 panic("failed to connect to database")
	}
   productdb,_:=db.DB()
	   fmt.Println("Connection opened to database")
   
   db.AutoMigrate(&Product{})
   
	router := mux.NewRouter()
   
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/product", product).Methods("GET")
	router.HandleFunc("/addproduct", Addproduct).Methods("POST")
	router.HandleFunc("/getproduct/{id}", Getproduct).Methods("GET")


	http.ListenAndServe(":8002", router)
   
	defer productdb.Close() 
}	
   
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Area")
}

func product(w http.ResponseWriter, r *http.Request) {

    var products []Product
	db.Find(&products)

	json.NewEncoder(w).Encode(products)
}   

func Addproduct(w http.ResponseWriter, r *http.Request) {

	var newproduct Product

	reqBody, _:= ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &newproduct)
	db.Create(&newproduct)

	json.NewEncoder(w).Encode(newproduct)

}

func Getproduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	productID := mux.Vars(r)["id"]

		 db.First(&product, productID) 
			json.NewEncoder(w).Encode(product)
		
}

