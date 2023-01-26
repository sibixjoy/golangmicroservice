package main

import (
	"io/ioutil"
	"net/http"

	"example.com/mainmicro/userdb"

	resty "github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
)

type Inventory struct {
	Id int `json:"ID"`
	ProductName string `json:"product_name"`
	ProductQuantity int `json:"product_quantity"`
}

func main() {
	var usrhandlerobj userdb.UsrHandler
	// usrhandlerobj.Connection("host.docker.internal","postgres","1234","orderdb","9000")
	usrhandlerobj.Connection("localhost","postgres","1234","orderdb","9000")

	router:=mux.NewRouter()
	router.HandleFunc("/health", userdb.HealthCheck).Methods("GET")
	router.HandleFunc("/user", usrhandlerobj.GetUser).Methods("GET")
	router.HandleFunc("/adduser", usrhandlerobj.AddUser).Methods("POST")
	router.HandleFunc("/delUser/{id}", usrhandlerobj.DelUser).Methods("DELETE")
	router.HandleFunc("/ordHome", OrdHome).Methods("GET")
	router.HandleFunc("/invHome", InvHome).Methods("GET")
	router.HandleFunc("/proHome", ProHome).Methods("GET")
	router.HandleFunc("/pro/{id}", SinglePro).Methods("GET")
	router.HandleFunc("/delInv/{id}", DelInv).Methods("DELETE")
	router.HandleFunc("/addInv", addInv).Methods("POST")
	router.HandleFunc("/addPro", addPro).Methods("POST")
	router.HandleFunc("/addOrd/{id}", addOrd).Methods("POST")


	http.Handle("/", router)
	http.ListenAndServe(":8400", router)

	dbinstance,_ := usrhandlerobj.DB.DB()
	defer dbinstance.Close()
}

func InvHome(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8200/inventory")
	// print the values in the response
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resp.Body()))
}

func OrdHome(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8300/order")
	// print the values in the response
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resp.Body()))
}

func ProHome(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8100/product")
	// print the values in the response
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resp.Body()))
}

func SinglePro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8100/product/"+vars["id"])
	// print the values in the response
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resp.Body()))
}

func DelInv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	client := resty.New()
	resp, err := client.R().Delete("http://localhost:8200/delinventory/"+vars["id"])
	// print the values in the response
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resp.Body()))
}

func addInv(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	w.Header().Set("Content-Type", "application/json")
	// var inventory Inventory
	d,_:=ioutil.ReadAll(r.Body)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(d)).
		Post("http://localhost:8200/addinventory")
	// print the values in the response
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resp.Body()))


}

func addPro(w http.ResponseWriter, r *http.Request) {
	client := resty.New()
	d,_:=ioutil.ReadAll(r.Body)
	// resp, err := client.R().Post("http://localhost:8100/addproduct")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(d)).
		Post("http://localhost:8100/addproduct")
	// print the values in the response
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resp.Body()))
}

func addOrd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	client := resty.New()
	
	// check the product is availabale in inventory
	resp, _ := client.R().Get("http://localhost:8200/singleinventory/"+vars["id"])

	// check the product is availabale in resp
	if resp.StatusCode() == 200 {
		resp, err := client.R().Post("http://localhost:8300/addorder")
		// print the values in the response
		if err != nil {
			panic(err)
		}
		w.Write([]byte(resp.Body()))
	} else {
		w.Write([]byte("Product is not available in inventory"))
	}

}