package composites

import "fmt"

func main() {

	// map[keyType]valueType => Go arkada bir hash table olusturdu
	serviceLatencies := make(map[string]int)

	serviceLatencies["AuthService"] = 45
	serviceLatencies["PaymentService"] = 150
	serviceLatencies["OrderService"] = 100
	serviceLatencies["NotifyService"] = 800

	// 3. TABLO FORMATINDA YAZDIRMA (Prefix/Suffix Kullanımı)
	// %-20s : Sola yasla, 20 karakter yer aç (Kalanı boşluk doldurur)
	// %10s  : Sağa yasla, 10 karakter yer aç
	fmt.Println("\n------------------------------------")
	fmt.Printf("| %-20s | %10s |\n", "SERVICE NAME", "LATENCY (ms)")
	fmt.Println("------------------------------------")

	// Elle tek tek yazdırıyoruz (Döngüye henüz girmiyoruz)
	// AuthService
	fmt.Printf("| %-20s | %10d |\n", "AuthService", serviceLatencies["AuthService"])
	// PaymentService
	fmt.Printf("| %-20s | %10d |\n", "PaymentService", serviceLatencies["PaymentService"])
	// OrderService
	fmt.Printf("| %-20s | %10d |\n", "OrderService", serviceLatencies["OrderService"])

	fmt.Println("\n[INFO] PaymentService optimize edildi...")
	serviceLatencies["PaymentService"] = 55 // 150'den 55'e düştü

	// %q kullanımı: String'i tırnaklı gösterir.
	fmt.Printf("Güncel Durum -> Servis: %q, Yeni Latency: %v\n", "PaymentService", serviceLatencies["PaymentService"])

	// 5. OLMAYAN VERİ KONTROLÜ (Comma-ok Idiom)
	targetService := "EmailService"
	fmt.Printf("\n[CHECK] %s durumu kontrol ediliyor...\n", targetService)

	// Go burada Exception fırlatmaz, latency 0 gelir. Ama servis var mı yok mu?
	latency, isOnline := serviceLatencies[targetService]

	if isOnline {
		fmt.Printf(" %s Aktif! Gecikme: %dms\n", targetService, latency)
	} else {
		// %T ile tipini görelim, %v ile değerini (false) görelim.
		fmt.Printf(" %s Bulunamadı! (isOnline tipi: %T, değeri: %v)\n", targetService, isOnline, isOnline)
	}

	//  SİLME İŞLEMİ (Delete)
	fmt.Println("\n[ALERT] OrderService kapatılıyor...")
	delete(serviceLatencies, "OrderService")

	// Silineni kontrol edelim
	_, varMi := serviceLatencies["OrderService"]
	fmt.Printf("OrderService hala var mı? : %t\n", varMi) // %t bool için özeldir
	serviceName := "SearchService"                        // İsmi değişkene alalım ki aşağıda kullanalım
	serviceLatencies[serviceName] = 15

	search, isOk := serviceLatencies["SearchService"]

	if isOk && search < 25 {
		fmt.Printf("[%s] cok hizliyiz %d\n", serviceName, search)
	}
}
