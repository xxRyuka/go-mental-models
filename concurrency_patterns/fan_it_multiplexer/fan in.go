package main

import (
	"fmt"
	"sync"
)

func jobGenerator(count int) <-chan int {
	jobs := make(chan int)
	go func() {
		defer close(jobs) // go funcun içinde olmalı dısarda olursa direk kapanır ve kapalı kanala veri gondermeeye calısırz
		for i := 1; i <= count; i++ {
			jobs <- i
		}
	}()
	return jobs
}

// wg yok dedik ya wg hangi durumda kullanılır hangi durumda kullanımaz ? bunuda bir netlestirmek istiyorum
func worker(workerID int, jobs <-chan int) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for job := range jobs {
			out <- fmt.Sprintf("Gorev %v , Worker : %v tarafından Tamamlandi", job, workerID)
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
	jobs := jobGenerator(15)
	worker1 := worker(1, jobs)
	worker2 := worker(2, jobs)
	worker3 := worker(3, jobs)

	Multiplexed_Results := Multiplexter(worker1, worker2, worker3)

	for sonuc := range Multiplexed_Results {
		fmt.Println(sonuc)
	}
}
