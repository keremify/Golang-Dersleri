package demos

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("control.txt")
	//nil değersizlik demek, err tarafına birşey atamadığımız için otomatik olarak nil oldu
	if err != nil { //err hata olarak düşün. eğer hata değeri nilden farklıysa (atanmışsa) aşağıdaki kodlar çalışır.
		if pErr, ok := err.(*os.PathError); ok { // hataları çeşitlendirmek için birdaha if ile yazarız. pErr burda patherror değerini aldı. yani eğer patherror değeri varsa (nil değilse) aşağısındaki dediği gibi hatanın türünü yazdıracak.
			fmt.Println("Dosya bulunamadı", pErr.Path)
			return
		} else { //path error harici bir hata ise bu çalışacak. Bunu da atayabilirdik ama vakit kaybetmeyelim
			fmt.Println("Bilinmeyen bir hata oluştu.")
		}
		fmt.Println("Dosya bulunamadı.")
	} else {
		fmt.Println(f.Name())
	}
}
