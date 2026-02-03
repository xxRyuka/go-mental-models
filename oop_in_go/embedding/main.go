package main

import "fmt"

type Yazar struct {
	ad    string
	soyad string
}

func (y *Yazar) TamAd() {
	fmt.Println(y.ad, y.soyad)
}

type Kitap struct {
	Yazar
	baslik string
}

func main() {
	b1 := Kitap{
		Yazar: Yazar{
			ad:    "victor",
			soyad: "hugola",
		},
		baslik: "sefiller",
	}

	b1.TamAd()
}
