package main

import (
	"fmt"
)

func alanHesap(uzunKenar, kisaKenar float64) (float64, float64) {

	var alan = uzunKenar * kisaKenar
	var cevre = (uzunKenar + kisaKenar) * 2
	return alan, cevre
}

func main() {
	//alan, cevre := alanHesap(10.8, 5.6)
	//fmt.Printf("Alan %f çevre %f", alan, cevre)

	//parametreyi atmak için "_" tanımlayıcı kullanır
	alan, _ := alanHesap(10.5, 5.8)
	fmt.Printf("Alan %f ", alan)
}
