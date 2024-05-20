package main

import "fmt"

func main() {
	//input
	var hargaAwal float64 = 370000
	var diskon float64 = 15

	//code disini

	// Menghitung jumlah diskon
	jumlahDiskon := (diskon / 100) * hargaAwal

	// Menghitung harga akhir setelah diskon
	hargaAkhir := hargaAwal - jumlahDiskon

	fmt.Printf("Harga yang harus dibayar adalah Rp. %.2f\n", hargaAkhir)
}
