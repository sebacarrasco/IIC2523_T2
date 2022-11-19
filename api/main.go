package main

import (
    "fmt"
    "log"
    "net/http"
		"io/ioutil"
		"encoding/json"
		"strconv"
		"github.com/gorilla/mux"
)

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Value int `json:"value"`
	Date string `json:"date"`
}

var Products []Product

// Endpoints
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Homepage of the Products API!")
    fmt.Println("Endpoint Hit: homePage")
}

func getProducts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: get Products")
	json.NewEncoder(w).Encode(Products)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: get Product by Id")
	vars := mux.Vars(r)
	pid, _ := strconv.Atoi(vars["id"])
	for _, product := range Products {
			if product.Id == pid {
					json.NewEncoder(w).Encode(product)
			}
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: post Products")
	reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product
    json.Unmarshal(reqBody, &product)
    Products = append(Products, product)
    json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: delete Product")
	vars := mux.Vars(r)
	pid, _ := strconv.Atoi(vars["id"])
	for index, product := range Products {
			if product.Id == pid {
					Products = append(Products[:index], Products[index+1:]...)
			}
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update Product")
	vars := mux.Vars(r)
	pid, _ := strconv.Atoi(vars["id"])
	for index, product := range Products {
			if product.Id == pid {
					reqBody, _ := ioutil.ReadAll(r.Body)
					var product Product
					json.Unmarshal(reqBody, &product)
					// change only the attributes that are not empty
					if product.Id != 0 {
							Products[index].Id = product.Id
					}
					if product.Name != "" {
							Products[index].Name = product.Name
					}
					if product.Description != "" {
							Products[index].Description = product.Description
					}
					if product.Value != 0 {
							Products[index].Value = product.Value
					}
					if product.Date != "" {
							Products[index].Date = product.Date
					}
					json.NewEncoder(w).Encode(Products[index])
			}
	}
}


// Handle all requests
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/products", getProducts)
	router.HandleFunc("/products/{id}", getProductById)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	fmt.Println("Corriendo en el puerto 3000")
	// Products = []Product{
  //       Product{Id: 0, Name: "Hello1", Description: "Article Description 1", Value: 10, Date: "hoy"},
  //       Product{Id: 1, Name: "Hello 2", Description: "Article Description 2", Value: 11, Date: "mañana"},
  //   }
  handleRequests()
}