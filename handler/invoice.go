package handler

import (
	"encoding/json"
	"net/http"
)

type RequestData struct {
	Item         string  `json:"item"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	Discount     float64 `json:"discount"`
	TypeDiscount string  `json:"typeDiscount"`
}

type ResponseData struct {
	Item       string  `json:"item"`
	Quantity   int     `json:"quantity"`
	Total      float64 `json:"total"`
	Discount   float64 `json:"discount"`
	GrandTotal float64 `json:"grandTotal"`
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate total
	total := float64(requestData.Quantity) * requestData.Price

	// Calculate discount based on typeDiscount
	var discountAmount float64
	if requestData.TypeDiscount == "%" {
		discountAmount = (requestData.Discount / 100) * total
	} else if requestData.TypeDiscount == "-" {
		discountAmount = requestData.Discount
	} else {
		http.Error(w, "Invalid typeDiscount", http.StatusBadRequest)
		return
	}

	// Calculate grand total
	grandTotal := total - discountAmount

	responseData := ResponseData{
		Item:       requestData.Item,
		Quantity:   requestData.Quantity,
		Total:      total,
		Discount:   discountAmount,
		GrandTotal: grandTotal,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
