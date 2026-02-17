package main

import (
	"fmt"
	rand2 "math/rand/v2"
	"time"
)

// timeout and voidable API
func GetNews(newsChannel chan<- string) {

	rand := rand2.IntN(5)
	time.Sleep(time.Duration(rand) * time.Second)
	newsChannel <- fmt.Sprintf("Haber %v \n", &newsChannel)
}
func main() {
	newsCh := make(chan string)

	go GetNews(newsCh)

	select {
	case response := <-newsCh:
		fmt.Printf("Basarili %v", response)

	case <-time.After(3 * time.Second):
		fmt.Println("Zamanasimi")
	}
}

//Zamanasimi
//Haber 0xc000088070
