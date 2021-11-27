package models

import (
	u "RestApiForGo/utils"

	"github.com/jinzhu/gorm"
)

// Modeller stract yapısı ile oluşturulur.
type Currency struct {
	gorm.Model  

	Code	string 		`json:"code"`
	Amount	float32 	`json:"amount"`
	Price	float32		`json:"price"`
}


// type CurrencyHistory struct {
// 	gorm.Model  
// 	Code	string 		
// 	History *History	`gorm:"foreignKey:Code;references:Code"`
// }

// type History struct {
// 	gorm.Model  
// 	Code	uint64 		
// 	Amount	float32 	
// 	Price	[]Price 	`gorm:"foreignKey:Code;references:Code"`
// }

// type Price struct {
// 	gorm.Model
// 	Old 	float32		
// 	Current float32		
// }


func (currency *Currency) Contains() (map[string]interface{}, bool) {
	currencyData := &Currency{}
	err := GetDB().Table("currencies").Where("code = ?", currency.Code).First(currencyData).Error
	if err != nil {		
		if err == gorm.ErrRecordNotFound {
			return u.Message(200, ""), true
		}
		return u.Message(500, "Server error"), false			
	} 
	return u.Message(403, "Currency already exists"), false	
}


func (currency *Currency) Validate() (map[string]interface{}, bool) {
	if currency.Code != "" {
		return u.Message(200, ""), true
	}
	return u.Message(400, "Bad request"), false
}

func (currency *Currency) CreateCurrency() map[string]interface{} {
	v_resp, v_status := currency.Validate()
	if !v_status {
		return v_resp
	}
	
	c_resp, c_status := currency.Contains()
	if !c_status {
		return c_resp
	}

	GetDB().Create(currency)

	if currency.ID <= 0 {
		return u.Message(500, "Server error")
	}

	response := u.Message(200, "Success")
	response["currency"] = currency
	return response
}

func ListCurrency() map[string]interface{} {

	var currencies []Currency
	GetDB().Table("currencies").Find(&currencies)

	if len(currencies) == 0{
		return u.Message(404, "Currency with that id does not exist")
	}

	response := u.Message(200, "Success")
	response["currency"] = currencies
	return response
}

func GetCurrency(id string) map[string]interface{} {

	if id  == "" {
		return u.Message(400, "Bad request")
	}

	currency := &Currency{}
	err := GetDB().Table("currencies").Where("id= ?", id).First(currency).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(404, "Currency with that id does not exist")
		}
		return u.Message(500, "Server Error")
	}

	response := u.Message(200, "Success")
	response["currency"] = currency
	return response
}

func (currency *Currency) EditCurrency(id string) map[string]interface{} {
	v_resp, v_status := currency.Validate()
	if !v_status {
		return v_resp
	}
	

	oldCurrency := &Currency{}
	err := GetDB().Table("currencies").Where("id= ?", id).Find(oldCurrency).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(404, "Currency with that id does not exist")
		}
		return u.Message(500, "Server Error")
	}

	if oldCurrency != nil {
		oldCurrency.Amount = currency.Amount
		oldCurrency.Price = currency.Price
		GetDB().Save(oldCurrency)
	} 

	response := u.Message(200, "Success")
	response["currency"] = oldCurrency
	return response
}


func DeleteCurrency(id string) map[string]interface{} {

	if id  == "" {
		return u.Message(400, "Bad request")
	}

	currency := &Currency{}
	err := GetDB().Table("currencies").Where("id= ?", id).First(currency).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(404, "Currency with that id does not exist")
		}
		return u.Message(500, "Server Error")
	}
	
	if currency != nil{
		GetDB().Delete(currency)
	}

	response := u.Message(200, "Success")
	return response
}