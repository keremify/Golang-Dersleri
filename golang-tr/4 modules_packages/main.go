package main

import variables "golesson/src"

func main() {
	variables.Demo1()
}

// go dilinde her bir dosyayı paket olarak varsayarız
// modüller de paketleri bir arada tutan bir alan gibi düşünebiliriz. Yani modüllerin içinde paketler bulunur
// modül oluşturmak için terminalde cmd kısmına "go mod init" yazıp çalıştırmalıyız.
// Modül oluşturma amacımız tüm paketleri tek bir fonksiyon ile çalıştırabilmemiz.
// Mesela yukarıda variables paketinden demo1 fonksiyonunu başka bir dosyadan çektik. Ve onlarca satır kodu tek bir yerde çalıştırdık
// importlarken de (import "modüladı/fonskiyonadı") formatıyla içeri aktarırız
