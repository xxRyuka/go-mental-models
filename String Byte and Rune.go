package main

import (
	"fmt"
)

func main() {

	fmt.Println("--- BÖLÜM 1: ASCII (İngilizce) Manipülasyonu ---")
	s1 := "cat"
	fmt.Println("phase1 ", s1)
	// String doğrudan değiştirilemez. Önce ameliyat masasına (Byte Slice) yatırıyoruz.
	// BU İŞLEM BELLEK KOPYALAMASI (ALLOCATION) YAPAR!
	s1Bytes := []byte(s1)                  // bytze dizisine alıyorum
	fmt.Println("Orijinal Byte:", s1Bytes) // [99 17 116] (ASCII kodları)
	s1Bytes[1] = 's'

	fmt.Println("Guncellenmis Byte:", s1Bytes) // [99 115 116] (ASCII kodları)

	s1Yeni := string(s1Bytes)
	fmt.Println("phase2 Guncellenmis ifade : ", s1Yeni)

	fmt.Println("\n--- BÖLÜM 2: Unicode (Türkçe) Tuzağı ---")

	s2 := "çilek"
	bytes2 := []byte(s2)
	// BURAYA DİKKAT!
	// 'ç' harfi bellekte KAÇ BYTE yer kaplıyor?
	// Go (UTF-8) 'ç' harfini 2 byte ile kodlar: [195 167]
	fmt.Println("'çilek' Byte Dizisi:", bytes2)
	// Biz saf bir niyetle 0. indeksi (sadece ilk byte'ı) 'd' yapıyoruz.
	// Ama 'ç' harfinin 2. byte'ı (167) orada sahipsiz kalacak!
	bytes2[0] = 'd'

	s2Bozuk := string(bytes2)
	fmt.Println("Bozuk Sonuç:", s2Bozuk)
	// Çıktıda 'd' harfinin yanında garip bir sembol göreceksin.
	// Çünkü 'd' (1 byte) + Artık kalan yarım 'ç' parçası = BOZUK UTF-8

	fmt.Println("\n--- BÖLÜM 3: Doğru Yöntem (Rune) ---")

	// Eğer karakter bazlı değişim yapacaksak Byte değil Rune kullanmalıyız.
	runes := []rune("çilek")

	runes[0] = 'd'

	s2Degisen := string(runes)
	fmt.Println("Degisen Rune : ", s2Degisen)
}
