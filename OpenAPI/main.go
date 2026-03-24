package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Petstore client başlıyor...")
	resp, err := http.Get("https://petstore.swagger.io/v2/pet/findByStatus?status=available")
	if err != nil {
		fmt.Println("Hata: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code: ", resp.StatusCode)
}
