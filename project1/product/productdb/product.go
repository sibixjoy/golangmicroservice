package productdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductDB struct {
	gorm.Model
	Id int `json:"ID"`
	Name string `json:"name"`
	Price int `json:"price"`
}

type ProHandler struct {
	DB *gorm.DB
}

func (proHandler *ProHandler) Connection(host,user,password,dbname,port string) {
	var err error

	dsn:="host="+host+" user="+user+" password="+password+" dbname="+dbname+" port="+port+" sslmode=disable"
	proHandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// proHandler.DB.AutoMigrate(ProductDB{})
	proHandler.DB.AutoMigrate(&ProductDB{})

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Super Secret Area")
}


func (proHandler *ProHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product []ProductDB
	proHandler.DB.Find(&product)
	json.NewEncoder(w).Encode(&product)
}

func (proHandler *ProHandler) GetIndProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product ProductDB
	proHandler.DB.Find(&product)
	json.NewEncoder(w).Encode(&product)
}

func (proHandler *ProHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product ProductDB
	d,_:=ioutil.ReadAll(r.Body)
	// json.NewDecoder(r.Body).Decode(&product)
	json.Unmarshal(d,&product)
	proHandler.DB.Create(&product)
	json.NewEncoder(w).Encode(&product)
}