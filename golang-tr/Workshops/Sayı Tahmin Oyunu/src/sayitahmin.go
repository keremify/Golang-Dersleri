package bulucu

import "fmt"

func SayiBul() {
	aklimdakiSayi := 80
	tahminEdilen := 0
	tahminHakki := 4

	fmt.Println("Bir sayı tahmin ediniz:")
	fmt.Println("Tahmin hakkınız: " + fmt.Sprint(tahminHakki+1))
	fmt.Scanln(&tahminEdilen)

	for tahminHakki >= 0 && aklimdakiSayi != tahminEdilen {
		if tahminHakki == 0 {
			fmt.Println("Tahmin hakkınız bitti. Bulmanz gereken sayı: " + fmt.Sprint(aklimdakiSayi))
			break
		}
		if tahminEdilen < 80 {
			fmt.Println("Daha büyük bir sayı tahmin ediniz.")
			fmt.Scanln(&tahminEdilen)
			tahminHakki--
		} else {
			fmt.Println("Daha küçük bir sayı tahmin ediniz.")
			fmt.Scanln(&tahminEdilen)
			tahminHakki--
		}

		if tahminEdilen == 80 {
			fmt.Println("Tebrikler doğru tahmin!")
			fmt.Printf("Kalan tahmin hakkınız: %d\n", tahminHakki)
		}
	}
}
