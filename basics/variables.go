package basics

import (
	"fmt"
)

// 1. GLOBAL SCOPE (Paket Seviyesi)
// Burada ':=' KULLANAMAYIZ. Zorunlu olarak 'var' kullanıyoruz.
// Bu değişkenler programın yaşam döngüsü boyunca erişilebilir (bu pakette).
var (
	appVersion    string = "1.0.0"
	totalRequests int    //Define etmediğim için go'daki zero Value kuralından suan 0
)

func main() {
	//iota.go file
	var tip myType = HttpStatusForbidden

	fmt.Println(tip)

	// 2. LOCAL SCOPE (Fonksiyon İçi)

	// Senaryo: Bir API isteği geldiğini varsayalım.

	// A. 'var' Kullanımı (Zero Value İhtiyacı)
	// Henüz bir hata oluşmadı, varsayılan olarak 'boş string' olmasını istiyoruz.
	// C#'taki gibi 'string lastError = null' değil, "" olur.
	var lastError string
	// Bir boolean flag. Varsayılanı 'false'tur.
	// "Henüz işlenmedi" anlamına gelir.
	var isProcessed bool

	fmt.Println("--- Başlangıç Durumu (Zero Values) ---")
	fmt.Printf("Toplam İstek: %d\n", totalRequests)               // %d => decimal
	fmt.Printf("Son Hata: '%s' (Burasi bos string)\n", lastError) //%s string
	fmt.Printf("İşlendi mi?: %v\n", isProcessed)                  // %v Value

	// 3. ':=' Kullanımı (Short Declaration)
	// Fonksiyon içinde olduğumuz için ve hemen değer atadığımız için en idiomatic yol budur.
	// Tipini (string) belirtmemize gerek yok, Go anlar (Type Inference).

	reqId := "req-123"
	clientIp := "127.0.0.1"

	totalRequests++
	isProcessed = true

	fmt.Println("\n--- İşlem Sonrası ---")
	fmt.Printf("[%s] ID'li istek %s adresinden alındı.\n", reqId, clientIp)
	fmt.Printf("Güncel Toplam İstek: %d\n", totalRequests)
	fmt.Printf("İşlem Durumu: %v\n", isProcessed)
}
