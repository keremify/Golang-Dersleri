package defers

import "fmt"

type product struct {
	productName string
	unitPrice   string
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi:", p.productName)
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: "5000$"}
	defer p.save()
	fmt.Println("Ürün eklendi")
}
