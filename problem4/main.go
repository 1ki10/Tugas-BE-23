package main

import (
	"fmt"
	"strings"
)

func palindrome(input string) bool {
	// Menghapus spasi dan mengubah semua huruf menjadi lowercase
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)

	// Membalik string
	reversed := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}

	// Membandingkan string asli dengan yang sudah dibalik
	return input == reversed
}

func main() {
	fmt.Println(palindrome("civic"))       // true
	fmt.Println(palindrome("katak"))       // true
	fmt.Println(palindrome("kasur rusak")) // true
	fmt.Println(palindrome("kupu-kupu"))   // false
	fmt.Println(palindrome("lion"))        // false
}
