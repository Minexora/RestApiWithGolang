package controllers

import (
	"RestApiForGo/models"
	u "RestApiForGo/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var CreateCurrency = func(w http.ResponseWriter, r *http.Request) {

	currency := &models.Currency{}
	err := json.NewDecoder(r.Body).Decode(currency) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
	if err != nil {
		u.Response(w, u.Message(400, "Bad request"))
		return
	}

	resp := currency.CreateCurrency()
	u.Response(w, resp)
}

var ListCurrency = func(w http.ResponseWriter, r *http.Request) {
	resp := models.ListCurrency()
	u.Response(w, resp)
}

var GetCurrency = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resp := models.GetCurrency(params["id"])
	u.Response(w, resp)
}

var DeleteCurrency = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Print(params)
	resp := models.DeleteCurrency(params["id"])
	u.Response(w, resp)
}

var UpdateCurrency = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	currency := &models.Currency{}
	err := json.NewDecoder(r.Body).Decode(currency) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
	if err != nil {
		u.Response(w, u.Message(400, "Bad request"))
		return
	}

	resp := currency.EditCurrency(params["id"])
	u.Response(w, resp)		
}