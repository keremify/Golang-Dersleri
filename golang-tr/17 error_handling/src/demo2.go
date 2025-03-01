package demos

// error kütüphanesi ile yazma
import (
	"errors"
	"fmt"
)

var sayi = 0

func Demo2() {
	result, err := NumCheck(sayi) //err boş kaldığından nil değeri aldı
	if err != nil {
		fmt.Println(err) // error varsa kütüphaneyi kullanarak yazdığımız çıktıyı gösterecek
	} else {
		fmt.Println("Girdiğiniz sayı: ", result)
	}
}

func NumCheck(num int) (int, error) {
	fmt.Println("Bir sayı giriniz :")
	fmt.Scanln(&sayi)
	if sayi%2 != 0 {
		return 0, errors.New("HATA: Numara çift değildir.") //error kütüphanesi ile yeni bir error oluşturduk ve çıktıyı söyledik
	}

	return sayi, nil //başta fonksiyondan int ve error değerleri çıkacağından ikisine de değer atadık. int = sayi ve error = nil atandı
}
