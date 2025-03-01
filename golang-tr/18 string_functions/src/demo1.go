package strings

//alias yani s harfini strings için atayabiliyoruz. Importlarken kullanılabilir kısaltma amaçlı
import (
	f "fmt"
	s "strings"
)

func Demo1() {
	// case sensitive yani büyük küçük harf duyarlıdır.
	isim := "Kerem"
	f.Println(s.Count(isim, "e"))    // sadece küçük "e" harfini sayar.
	f.Println(s.Contains(isim, "e")) // e harfi olup olmadığını kontrol eder
	sonuc := s.Index(isim, "r")      // hangi index değerinde yani baştan kaçıncı sırada olduğunu söyler (0'dan başlayarak)
	// index yazdırması -1 çıkarsa aradığımız harf yoktur
	if sonuc != -1 {
		f.Println("r harfi var")
	} else {
		f.Println("r harfi yoktur")
	}

	f.Println(s.ToLower(isim)) // hepsini küçük harf yapar
	f.Println(s.ToUpper(isim)) // hepsini büyük yazdır
}
