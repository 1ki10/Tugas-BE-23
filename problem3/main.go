package main

import (
	"fmt"
	"math"
)

func main() {

	//input
	var T float64 = 20 // Tinggi tabung
	var r float64 = 4  // Jari-jari tabung

	//kode di simi

	// Rumus luas permukaan tabung: 2 * π * r * (r + T)
	luasPermukaan := 2 * math.Pi * r * (r + T)

	fmt.Printf("Luas permukaan tabung dengan tinggi %.2f cm dan jari-jari %.2f cm adalah %.2f cm²\n", T, r, luasPermukaan)
}
