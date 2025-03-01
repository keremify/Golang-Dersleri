package demos

import (
	"fmt"
)

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{tahmin, "Sadece 1-100 arası sayı giriniz."}
	}
	return "Bildiniz", nil
}
