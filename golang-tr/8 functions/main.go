package main

import (
	"fmt"
	funcs "func/src"
)

func main() {
	//var sonuc = funcs.Topla(2, 7)
	//fmt.Println(sonuc)
	//funcs.SelamVer("Kerem")
	sonuc1, sonuc2, sonuc3, sonuc4 := funcs.DortIslem()
	fmt.Println("Toplama sonucu = ", sonuc1)
	fmt.Println("Çıkarma Sonucu = ", sonuc2)
	fmt.Println("Çarpma Sonucu = ", sonuc3)
	fmt.Println("Bölme Sonucu = ", sonuc4)
}
