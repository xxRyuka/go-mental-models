package main

import (
	"fmt"
	"sync"
	"time"
)

func JobGenerator(count int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for i := 0; i <= count; i++ {
			out <- fmt.Sprintf("Job No : %v", i)
		}
	}()

	return out
}

func Woker(workerID int, jobCh <-chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for ch := range jobCh {
		fmt.Printf("Worker : %v Calisiyor, %v Uzerinde Calisiyor\n", workerID, ch)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker : %v , %v Görevini Tamamladi\n", workerID, ch)

	}
}

func main() {
	fmt.Println(time.Now().Second())

	var wg sync.WaitGroup
	jobs := JobGenerator(15)

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go Woker(i, jobs, &wg)
	}

	wg.Wait()
	fmt.Println(time.Now().Second())
}
