package main

import (
	"RestApiForGo/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main(){
	//Init Router
	route := mux.NewRouter()
	
	//Router Handlers / Endpoints
	route.HandleFunc("/currency", controllers.CreateCurrency).Methods("POST")
	route.HandleFunc("/currencies", controllers.ListCurrency).Methods("GET")
	route.HandleFunc("/currencies/{id}", controllers.GetCurrency).Methods("GET")
	route.HandleFunc("/currencies/{id}", controllers.DeleteCurrency).Methods("DELETE")
	route.HandleFunc("/currencies/{id}", controllers.UpdateCurrency).Methods("PUT")

	
	// Server Init
	port := os.Getenv("PORT") // Environment dosyasından port bilgisi getirilir
	fmt.Println(port)
	if port == "" {
		port = "8080" //localhost:8080
	}
	fmt.Println(port)

	err := http.ListenAndServe(":"+port, route)
	if err != nil {
		fmt.Print(err)
	}
}