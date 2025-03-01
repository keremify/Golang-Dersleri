package main

import "fmt"

func main() {
	sozluk := make(map[string]string)
	sozluk["Book"] = "Kitap"
	sozluk["Masa"] = "Table"

	fmt.Println(sozluk["Book"])
	delete(sozluk, "Book")
}
