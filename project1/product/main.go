package main

import (
	"net/http"

	"example.com/product/productdb"

	"github.com/gorilla/mux"
)

func main() {
	var prohandlerobj productdb.ProHandler

	prohandlerobj.Connection("localhost","postgres","1234","orderdb","9000")
	
	router:=mux.NewRouter()
	router.HandleFunc("/health", productdb.HealthCheck).Methods("GET")
	router.HandleFunc("/product", prohandlerobj.GetProduct).Methods("GET")
	router.HandleFunc("/product/{id}", prohandlerobj.GetIndProduct).Methods("GET")
	router.HandleFunc("/addproduct", prohandlerobj.AddProduct).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8100", router)

	dbinstance,_ := prohandlerobj.DB.DB()
	defer dbinstance.Close()



}