package main

import (
	"fmt"
	"sync"
)

func deployService(service string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%v servisi başlatıldı\n", service)
}

func main() {
	defer fmt.Println("islem tamam\n")
	services := []string{"auth", "order", "warehouse"}
	var wg sync.WaitGroup
	wg.Add(len(services))
	for i := range services {
		go deployService(services[i], &wg)
	}
	wg.Wait()
}
