package main

import "fmt"

func main() {
	sayiList := []int{2, 1, 4, 6, 5, 8, 9, 7}

	sayiMap := make(map[int]int)

	fmt.Println(sayiMap)
	for i := 0; i <= len(sayiList); i++ {
		fmt.Println(i)
	}

}
