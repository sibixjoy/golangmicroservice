package main

import (
	"net/http"

	"example.com/order/orderdb"

	"github.com/gorilla/mux"
)

func main() {
	var ordhandlerobj orderdb.OrdHandler

	ordhandlerobj.Connection("localhost","postgres","1234","orderdb","9000")

	router:=mux.NewRouter()
	router.HandleFunc("/health", orderdb.HealthCheck).Methods("GET")
	router.HandleFunc("/order", ordhandlerobj.GetOrder).Methods("GET")
	router.HandleFunc("/addorder", ordhandlerobj.AddOrder).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8300", router)

	dbinstance,_ := ordhandlerobj.DB.DB()
	defer dbinstance.Close()



}