package asalsay

import "fmt"

func AsalSayiBul() {
	sayi := 0
	fmt.Println("Bir sayı giriniz: ")
	fmt.Scan(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
			break
		}
	}
	if asalMi {
		fmt.Println(sayi, "sayısı asal bir sayıdır.")
	} else {
		fmt.Println(sayi, "sayısı asal bir sayı değildir.")
	}
}
