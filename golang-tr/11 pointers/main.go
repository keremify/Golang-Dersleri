/*
* ve & operatörleri :
* operatörü bir değişkenin değerini döndürürken, & operatörü bir değişkenin adresini döndürür.
Bir değişkenin adresini almak için & operatörü kullanılır.
Bir değişkenin değerini almak için * operatörü kullanılır.
*/

package main

import "fmt"

func main() {
	var x int = 42  // x değişkeni oluştur ve değerini ata
	var p *int = &x // x değişkeninin bellek adresini p'ye ata (p, x'in adresini tutar)

	if p != nil { // p'nin değeri nil değilse
		fmt.Println(*p) // p'nin işaret ettiği değeri yazdır
	} // 42

	fmt.Println("x'in değeri:", x)                // 42
	fmt.Println("x'in adresi:", &x)               // Bellek adresi
	fmt.Println("p'nin tuttuğu adres:", p)        // Bellek adresi
	fmt.Println("p'nin işaret ettiği değer:", *p) // 42

	// p ile x'in değerini değiştirme
	*p = 100                            // p'nin işaret ettiği değeri değiştir
	fmt.Println("x'in yeni değeri:", x) // 100 olacak, çünkü p'nin işaret ettiği değeri değiştirdik ve ikisi aynı bellek adresini tutuyor
}

// dizilerde pointer kullanımı farklıdır. Diziler zaten pointer varmış gibi davranır ve başka bir diziye atandığında,
// aslında aynı bellek adresini paylaşırlar. Yani bir dizinin bir kopyasını oluşturmak için yeni bir bellek alanı ayrılmaz.
