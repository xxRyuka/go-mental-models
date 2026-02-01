package composites

import "fmt"

func main() {
	// Dizi Oluturuyorum
	var Dizi [5]int = [5]int{1, 2, 3, 4, 5}

	mySlice := Dizi[1:4]
	fmt.Println(" mySlice ilk hali", mySlice)
	fmt.Println("Dizinin İlk hali ", Dizi)

	mySlice[2] = 2

	// Sliceler arrayin pointeridir slice değişirse Arrayde değişir
	fmt.Println(" mySlice son hali", mySlice)
	fmt.Println("Dizinin son hali ", Dizi)

	//make([]T, length, capacity)
	slice2 := make([]int, 0, 2)
	fmt.Println("len , cap,  &slice2 , slice2 : ", len(slice2), cap(slice2), &slice2, (slice2))
	slice2 = append(slice2, 1)
	slice2 = append(slice2, 2)
	//slice2 = append(slice2, 8)
	slice3 := append(slice2, 5) // çağıranı ETKİLEMEZ yeni bir slice olusturuyor

	fmt.Println("3 append yapıldı 2 kapasileli sliceye")
	fmt.Println("len , cap,  &slice2 , slice2 : ", len(slice2), cap(slice2), &slice2, (slice2))

	fmt.Println(" slice3 : len , cap,  &slice3 , slice3 : ", len(slice3), cap(slice3), &slice3, (slice3))

	fmt.Println("--- Ayrılık Testi ---")
	// slice3'ün ilk elemanını değiştiriyorum (Eskiden 1 idi)
	slice3[0] = 999
	// Slice kapasitesi değişmezse ayni arraya bakmaya devam edeceği için ikisinide etkiler ama
	//cpasity değişirse artık farklı arraylara baktiklari için ayrilmislar demektir
	fmt.Println("Slice3 (999 olmalı):", slice3)
	fmt.Println("Slice2 (1 mi kaldı?):", slice2)
}
