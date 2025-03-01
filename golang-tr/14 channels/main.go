// channel aynı zamanda "kanal" anlamına gelir.
// channel ile bir fonksiyonun işlemi bittiğinde diğer fonksiyonun çalışmasını sağlayabiliriz.
// yani işlem sırasını belirleyebiliriz.

package main

import (
	"fmt"
	"time"
)

func main() {
	ciftSayiToplamCn := make(chan int) // ciftSayiToplamCn adında bir kanal oluştur
	tekSayiToplamCn := make(chan int)  // tekSayiToplamCn adında bir kanal oluştur
	go tekSayilar(tekSayiToplamCn)     // tekSayilar fonksiyonunu çağır ve tekSayiToplamCn kanalını parametre olarak gönder
	go ciftSayilar(ciftSayiToplamCn)   // ciftSayilar fonksiyonunu çağır ve ciftSayiToplamCn kanalını parametre olarak gönder
	time.Sleep(1 * time.Second)        // 3 saniye bekle

	ciftSayiToplam, tekSayiToplam := <-ciftSayiToplamCn, <-tekSayiToplamCn // ciftSayiToplamCn ve tekSayiToplamCn kanallarından veri al
	carpim := ciftSayiToplam * tekSayiToplam                               // ciftSayiToplam ve tekSayiToplam değişkenlerini çarp
	fmt.Println("Çarpım : ", carpim)                                       // carpim değişkenini yazdır
}

func ciftSayilar(ciftSayiCn chan int) { // ciftSayiCn adında bir kanal oluştur
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
	}
	ciftSayiCn <- toplam // toplam değişkenini ciftSayiCn kanalına gönder
}

func tekSayilar(tekSayiCn chan int) { // tekSayiCn adında bir kanal oluştur
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
	}
	tekSayiCn <- toplam // toplam değişkenini tekSayiCn kanalına gönder
}

// return ifadesiyle çalışsaydık, fonksiyonlar aynı anda çalışmayacaktı.
// channel ile hem asenkron çalışmasını sağladık, hem de işlem sırasını belirledik.
