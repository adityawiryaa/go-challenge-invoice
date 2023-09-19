package main

import (
	"fmt"
	"go-invoice/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/invoice", handler.GetInvoice)

	port := 6005
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
