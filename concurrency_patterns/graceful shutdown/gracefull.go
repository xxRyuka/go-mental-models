package main

import (
	"context"
	"fmt"
	"sync"
)

func jobGenerator(ctx context.Context, count int) <-chan int {
	jobs := make(chan int)
	go func() {
		defer close(jobs) // go funcun içinde olmalı dısarda olursa direk kapanır ve kapalı kanala veri gondermeeye calısırz
		for i := 1; i <= count; i++ {

			select {
			case jobs <- i:
			case <-ctx.Done():
				fmt.Println("[JobGenerator] İptal Edildi Üretim Duruyor")
				return
			}

		}
	}()
	return jobs
}

// wg yok dedik ya wg hangi durumda kullanılır hangi durumda kullanımaz ? bunuda bir netlestirmek istiyorum
func worker(ctx context.Context, workerID int, jobs <-chan int) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for {

			select {

			case <-ctx.Done():
				fmt.Printf("[Worker] {%v} Veri Beklerken İptal Edildi Üretim Duruyor\n", workerID)
				return

			case job, ok := <-jobs:
				{
					if ok != true {
						return
					}

					select {
					case out <- fmt.Sprintf("Gorev %v , Worker : %v tarafından Tamamlandi", job, workerID):

					case <-ctx.Done():
						fmt.Printf("[Worker] {%v} Veri Gonderirken İptal Edildi Üretim Duruyor\n", workerID)
						return
					}
				}

			}

		}
	}()
	return out
}

// Multiplexter Merge
// variadic kw "..."
func Multiplexter(chans ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	wg.Add(len(chans))

	for _, strings := range chans {

		go func(ch <-chan string) {
			for s := range ch {
				out <- fmt.Sprintf("Multiplexed : %v", s)
			}
			wg.Done()
		}(strings)

	}

	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	jobs := jobGenerator(ctx, 105)
	worker1 := worker(ctx, 1, jobs)
	worker2 := worker(ctx, 2, jobs)
	worker3 := worker(ctx, 3, jobs)

	Multiplexed_Results := Multiplexter(worker1, worker2, worker3)

	counter := 1
	for sonuc := range Multiplexed_Results {
		fmt.Println(sonuc)
		counter++
		if counter > 5 {
			cancel()
			break
		}
	}
}
