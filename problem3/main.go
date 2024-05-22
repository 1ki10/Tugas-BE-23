package main

import "fmt"

func main() {
	//input
	var T float64 = 20 // Tinggi tabung
	var r float64 = 4  // Jari-jari tabung

	//kode di sini
	// Rumus luas permukaan tabung: 2 * π * r * (r + T)
	luasPermukaan := 2 * 3.14 * r * (r + T)

	// output
	fmt.Println(luasPermukaan)
}
