//rectprops.go
package rectangle

import (
	"fmt"
	"math"
)

//ekrana yazı basan basit bir init fonksiyonu
func init() {

	fmt.Println("Rectangle paketi yüklendi")

}

func Alan(len, wid float64) float64 {
	alan := len * wid
	return alan
}

func Kosegen(len, wid float64) float64 {
	kosegen := math.Sqrt((len * len) + (wid * wid))
	return kosegen
}
