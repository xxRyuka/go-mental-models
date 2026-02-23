package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func BackgroundWorker(wg *sync.WaitGroup, ctx context.Context, id int) {
	defer wg.Done()
	i := 0
	for {
		select {

		case <-ctx.Done():
			{
				fmt.Printf("Worker %v: Kapatılma sinyali alındı, temizlik yapılıyor... \n", id)
				wg.Done()
				return
			}

		case <-time.After(time.Millisecond * 500):
			{
				i++
				fmt.Printf("Worker %v: Veri İşliyor... rotate : %v\n", id, i)
			}

		}
	}
}

func main() {
	ctx, cancFunc := context.WithCancel(context.Background())
	defer cancFunc()

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go BackgroundWorker(&wg, ctx, i)
	}
	time.Sleep(2 * time.Second)
	cancFunc()
	wg.Wait()
	fmt.Println("Sistem güvenle kapatıldı.")
}
