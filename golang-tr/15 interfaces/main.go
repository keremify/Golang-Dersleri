// interface kavramı nedir ve nerede kullanılır
// interface anahtar kelimesi ile bir interface oluşturulur. Interface, bir veya birden fazla fonksiyonun imzasını tanımlar.

package main

import (
	"fmt"
	"math"
)

type sekil interface { // sekil adında bir interface oluştur
	alan() float64  // alan adında bir fonksiyon oluştur
	cevre() float64 // cevre adında bir fonksiyon oluştur
}

type dikdortgen struct { // dikdortgen adında bir struct oluştur
	genislik, yukseklik float64 // genislik ve yukseklik adında iki değişken tanımla
}

type daire struct { // daire adında bir struct oluştur
	yaricap float64 // yaricap adında bir değişken tanımla
}

func (d dikdortgen) alan() float64 { // dikdortgen tipine alan fonksiyonunu tanımla
	return d.genislik * d.yukseklik // dikdortgen tipinin alanını hesapla ve döndür
}
func (c daire) alan() float64 { // daire tipine alan fonksiyonunu tanımla
	return c.yaricap * c.yaricap * math.Pi // daire tipinin alanını hesapla ve döndür
}
func (d dikdortgen) cevre() float64 { // dikdortgen tipine cevre fonksiyonunu tanımla
	return 2 * (d.genislik + d.yukseklik) // dikdortgen tipinin çevresini hesapla ve döndür
}
func (c daire) cevre() float64 { // daire tipine cevre fonksiyonunu tanımla
	return 2 * math.Pi * c.yaricap // daire tipinin çevresini hesapla ve döndür
}

func school(s sekil) { // school adında bir fonksiyon oluştur ve sekil tipinde bir parametre al
	fmt.Println("Şeklin alanı :", s.alan())    // şeklin alanını yazdır
	fmt.Println("Şeklin çevresi :", s.cevre()) // şeklin çevresini yazdır
}

func main() { // main fonksiyonu
	d := dikdortgen{genislik: 10, yukseklik: 20}   // d adında bir dikdortgen oluştur ve genislik ve yukseklik değerlerini ata (structa)
	c := daire{yaricap: 6}                         // c adında bir daire oluştur ve yaricap değerini ata (structa)
	fmt.Println("Dikdörtgenin Alanı ve Çevresi :") // dikdörtgenin çevresi ve alanı
	school(d)                                      // school fonksiyonunu çağır ve d değişkenini parametre olarak gönder
	fmt.Println("Dairenin Alanı ve Çevresi :")     // dairenin çevresi ve alanı
	school(c)                                      // school fonksiyonunu çağır ve c değişkenini parametre olarak gönder
}
