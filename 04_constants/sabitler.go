/*Sabitler(Constant)*/

package main

import (
	"fmt"
)

func main() {
	//var i = 5
	//var f = 5.7
	//var c = 5 + 6i
	//fmt.Printf("i'nin türü %T, f'nin türü %T, c'nin türü %T", i,f,c)

	//sabit atama keyword const
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nnfloatVar", float64Var, "\ncomplex64Var", complex64Var)

}
