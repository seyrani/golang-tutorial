package main

import (
	"fmt"
)

func main() {

	/*
		1. kullanım yöntemi
	*/
	sayi := 15

	if sayi%4 == 0 {
		fmt.Println("Sayı tek sayıdır.")
	} else {
		fmt.Println("Sayı çift sayıdır")
	}

	/*
		2. kullanım yöntemi
	*/

	//sayi2 içerden erişimde kullanılır
	if sayi2 := 8; sayi2%2 == 0 {
		fmt.Println(sayi2, "sayı çift")
	} else {
		fmt.Println(sayi2, "sayı tek")
	}

	//ornek 2

	sayi3 := 101

	if sayi3 <= 50 {
		fmt.Println("Sayı 50'den küçük veya eşit")
	} else if sayi3 >= 51 && sayi3 <= 100 {
		fmt.Println("Sayı 51 ve 100 arasında")
	} else {
		fmt.Println("Sayı 100'den büyüktür.")
	}

}
