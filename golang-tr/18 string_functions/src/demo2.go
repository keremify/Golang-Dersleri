package strings

import (
	"fmt"
	"strings"
)

func Demo2() {
	isim := "Kerem"
	fmt.Println(strings.HasPrefix(isim, "Ker")) // Sorguladığımız harflerle başlayıp başlamadığımızı gösterir
	fmt.Println(strings.HasSuffix(isim, "im"))  // Sorguladığımız harflerle bitip bitmediğini gösterir
	fmt.Println(strings.Index(isim, "re"))      // "re" harf grubunun hangi sırada başladığını söyler. Eğer yoksa -1 değeri
	harfler := []string{"k", "e", "r", "e", "m"}
	sonuc := strings.Join(harfler, "*")
	fmt.Println(sonuc)                               // harfler dizisini tek bir string yapıp arasına * yazdırdık. Doldurmazsak da ismi yazar.
	fmt.Println(strings.Replace(sonuc, "e", "i", 2)) //2 tane e harfini i ile değiştirir.
	fmt.Println(strings.Split(sonuc, "*"))           // yıldızları siler ama boşluklar kalır. Eğer olmayan birşeyi splitlersek sonuç normal hali çıkar
	fmt.Println(strings.Repeat(sonuc, 4))            // yazdırırken aynı sonucu 4 kere yan yana yazar
}
