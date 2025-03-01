package strings

import "fmt"

func Slice() {
	sehirler := []string{"İstanbul", "Ankara", "İzmir"}
	//fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	sehirler = append(sehirler, "Adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)
}
