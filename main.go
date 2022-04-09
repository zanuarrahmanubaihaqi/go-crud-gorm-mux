package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mux_crud/handler"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

var db *gorm.DB
var err error

type Product struct {
	ID int `form:"id" json:"id"`
	Code string `form:"code" json:"code"`
	Name string `form:"name" json:"name"`
	Price decimal.Decimal `form:"price" json:"price" sql:"type:decimal(16,2);"`
}

type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

func main() {
    // dsn := "root:@tcp(127.0.0.1:3306)/golang_mux?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err = gorm.Open("mysql", "root:@/golang_mux?charset=utf8&parseTime=True")

    if err != nil{
		log.Println("Connection Failed to Open", err)
    } else { 
		log.Println("Connection Established")
    }

	db.AutoMigrate(&Product{})
	handleRequests()
}

func handleRequests() {
	log.Println("start the development at http://127.0.0.1:9000")

	webRoute := mux.NewRouter().StrictSlash(true)

	webRoute.NotFoundHandler = http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method Not Found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	webRoute.MethodNotAllowedHandler = http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 403, Message: "Method Not Allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	webRoute.HandleFunc("/", homePage)
	webRoute.HandleFunc("/api/product", handler.CreateProduct).Methods("POST")
	webRoute.HandleFunc("/api/products", handler.GetProducts).Methods("GET")
	webRoute.HandleFunc("/api/product/{id}", handler.GetProductById).Methods("GET")
	webRoute.HandleFunc("/api/product/{id}", handler.UpdateProduct).Methods("PUT")
	webRoute.HandleFunc("/api/product/{id}", handler.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", webRoute))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome at home page !")
}

