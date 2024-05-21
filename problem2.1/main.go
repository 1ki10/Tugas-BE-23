package main

import "fmt"

func main() {
	// Input bilangan
	var bilangan int
	fmt.Println("Masukkan bilangan: ")
	fmt.Scanf("%d", &bilangan)

	// Proses untuk mencari faktor-faktor
	fmt.Println("Faktor-faktor dari", bilangan, "adalah:")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}
