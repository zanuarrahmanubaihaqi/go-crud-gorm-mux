package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:@/golang_mux?charset=utf8&parseTime=True")
	
	playloads, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(playloads, &product)
	
	db.Create(&product)

	res := Result{Code: 200, Data: product, Message: "product data created"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:@/golang_mux?charset=utf8&parseTime=True")
	fmt.Println("get all products")

	products := []Product{}
	db.Find(&products)

	res := Result{Code: 200, Data: products, Message: "success get product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:@/golang_mux?charset=utf8&parseTime=True")
	
	vars := mux.Vars(r)
	productId := vars["id"]
	fmt.Println("get product by id : ", productId)
	
	var product Product
	db.First(&product, productId)

	res := Result{Code: 200, Data: product, Message: "succes get product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:@/golang_mux?charset=utf8&parseTime=True")
	
	vars := mux.Vars(r)
	productId := vars["id"]
	fmt.Println("update product with id : ", productId)

	playloads, _ := ioutil.ReadAll(r.Body)

	var dataUpdate Product
	json.Unmarshal(playloads, &dataUpdate)

	var product Product
	db.First(&product, productId)
	db.Model(&product).Updates(dataUpdate)

	res := Result{Code: 200, Data: product, Message: "updated success"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:@/golang_mux?charset=utf8&parseTime=True")

	vars := mux.Vars(r)
	productId := vars["id"]
	fmt.Println("delete product with id : ", productId)

	var product Product
	db.First(&product, productId)
	db.Delete(&product)

	res := Result{Code: 200, Message: "deleted succes"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http. StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}