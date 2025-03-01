/* Structs Konusu
- Yani aslında farklı özellikleri olan bir veri tipi oluşturmak için kullanılır.
- Struct'lar, birbirine bağlı farklı veri tiplerini bir arada tutmak için kullanılır.
*/

package main

import "fmt"

type product struct { // product struct'ı oluştur
	name      string // name, unitPrice ve brand alanlarına sahip
	unitPrice float64
	brand     string
}

type customer struct { // customer struct'ı oluştur
	firstName string // firstName, lastName, age ve credit alanlarına sahip
	lastName  string
	age       int
	credit    float64
}

func main() {

	fmt.Println(product{"Bilgisayar", 17999, "Lenovo"}) // {Bilgisayar 17999 Lenovo}
	demo2()                                             //demo2 fonskiyonunu çağır. Bu fonksiyon customer struct'ı için save fonksiyonunu çağırır.
}

func (c customer) save() { // customer struct'ı için save fonksiyonu oluştur
	fmt.Println("Eklendi : ", c.firstName) // Eklendi :  Kerem
	// Burada demo2 fonksiyonunda atadığımız c.firstName'i yazdırır
}

func demo2() { // customer struct'ı için save fonksiyonunu çağır
	c := customer{firstName: "Kerem", lastName: "Özdemir", age: 18} // customer struct'ının içini doldur
	c.save()                                                        // save fonksiyonunu çağır (yukarıda tanımlanan fonksiyon)
}
