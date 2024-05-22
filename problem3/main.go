package main

import "fmt"

// cetakTabelPerkalian
func cetakTabelPerkalian(angka int) {
	for i := 1; i <= angka; i++ {
		for j := 1; j <= angka; j++ {
			// Cetak hasil perkalian dan atur spasi
			fmt.Print(i*j, " ")
		}
		// Ganti baris setelah setiap baris
		fmt.Println()
	}
}

func main() {
	cetakTabelPerkalian(4)
	// Output :
	// 1 2 3 4
	// 2 4 6 8
	// 3 6 9 12
	// 4 8 12 16
}
