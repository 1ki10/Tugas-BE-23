package main

import "fmt"

func main() {
	// Input nilai mahasiswa (misalnya 80)
	var studentScore int = 80

	// Konversi nilai numerik menjadi nilai huruf
	var grade string
	switch {
	case studentScore >= 80:
		grade = "A"
	case studentScore >= 65:
		grade = "B+"
	case studentScore >= 50:
		grade = "B"
	case studentScore >= 35:
		grade = "C"
	default:
		grade = "D"
	}

	// Output
	fmt.Printf("Nilai huruf: %s\n", grade)
}
