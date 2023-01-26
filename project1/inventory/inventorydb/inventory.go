package inventorydb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Id int `json:"ID"`
	ProductName string `json:"product_name"`
	ProductQuantity int `json:"product_quantity"`
}

type InvHandler struct {
	DB *gorm.DB
}

func (invhandler *InvHandler) Connection(host,user,password,dbname,port string) {
	var err error

	dsn:="host="+host+" user="+user+" password="+password+" dbname="+dbname+" port="+port+" sslmode=disable"
	invhandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// invhandler.DB.AutoMigrate(Inventory{})
	invhandler.DB.AutoMigrate(&Inventory{})

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Super Secret Area")
}


func (invhandler *InvHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inventory []Inventory
	invhandler.DB.Find(&inventory)
	json.NewEncoder(w).Encode(&inventory)
}

func (invhandler *InvHandler) AddInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inventory Inventory
	d,_:=ioutil.ReadAll(r.Body)
	fmt.Println("AddInventory",string(d))
	// json.NewDecoder(r.Body).Decode(&inventory)
	json.Unmarshal(d,&inventory)
	// fmt.Printf("AddInventory %+v",inventory)
	// print error if any
	// err, ok := invhandler.DB.DB()
	// if ok != nil {
	// 	fmt.Println("Error in DB connection", err)
	// }
	result:=invhandler.DB.Create(&inventory)
	fmt.Println("AddInventory",result)
	fmt.Println("Error",result.Error)
	json.NewEncoder(w).Encode(&inventory)
}

func (invhandler *InvHandler) DelInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inventory Inventory
	params := mux.Vars(r)
	invhandler.DB.Delete(&inventory, params["id"])
	json.NewEncoder(w).Encode(&inventory)
}

func (invhandler *InvHandler) GetIndInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inventory Inventory
	params := mux.Vars(r)
	invhandler.DB.Find(&inventory, params["id"])
	json.NewEncoder(w).Encode(&inventory)
	
}