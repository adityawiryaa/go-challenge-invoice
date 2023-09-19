package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func GetInvoice(c *gin.Context) {

	var requestData RequestData
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Log the request body
	fmt.Printf("Request Body: %s\n", string(body))

	// Unmarshal the JSON data into the requestData struct
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
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
		log.Println("ERROR Type Discount", requestData)
		// http.Error(c.Writer, "Invalid typeDiscount", http.StatusBadRequest)
		// return
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

	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(responseData)
}
