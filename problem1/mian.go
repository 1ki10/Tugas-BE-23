package main

import "fmt"

func main() {
	// Input nilai dan nama mahasiswa
	var studentName string
	var studentScore int

	fmt.Println("Masukkan nama mahasiswa:")
	fmt.Scanln(&studentName)

	fmt.Println("Masukkan nilai mahasiswa:")
	fmt.Scanln(&studentScore)

	// Proses penilaian
	var grade string

	switch {
	case studentScore >= 80 && studentScore <= 100:
		grade = "A"
	case studentScore >= 65 && studentScore <= 79:
		grade = "B+"
	case studentScore >= 50 && studentScore <= 64:
		grade = "B"
	case studentScore >= 35 && studentScore <= 49:
		grade = "C"
	case studentScore >= 0 && studentScore <= 34:
		grade = "D"
	default:
		grade = "Nilai tidak valid"
	}

	// Output
	fmt.Printf("Nama Mahasiswa: %s\n", studentName)
	fmt.Printf("Nilai Huruf: %s\n", grade)
}
