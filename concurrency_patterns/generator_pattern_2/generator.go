package main

import (
	"context"
	"fmt"
)

func Generator(ctx context.Context, count int) <-chan string {

	out := make(chan string)

	go func() {
		defer close(out)
		for i := 0; i <= count; i++ {

			data := fmt.Sprintf("Data-FF-%v", i)

			out <- data
			//select {
			//case out <- data:
			//	{
			//		fmt.Printf("ok %v\n", i)
			//	}
			//case <-ctx.Done():
			//	fmt.Println("[GENERATOR] Tüketici fişi çekti, üretim derhal durduruluyor!")
			//	return
			//}
		}
	}()
	return out
}
func main() {

	ctx, canc := context.WithCancel(context.Background())
	defer canc()
	stream := Generator(ctx, 1001)
	for data := range stream {
		fmt.Println(data)
		if data == "Data-FF-5" {
			//canc()
			break

		}
	}

	//time.Sleep(1 * time.Second)
}
