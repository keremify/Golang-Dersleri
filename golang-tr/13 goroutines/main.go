// goroutine ve time.Sleep() fonksiyonu kullanımı
// time.Sleep() fonksiyonu ile belirli bir süre bekletme işlemi yapabiliriz.
// goroutine ile aynı anda birden fazla işlemi yapabiliriz.
// başka bir deyişle, goroutine ile bir fonksiyonu çağırıp, aynı anda başka bir fonksiyonu çağırabiliriz.
// normalde bir fonksiyon çağrıldığında, diğer fonksiyonun çağrılmasını bekler. Yani sırayla çalışır.

package main

import (
	"fmt"
	"time"
)

func main() {
	go tekSayilar()             // tekSayilar fonksiyonunu çağır
	go ciftSayilar()            // ciftSayilar fonksiyonunu çağır
	time.Sleep(3 * time.Second) // 1 saniye bekle
	fmt.Println("Main fonksiyon bitti")
}

func ciftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayılar : ", i)
	}
}

func tekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayılar : ", i)
	}
}
