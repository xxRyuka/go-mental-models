package main

import "fmt"

type Ogrenci struct {
	Ad    string
	Vize  int
	Final int
}

func (o Ogrenci) OrtalamaHesapla() float64 {
	return float64(o.Vize*1/3) + float64(o.Final)*2/3
}

func (o Ogrenci) isPassed(baraj float64) bool {
	ort := o.OrtalamaHesapla() // bu verdiğimiz o parametresini dısardan nasıl alıyor dilin inceliklerinide ogrenmek istiyorum
	if ort > baraj {
		return true
	} else {
		return false
	}
}

func main() {
	ogrenci1 := Ogrenci{
		Ad:    "Ahmet",
		Vize:  70,
		Final: 90,
	}

	ortalama := ogrenci1.OrtalamaHesapla()

	fmt.Printf(" %q isimli ogrencinin ortalamasi: %f \n", ogrenci1.Ad, ortalama)

	ogrenci2 := Ogrenci{
		Ad:    "asyse",
		Vize:  50,
		Final: 65,
	}

	ortalama2 := ogrenci2.OrtalamaHesapla()

	fmt.Printf(" %q isimli ogrencinin ortalamasi: %f \n", ogrenci2.Ad, ortalama2)

	fmt.Printf("%t", ogrenci1.isPassed(50))
}
