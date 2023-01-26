package orderdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id int `json:"ID"`
	Product string `json:"product"`
	Quantity int `json:"quantity"`
	User_name string `json:"user_name"`
}

type OrdHandler struct {
	DB *gorm.DB
}

func (ordhandler *OrdHandler) Connection(host,user,password,dbname,port string) {
	var err error

	dsn:="host="+host+" user="+user+" password="+password+" dbname="+dbname+" port="+port+" sslmode=disable"
	ordhandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// ordhandler.DB.AutoMigrate(Order{})
	ordhandler.DB.AutoMigrate(&Order{})

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Super Secret Area")
}


func (ordhandler *OrdHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Order []Order
	ordhandler.DB.Find(&Order)
	json.NewEncoder(w).Encode(&Order)
}

func (ordhandler *OrdHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Order Order
	d,_:=ioutil.ReadAll(r.Body)
	// json.NewDecoder(r.Body).Decode(&Order)
	json.Unmarshal(d,&Order)
	ordhandler.DB.Create(&Order)
	json.NewEncoder(w).Encode(&Order)
}
