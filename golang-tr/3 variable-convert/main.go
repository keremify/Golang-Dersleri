//veri tipi dönüşümü

package main

import "fmt"

func main() {
	x := 10
	y := 10.5

	// float64 olan bir değeri int yaptık

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	fmt.Println(float64(x) + y)
}

//sonuç olarak aynı tiplerde işlem yapmamız gerekiyor. Farklı tipler toplandığında hata alırız. Sonuç da aynı tip çıkar.
