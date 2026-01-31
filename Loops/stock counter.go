package main

import (
	"fmt"
	"sort"
)

func main() {
	stoklar := map[string]int{
		"Laptop":  10,
		"Telefon": 25,
		"Tablet":  0, // Stok bitmiş
		"Monitör": 8,
		"Klavye":  30,
		"Fare":    0, // Stok bitmiş
		"Yazıcı":  5,
	}

	fmt.Println("--- 1. TÜM ENVANTER (Key & Value) ---")

	// C#: foreach (var item in stoklar) { item.Key, item.Value }
	// Go: k (key), v (value) := range map
	for urun, adet := range stoklar {
		fmt.Printf("Ürün: %s, Adet: %d\n", urun, adet)
		// NOT: Çıktı sırası her çalıştırdığında değişebilir! cünkü mapler siralı değildir.
	}

	fmt.Println("\n--- 2. STOKTA OLMAYANLAR (Sadece Key İle Yazildi Dongu) ---")
	for urun := range stoklar {
		if stoklar[urun] == 0 {
			fmt.Printf("Stokta Yok: %s\n", urun)
		}
	}

	fmt.Println("\n--- 3. TOPLAM ÜRÜN SAYISI (Sadece Value Lazım) ---")

	toplamUrun := 0
	for _, adet := range stoklar {

		toplamUrun += adet
	}
	fmt.Printf("Toplam Stoktaki Ürün Adedi: %d\n", toplamUrun)

	fmt.Println("Mapi Sıralı Hale Getireceğiz ")

	keys := make([]string, 0, len(stoklar))

	for k := range stoklar {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	// mapten değerleri çekip yazdir ne onu anlamadım
	for _, k := range keys {
		fmt.Printf("Ürün: %v, Adet: %v\n", k, stoklar[k])
	}
}
