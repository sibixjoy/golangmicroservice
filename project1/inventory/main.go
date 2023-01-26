package main

import (
	"net/http"

	"example.com/inventory/inventorydb"

	"github.com/gorilla/mux"
)

func main() {
	var invhandlerobj inventorydb.InvHandler

	invhandlerobj.Connection("localhost","postgres","1234","orderdb","9000")
	
	router:=mux.NewRouter()
	router.HandleFunc("/health", inventorydb.HealthCheck).Methods("GET")
	router.HandleFunc("/inventory", invhandlerobj.GetInventory).Methods("GET")
	router.HandleFunc("/addinventory", invhandlerobj.AddInventory).Methods("POST")
	router.HandleFunc("/singleinventory/{id}", invhandlerobj.GetIndInventory).Methods("GET")
	router.HandleFunc("/delinventory/{id}", invhandlerobj.DelInventory).Methods("DELETE")


	http.Handle("/", router)
	http.ListenAndServe(":8200", router)

	dbinstance,_ := invhandlerobj.DB.DB()
	defer dbinstance.Close()



}