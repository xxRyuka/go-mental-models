package main

import "fmt"

// Hesapla: İçine ne atarsan at (any), sayıları ayıklayıp toplar.
func Hesapla(veriler []any) float64 {
	var toplam float64 = 0

	fmt.Println("--- Hesaplama Başlıyor ---")

	// Döngü ile slice içindeki her bir 'any' kutusunu geziyoruz.
	for i, veri := range veriler {

		//  TYPE SWITCH (TİP ANAHTARI)
		// 'veri' değişkeni (interface{}) burada analiz ediliyor.
		// 'v' değişkeni ise "Bukalemun"dur. Hangi case'e girerse o tipe dönüşür.
		switch v := veri.(type) {

		case int:
			// BURADA: v değişkeni artık resmen bir 'int'tir.
			// Go'da int ile float64 direkt toplanamaz, o yüzden 'float64(v)' yaptık.
			fmt.Printf("[%d] int bulundu: %d -> Toplama ekleniyor.\n", i, v)
			toplam += float64(v)

		case float64:
			// BURADA: v değişkeni artık bir 'float64'tür.
			// Direkt toplama ekleyebiliriz.
			fmt.Printf("[%d] float bulundu: %.2f -> Toplama ekleniyor.\n", i, v)
			toplam += v

		case string:
			// BURADA: v değişkeni bir 'string'tir.
			// Matematiksel değeri olmadığı için sadece log basıp geçiyoruz.
			fmt.Printf("[%d] String (Yazı) geldi, atlanıyor: \"%s\"\n", i, v)

		case bool:
			// BURADA: v değişkeni bir 'bool'dur.
			// Senaryo: Eğer değer 'false' ise "İşlemi bitir" komutu sayıyoruz.
			if !v {
				fmt.Printf("[%d]  DUR (false) sinyali alındı! Hesaplama kesiliyor.\n", i)
				return toplam // Fonksiyondan anında çıkar ve o anki toplamı döner.
			}
			fmt.Printf("[%d] bool (true) geldi, yola devam ediliyor.\n", i)

		default:
			// Hiçbirine uymazsa (örneğin struct gelirse) burası çalışır.
			// %T formatı, verinin orijinal tipini gösterir.
			fmt.Printf("[%d] Bilinmeyen tip: %T, işlem yapılmadı.\n", i, v)
		}
	}

	return toplam
}

func main() {
	// 1. Senaryo Gereği Karışık Veri Seti Oluşturuyoruz ([]any)
	// İçinde int, float, string ve bool var.
	karisikVeriler := []any{
		10,          // int (Toplanacak)
		25.50,       // float64 (Toplanacak)
		"HATA_LOGU", // string (Atlanacak)
		5,           // int (Toplanacak)
		true,        // bool (Devam et)
		"System32",  // string (Atlanacak)
		false,       // bool (STOP Sinyali! Buradan sonrasına bakılmamalı)
		1000,        // int (Buna asla ulaşılmamalı çünkü üstte false var)
	}

	// 2. Fonksiyonu Çağırıyoruz
	sonuc := Hesapla(karisikVeriler)

	// 3. Sonucu Yazdırıyoruz
	fmt.Println("--------------------------------")
	fmt.Printf(" İŞLEM SONUCU: %.2f\n", sonuc)

	//	--- Hesaplama Başlıyor ---
	//		[0] int bulundu: 10 -> Toplama ekleniyor.
	//	[1] float bulundu: 25.50 -> Toplama ekleniyor.
	//	[2] String (Yazı) geldi, atlanıyor: "HATA_LOGU"
	//	[3] int bulundu: 5 -> Toplama ekleniyor.
	//	[4] bool (true) geldi, yola devam ediliyor.
	//	[5]	String (Yazı) geldi, atlanıyor: "System32"
	//	[6]  DUR (false) sinyali alındı! Hesaplama kesiliyor.
	//	--------------------------------
	//	İŞLEM SONUCU: 40.50

}
