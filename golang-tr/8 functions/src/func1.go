package funcs

import "fmt"

func Topla(sayi1 int, sayi2 int) int {
	var toplam = sayi1 + sayi2
	return toplam
}

func SelamVer(kullaniciAdi string) {
	fmt.Println("Merhaba, " + kullaniciAdi)
}
