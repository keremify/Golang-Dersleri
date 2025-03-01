package main

// json-server --watch db.json
// başlamadan önce yukardaki komutu terminale yazarak localhostta başlat.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// json etiketlerinde boşluk bulunmaz
type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products/")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) //JSON diline benzer bir şekilde yazar. Direk yazdırırsan anlamsız sayı dizileri çıkar.
	//wishlistBeta := string(bodyBytes)             // stringe çevirerek daha okunaklı biçime getiririz. (bu kısım gerekli değil)
	//fmt.Println(wishlistBeta)

	var wishlist []Product
	json.Unmarshal(bodyBytes, &wishlist) // unmarshall ile JSON halini stringe çevirdik []product'a dikkat. Marshal ise JSON'a geri çevirir.

	return wishlist, nil
}

func AddProducts() (Product, error) {
	product := Product{Id: 6, ProductName: "Tablet", CategoryId: 1, UnitPrice: 13000} // bu product bilgilerini sürekli değiştirip yeniden kodu çalıştırarak postlama işlemi gerçekleştirilir.
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err // boş bir product döner çünkü pointer değeri olmayan bir şeyi (err), nil olarak döndüremeyiz
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	fmt.Println("Başarıyla kaydedildi")
	return productResponse, nil

}

func main() {
	AddProducts()
	wishlists, _ := GetAllProducts()
	for i := 0; i < len(wishlists); i++ { // alt alta yazabiliriz böylece
		fmt.Println(wishlists[i].ProductName) // struct olduğu için istediğimiz bilgiyi alt alta yazdırabiliriz.
	}
}

// API'nin çalışma şekline göre post işlemi değişir. Sırayla ID yükselterek postlarken bir numara atlarsanız burda bir sıkıntı olmaz ve
// eklenir. Ama başka API ile çalışırken hata olabilir. Var olan ID üzerine başka birşey eklersek kaydedilmez. ID 5 varken başka bir 5 id
// eklenmemeli.
