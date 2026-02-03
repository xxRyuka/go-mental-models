package main

import "fmt"

func addItem(slice []string, item string) []string {
	return append(slice, item)
}

// Bunu alamadım parametre olarak *slice veriyoruz yani = 06cx568
// appendin içinede *slice'nin pointerini yani **slice veriyoruz = 06ax875 veriyoruz burda kafam karıstı gercekten devam ettirmedim burayı netleştirmek istiyorum
func addItemWithPointer(slice *[]string, item string) {
	fmt.Println("pointer ile add slice fonlsioynu calisyi")
	*slice = append(*slice, item)
}
func main() {
	envanter := []string{"canta", "terlik"}
	fmt.Println(envanter)

	x := addItem(envanter, "cabbar")
	fmt.Println(x, "x")

	fmt.Printf("%v , pointer : %p \n", envanter, &envanter)

	addItemWithPointer(&envanter, "newitem")

	fmt.Printf("%v , pointer : %p", envanter, &envanter)

	// Adresin değişmiyor olmasının sebebi sanırım su biz pointer ekleme fonksiyonunda adresin referans gösterdiği arrayi değiştirdik değil mi :D

	//
	//	[canta terlik]
	//	[canta terlik cabbar] x
	//	[canta terlik] , pointer : 0xc00009a090
	//	pointer ile add slice fonlsioynu calisyi
	//	[canta terlik newitem] , pointer : 0xc00009a090

}
