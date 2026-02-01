package main

import "fmt"

func Divide(dividend, divisor float64) (float64, string) {
	if divisor == 0 {
		return 0, "Sifira bolme hatasi"
	}
	return float64(dividend / divisor), "Basarili"
}

func main() {
	_, msg := Divide(10, 0)
	fmt.Println(msg)

	result, msg := Divide(10, 2)
	fmt.Println(result, msg)

	results, msg3 := Divide(10, 50)
	fmt.Println(results, msg3)

	fmt.Printf("%T\n", results)
}
