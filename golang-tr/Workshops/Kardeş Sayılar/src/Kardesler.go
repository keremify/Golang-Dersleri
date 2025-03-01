package kardes

import "fmt"

func KardesSayiBul() {
	sayi1 := 0
	sayi2 := 0
	toplam1 := 0
	toplam2 := 0
	fmt.Println("Birinci sayıyı giriniz: ")
	fmt.Scanln(&sayi1)
	fmt.Println("İkinci sayıyı giriniz: ")
	fmt.Scanln(&sayi2)

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 = toplam1 + i
		}
	}
	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 = toplam2 + i
		}
	}
	if sayi1 == toplam2 && sayi2 == toplam1 {
		fmt.Printf("%d ve %d sayıları birbirinin kardeş sayılarıdır.", sayi1, sayi2)
	} else {
		fmt.Printf("%d ve %d sayıları birbirinin kardeş sayıları değildir.", sayi1, sayi2)
	}
}
