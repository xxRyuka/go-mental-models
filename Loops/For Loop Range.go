package main

import "fmt"

func main() {
	// 1. DÜZELTME: Bozuk olanı öne aldık ki 'break' olmadan önce test edebilelim.
	urls := []string{"api.bozuk.com", "api.google.com", "api.github.com", "s"}

ScanAllUrls:
	for _, url := range urls {
		fmt.Println("\n--------------------")
		fmt.Println("URL:", url)

		// 2. DÜZELTME: '<= 3' yaptık ki gerçekten 3 kere denesin.
		for deneme := 1; deneme <= 3; deneme++ {

			// 3. DÜZELTME: '\n' ekledik, çıktılar alt alta olsun.
			fmt.Printf("   -> Deneme Sayısı: %d\n", deneme)

			if url == "api.bozuk.com" {
				fmt.Println("       Sunucu Cevap Vermiyor (Retry Bekleniyor...)")
				// continue demeye gerek yok, loop zaten döner.
			} else if url == "api.github.com" {
				fmt.Println("       HEDEF BULUNDU! Tüm operasyon durduruluyor.")
				break ScanAllUrls // FİŞİ ÇEK!
			} else {
				fmt.Println("       Bağlantı Başarılı.")
				break // Sadece bu URL'i geç, sıradakine bak.
			}
		}
	}
	fmt.Println("\n Program Bitti.")
}
