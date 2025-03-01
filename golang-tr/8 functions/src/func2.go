package funcs

import "fmt"

func DortIslem() (int, int, int, float32) {
	var sayi1, sayi2 int
	fmt.Print("Lütfen 1. Sayıyı Giriniz = ")
	fmt.Scanln(&sayi1)
	fmt.Print("Lütfen 2. Sayıyı Giriniz = ")
	fmt.Scanln(&sayi2)

	toplam := sayi1 + sayi2
	cikarma := sayi1 - sayi2
	carpma := sayi1 * sayi2
	bolme := float32(sayi1) / float32(sayi2)

	return toplam, cikarma, carpma, bolme
}
