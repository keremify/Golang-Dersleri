package main

import "fmt"

func main() {
	// üç adet değer tanımla, ekrana en büyüğünü yazdır
	var sayi1, sayi2, sayi3 int = 10, 71, 91
	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}
	fmt.Printf("En büyük sayı = %v", enBuyuk)
}

//ingilizce konuşmak gibi. Eğer komutuyla alternatif senaryolarda ne yapılacağını gösteririz.
