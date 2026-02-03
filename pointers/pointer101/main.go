package main

import "fmt"

func main() {
	hazine := "altın"
	harita := &hazine

	fmt.Printf("Hazinenin RAM Adresi : %v , Hariadan hazinenin değeri %v\n", harita, *harita)

	*harita = "elmas"

	fmt.Println(hazine)
}
