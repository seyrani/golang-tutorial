package main

import (
	"fmt"
	"geometry/rectangle" /*oluşturulan paketimizi import ettik*/
	"log"
)

/*Paket değişkenleri*/
var rectLen, rectWidth float64 //= -1, -7

/*uzunluk ve genişliğin sıfırdan küçü olup olmadığını kontrol eden init fonksiyonu*/
func init() {
	println("Ana paket başlatıldı")

	if rectLen < 0 {
		log.Fatal("Girilen uzunluk değeri sıfırdan küçük")
	}
	if rectWidth < 0 {
		log.Fatal("Girilen genişlik değeri sıfırdan küçük")
	}
}

func main() {

	/*Kullanıcıdan veri alır paket değişkenlerine atama yapar*/
	fmt.Print("Uzunluk değeri giriniz:")
	fmt.Scanf("%v", &rectLen)
	fmt.Print("Genişlik değeri giriniz:")
	fmt.Scanf("%v", &rectWidth)

	fmt.Println("Geometrik şeklin özellikleri")
	/*Dikdörtgen paketin alan hesabı*/
	fmt.Printf("Dikdörtgenin alanı %.2f\n", rectangle.Alan(rectLen, rectWidth))
	/*Köşegenin hesaplama işlemi*/
	fmt.Printf("Dikdörtgenin köşegenleri %.2f ", rectangle.Kosegen(rectLen, rectWidth))

}
