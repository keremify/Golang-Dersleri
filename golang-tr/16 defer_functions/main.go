package main

import (
	defers "deferler/src"
	"fmt"
)

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}

func B() {
	fmt.Println("B fonskiyonu çalıştı")
}

func C() {
	fmt.Println("C fonksiyonu çalıştı")
}

func main() {
	//defer A()
	//defer C()
	//fmt.Println("Bitti")
	//defers.Test()
	defers.Demo3()
}

// defer ifadesi yığın olarak ilerler. Hata olup olmadığını kontrol etmek için kullanılır. Yani defersiz olarak A() ve
// C() yazdığımızda ve main çalışırken önceki satırlarda hata olduğunda iki fonksiyon da çalışmazdı. Yani o fonksiyonların
// her şartta çalışacağını sağlarız. Fonskiyon çalışırken önce fonksiyon tamamlanır sonra defer ifadelerine geçilir. defer ifadeleri
// çalıştırılırken en son yazılan ilk çalışır. Üst üste atmak gibi düşün.
