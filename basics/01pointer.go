package basics

import "fmt"

func main() {
	sayi := 50
	fmt.Println("sayi: ", sayi)
	pointer := &sayi
	fmt.Println("pointer: ", pointer)
	*pointer = 100
	fmt.Println("pointer: ", pointer)

	fmt.Println("sai ", sayi)
}
