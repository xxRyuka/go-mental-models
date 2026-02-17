package main

import (
	"fmt"
	"sync"
)

func add(wg *sync.WaitGroup, num *int, mtx *sync.Mutex) {
	defer wg.Done()
	mtx.Lock()
	*num = *num + 1
	mtx.Unlock()
}

func main() {
	wg := new(sync.WaitGroup)
	mtx := new(sync.Mutex)
	num := 0
	wg.Add(10000)

	for i := 1; i <= 10000; i++ {
		go add(wg, &num, mtx)
	}
	wg.Wait()

	fmt.Println(num)

	// out : 9774

}
