package main

import (
	"fmt"
	"math"
)

// Fungsi untuk memeriksa apakah suatu bilangan adalah bilangan prima
func isPrima(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk memeriksa apakah suatu bilangan adalah Full Prima
func FullPrima(N int) bool {
	if !isPrima(N) {
		return false
	}

	// Memeriksa setiap digit bilangan
	for N > 0 {
		digit := N % 10
		if !isPrima(digit) {
			return false
		}
		N /= 10
	}
	return true
}

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
