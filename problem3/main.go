package main

import "fmt"

// primeNumber memeriksa apakah suatu bilangan adalah bilangan prima.
func primeNumber(bilangan int) bool {
	// Langsung mengembalikan false jika bilangan kurang dari 2.
	if bilangan < 2 {
		return false
	}

	// Mengecek faktor dari 2 sampai akar kuadrat dari bilangan.
	for i := 2; i*i <= bilangan; i++ {
		if bilangan%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(primeNumber(11)) // true
	fmt.Println(primeNumber(13)) // true
	fmt.Println(primeNumber(17)) // true
	fmt.Println(primeNumber(20)) // false
	fmt.Println(primeNumber(35)) // false
}
