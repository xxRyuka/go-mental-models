package main

import "fmt"

func kelimeEkle(dict map[string]string, k, v string) {
	dict[k] = v
}
func main() {
	sozluk := map[string]string{
		"kitap":  "book",
		"pencil": "kalem",
	}

	fmt.Println(sozluk)

	kelimeEkle(sozluk, "furkan", "kus")
	fmt.Println(sozluk)
}
