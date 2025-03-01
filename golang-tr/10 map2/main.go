package main

import "fmt"

func main() {
	sozluk := map[string]string{"Kitap": "Book", "Masa": "Table"}

	//k ilk değer v diğer değer oluyor

	for k, v := range sozluk {
		fmt.Println(k)
		fmt.Println(v)
	}

}

// sözlük yapmada key ve value değerlerini belirleyip for döngüsü ile yazdırabiliriz.
