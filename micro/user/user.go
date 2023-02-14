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
type User struct {

	Id int `json:"ID"`
	User_name string `json:"user_name"`
	Email string `json:"email"`
}
var db *gorm.DB 

func main() {
	var err error
   
	dsn := "host=host.docker.internal user=postgres password=1234 dbname=micro port=9000 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	 panic("failed to connect to database")
	}
   userdb,_:=db.DB()
	   fmt.Println("Connection opened to database")
   
   db.AutoMigrate(&User{})
   
	router := mux.NewRouter()
   
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/user", user).Methods("GET")
	router.HandleFunc("/adduser", Adduser).Methods("POST")
	router.HandleFunc("/getuser/{id}", Getuser).Methods("GET")
    router.HandleFunc("/deleteuser/{id}",Deluser) .Methods("DELETE")


	http.ListenAndServe(":8004", router)
   
	defer userdb.Close() 
}	

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Area")
}

func user(w http.ResponseWriter, r *http.Request) {

    var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)

}

func Adduser(w http.ResponseWriter, r *http.Request) {

	var newuser User

	reqBody, _:= ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &newuser)
	db.Create(&newuser)

	json.NewEncoder(w).Encode(newuser)

}
func Getuser(w http.ResponseWriter, r *http.Request) {
	var user User
	userID := mux.Vars(r)["id"]

		 db.First(&user, userID) 
			json.NewEncoder(w).Encode(user)
		
}
func Deluser(w http.ResponseWriter, r *http.Request) {
	var deluser User
	userID := mux.Vars(r)["id"]

		 db.Delete(&deluser, userID) 
			json.NewEncoder(w).Encode(deluser)
		
}