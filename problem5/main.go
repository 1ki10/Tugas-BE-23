package main

import "fmt"

// Fungsi menghitung pangkat dari sebuah bilangan.
func pangkat(basis, eksponen int) int {
	hasil := 1
	for i := 0; i < eksponen; i++ {
		hasil *= basis
	}
	return hasil
}

func main() {
	fmt.Println(pangkat(2, 3))  // 8
	fmt.Println(pangkat(5, 3))  // 125
	fmt.Println(pangkat(10, 2)) // 100
	fmt.Println(pangkat(2, 5))  // 32
	fmt.Println(pangkat(7, 3))  // 343
}
