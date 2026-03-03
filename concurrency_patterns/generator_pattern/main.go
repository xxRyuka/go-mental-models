package main

import (
	"context"
	"fmt"
	rand2 "math/rand/v2"
	"time"
)

//Verileri direk bir slice,List... olarak tutup rami sisirmek yerine Generator Pattern ile pipe içinde tutacağız

func GenerateFakeTransaction(ctx context.Context, count int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for i := 0; i < count; i++ {

			select {
			case out <- fmt.Sprintf("TXN-LOG-%vFF-%v", rand2.IntN(8975), i):
			case <-ctx.Done():
				{
					fmt.Println("[GENERATOR] Tüketici okumayı bıraktı,üretim durduruluyor")
					return
				}
			}
		}
	}()

	return out
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream := GenerateFakeTransaction(ctx, 150)

	counter := 1
	for t := range stream {
		fmt.Println(t)
		counter++
		if counter == 5 {
			cancel()
			break
		}
	}

	time.Sleep(1 * time.Second) // Printleri görmek için ufak bir bekleme
}
