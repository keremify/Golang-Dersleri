package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya"

	for i := 1; i <= 5; i++ {
		fmt.Println(i, metin)
	}
	fmt.Println("Bu kadardi.")
}
