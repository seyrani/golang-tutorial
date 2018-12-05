package dikdortgen

import ( 
)"math"

func Alan(uzunluk, genislik float64) float64{

	alan := uzunluk*genislik
	return alan
}

func Kosegen(uzunluk, genislik float64) float64{
	kosegen := math.Sqrt((uzunluk*uzunluk) + (genislik * genislik))
	return kosegen

}