package main

import "fmt"

func main() {

	a := true
	b := false
	fmt.Println("a :", a, "b :", b)

	c := a && b
	fmt.Println("c: ", c)

	d := a || b
	fmt.Println("d: ", d)

	//tür dönüşüm

	i := 55
	j := 65.5
	sum := i + int(j) //j integer'a dönüştü
	fmt.Println("Toplam : ", sum)

	i2:=10
	var j1 float64 = float64(i2) //float64'e dönüştürülür ve sonra j'ye atanır
	fmt.Println("j :", j1)

}
