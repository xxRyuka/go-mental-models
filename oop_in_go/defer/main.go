package main

import "fmt"

func DbProcess() {
	fmt.Println("conn open")

	defer fmt.Println("conn closed")

	fmt.Println("db migrating")

	if true {
		fmt.Println("(error simulating) an Error occured")
		return
	}
}

func main() {
	DbProcess()
	// output :
	//conn open
	//db migrating
	//(error simulating) an Error occured
	//conn closed

}
