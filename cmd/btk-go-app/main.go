package main

import (
	"fmt"
	"golesson/refactoring"
)

func main() {
	fmt.Println("BTK Go Course Gateway Başlatıldı...")
	
	// Refactored project logic call
	products, _ := refactoring.GetAllProducts()
	fmt.Printf("Sistemde %d ürün yüklü.\n", len(products))
}
