package demos

// get operasyonuyla çalıştık. sadece go hazır kütüphaneleriyle yapıldı. Workshop dosyasındaki to-do app'te başka kütüphanelere
// bakabilirsiniz.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json: "userId"`
	Id        int    `json: "id"`
	Title     string `json: "title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // veri aldığımız site urlsi
	if err != nil {
		fmt.Println("Bir hata oluştu")
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body) //ioutil.ReadAll alacağımız verinin body kısmını okur. Bunu da bodyBytes'a atarız.

	bodyString := string(bodyBytes) // bodyBytes'ta okunan veriyi stringe çeviririz. Okumamız kolaylaşır. JSON Kodlu verilerdir.
	fmt.Println(bodyString)         // yazdırırız.

	var todo Todo
	json.Unmarshal(bodyBytes, todo) //Unmarshal, JSON kodlu verileri ayrıştırır ve sonucu todo yani virgül sonrası değerde depolar
	fmt.Println(todo)               // yazdırırz.
}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş yap", false}
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Bir hata oluştu")
	}
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	// "gönderilen site, içerik tipi (json), gönderilecek data" sırasıyla doldururuz.
	if err != nil {
		fmt.Println("Bir hata oluştu")
	}
	// Datanın gönderildiğini anlamak için aşağıdaki gibi yeniden okuma isteği göndeririz. Eğer gönderdiğimiz datayı okuyabiliyorsak
	// Post işlemi başarılı olmuştur.

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, todoResponse)
	fmt.Println(todoResponse)
}
