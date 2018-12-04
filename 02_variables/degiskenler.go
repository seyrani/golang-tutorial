package main

import(
	"fmt"
	"math"
)


func main(){

	var age int
	fmt.Println("Yaşınız :", age)
	age = 32
	fmt.Println("Yaşınız :", age)
	age = 54
	fmt.Println("Yeni Yaşınız: ", age)

	//degişkene değer atanabilir
	var age2 int = 42
	fmt.Println("Başlangıç yaş: ", age2)

	var age3 = 15
	fmt.Println("değişken tür girilmezse go otomatik türü bilir ", age3)

	//çoklu değişken ıtanımlama
	var genislik, yukseklik int =100,50
	fmt.Println("Genişlik ", genislik, "Yükseklik ", yukseklik)

	//farklı türlere ait değiken atama
	var(
		ad = "Ali"
		yas = 46
		boy = 1.75		
	)

	fmt.Println("Adı :", ad, "Yaşı: ", yas, "Boy: ", boy )

	/*kısa syntax kullanımı := sol taraftaki 
	değeri sağdaki değikenlere aktarır*/

	a,b := 35,45
	fmt.Println("a'nın değeri:", a, "b'nin degeri ", b)

	/*Değişkenlere, çalışma süresi boyunca hesaplanan 
	değerler de atanabilir.*/ 

	a1, b1 := 145.8, 543.8
	c1 := math.Min(a1,b1)
	fmt.Println("minimum değer ", c1)



	

}
