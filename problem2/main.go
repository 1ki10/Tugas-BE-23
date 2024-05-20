package main

import "fmt"

func main() {

	//input
	var alas float64 = 20
	var tinggi float64 = 25

	//code disini

	// Rumus luas segitiga: 1/2 * alas * tinggi
	luas := 0.5 * alas * tinggi

	fmt.Printf("Luas segitiga dengan alas %.2f cm dan tinggi %.2f cm adalah %.2f cm²\n", alas, tinggi, luas)
}
