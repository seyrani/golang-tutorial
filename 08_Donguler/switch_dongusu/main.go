package main

import "fmt"

func sayi() int {
	s1 := 15 * 5
	return s1
}

func main() {

	ay := 3

	switch ay {
	case 1:
		fmt.Println("Ocak")
	case 2:
		fmt.Println("Şubat")
	case 3:
		fmt.Println("Nisan")
	case 4:
		fmt.Println("Mayıs")
	case 5:
		fmt.Println("Haziran")
	case 6:
		fmt.Println("Temmuz")
	case 7:
		fmt.Println("Ağustos")
	case 8:
		fmt.Println("Eylül")
	case 9:
		fmt.Println("Ekim")
	case 10:
		fmt.Println("Kasım")
	case 11:
		fmt.Println("Aralık")
	}

	//sesli harfleri kontrol eder
	harf := "i"

	switch harf {
	case "a", "e", "i", "o", "u":
		fmt.Println("harf sesli")
	default:
		fmt.Println("harf sessiz")
	}

	//switch kontrol örneği
	switch s1 := sayi(); { //num is not a constant
	case s1 < 50:
		fmt.Printf("%d sayısı 50'den küçük\n", s1)
		fallthrough
	case s1 < 100:
		fmt.Printf("%d sayısı 100'den küçük\n", s1)
		fallthrough
	case s1 < 200:
		fmt.Printf("%d sayısı 200'den küçük", s1)
	}

}
