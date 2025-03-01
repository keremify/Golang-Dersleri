package main

import (
	"fmt"
	"math"
)

type sekil interface {
	alan() float64
	cevre() float64
}

type kare struct {
	kenar float64
}

type daire struct {
	yariCap float64
}

func (k kare) alan() float64 {
	kareAlan := k.kenar * k.kenar
	return kareAlan
}

func (k kare) cevre() float64 {
	kareCevre := k.kenar * 4
	return kareCevre
}

func (d daire) alan() float64 {
	daireAlan := d.yariCap * d.yariCap * math.Pi
	return daireAlan
}

func (d daire) cevre() float64 {
	daireCevre := 2 * d.yariCap * math.Pi
	return daireCevre
}

func hesapla(s sekil) {
	fmt.Println("Şeklin alanı :", s.alan())
	fmt.Println("Şeklin çevresi :", s.cevre())
}

func main() {
	var girilenKenar, girilenYariCap float64

	fmt.Println("Karenin kenar uzunluğunu giriniz:")
	fmt.Scanln(&girilenKenar)
	k := kare{kenar: girilenKenar}

	fmt.Println("Dairenin yarıçapını giriniz:")
	fmt.Scanln(&girilenYariCap)
	d := daire{yariCap: girilenYariCap}

	fmt.Println("Karenin Özellikleri :")
	hesapla(k)
	fmt.Println("Dairenin Özellikleri :")
	hesapla(d)
}
