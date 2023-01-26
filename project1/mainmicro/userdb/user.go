package userdb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id int `json:"ID"`
	User_name string `json:"user_name"`
	Email string `json:"email"`
}

type UsrHandler struct {
	DB *gorm.DB
}

func (ordhandler *UsrHandler) Connection(host,user,password,dbname,port string) {
	var err error

	dsn:="host="+host+" user="+user+" password="+password+" dbname="+dbname+" port="+port+" sslmode=disable"
	ordhandler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// ordhandler.DB.AutoMigrate(User{})
	ordhandler.DB.AutoMigrate(&User{})

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Super Secret Area")
}


func (usrhandler *UsrHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User []User
	usrhandler.DB.Find(&User)
	json.NewEncoder(w).Encode(&User)
}

func (usrhandler *UsrHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User User
	json.NewDecoder(r.Body).Decode(&User)
	usrhandler.DB.Create(&User)
	json.NewEncoder(w).Encode(&User)
}

func (usrhandler *UsrHandler) DelUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User User
	json.NewDecoder(r.Body).Decode(&User)
	usrhandler.DB.Delete(&User)
	json.NewEncoder(w).Encode(&User)
	
}