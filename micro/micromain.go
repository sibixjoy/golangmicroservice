package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

    "github.com/go-resty/resty/v2"
	"io/ioutil"
	"encoding/json"

)

var db *gorm.DB 

func main() {
	var err error
   
	dsn := "host=localhost user=postgres password=1234 dbname=micro port=9000 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	 panic("failed to connect to database")
	}
    maindb,_:=db.DB()
	   fmt.Println("Connection opened to database")
   
   
	router := mux.NewRouter()
   
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/userlist",userlist).Methods("GET")
	router.HandleFunc("/productlist", productlist).Methods("GET")
	router.HandleFunc("/orderlist", Orderlist).Methods("GET")
	router.HandleFunc("/GETuser/{id}", Getuserbyid).Methods("GET")
	router.HandleFunc("/GETproduct/{id}", Getproductbyid).Methods("GET")
	router.HandleFunc("/ADDuser", AddUser).Methods("POST")
	router.HandleFunc("/ADDorder/{id}", AddOrder).Methods("POST")

	
	http.ListenAndServe(":8000", router)
   
	defer maindb.Close() 
}	
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Area")
}

func userlist(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	resp, _ := client.R().
      Get("http://localhost:8004/user")
	
	w.Write([]byte(resp.Body()))


}

func productlist(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	resp, _ := client.R().
	
	Get("http://localhost:8002/product")
	

	w.Write([]byte(resp.Body()))

}

func Orderlist(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	resp, _ := client.R().
	Get("http://localhost:8001/order")
	

	w.Write([]byte(resp.Body()))

}

func Getuserbyid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	client := resty.New()
	resp, _ := client.R().
	Get("http://localhost:8004/getuser/"+id)
	

	w.Write([]byte(resp.Body()))

}

func Getproductbyid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	client := resty.New()
	resp, _ := client.R().
	Get("http://localhost:8002/getproduct/"+id)
	

	w.Write([]byte(resp.Body()))

}



func AddUser(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	reqBody, _ := ioutil.ReadAll(r.Body)
	

	resp, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post("http://localhost:8004/adduser")
		
	w.Write([]byte(resp.Body()))
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	var jsonString string
	id := mux.Vars(r)["id"]

	resp, _ := client.R().
	 Get("http://localhost:8003/getinventory/"+id)
		
    if resp.StatusCode()== 200 {

		reqBody, _ := ioutil.ReadAll(r.Body)
	    resporeder, _ := client.R().
		    SetHeader("Content-Type", "application/json").
		    SetBody(reqBody).
		    Post("http://localhost:8001/addorder")
		
	w.Write([]byte(resporeder.Body()))


    }else{
	    jsonString = `"message" : "no stockavilable"`
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonString)
    }
}