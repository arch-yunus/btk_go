package main

import "fmt"

// Add, iki tamsayıyı toplar. (TDD örneği)
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Testing Fundamentals Modülü")
	fmt.Printf("2 + 3 = %d\n", Add(2, 3))
}
