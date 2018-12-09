/*
* For Döngüsü
* Sözdizimi :
* for initialisation; condition; post {
* }
 */

package main

import "fmt"

func main() {

	//1-10 arası sayıları sıralar
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	//0-10 arası çift sayıları yazdır

	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}

	//
	for s1, i := 10, 1; i <= 10 && s1 <= 19; i, s1 = i+1, s1+1 {
		fmt.Printf("%d * %d = %d\n", s1, i, s1*i)
	}
}
