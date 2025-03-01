package defers

import "fmt"

func sayiKontrol(sayi int) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Sayı Çifttir"
	}

	if sayi%2 != 0 {
		return "Sayı Tektir"
	}
	return "Sayı Belirsizdir"
}

func Test() {
	//sonuc := sayiKontrol(9)
	fmt.Println(sayiKontrol(9))
}

func DeferFunc() {
	fmt.Println("Kontrol tamamlandı")
}

// sayiKontrol fonskiyonu çalıştığında returnu görünce direk fonskiyonun en sonuna gider ve fonskiyon biter.
// defer ifadesi ise fonskiyonun tamamlandığını gördüğünde çalışır. Yani defer ifadesi fonskiyonun sonunda çalışır.
// ama yığın mantığıyla ilerlediği için en son yazılan ilk çalışır.
