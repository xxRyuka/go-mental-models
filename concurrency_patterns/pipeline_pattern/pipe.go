package main

import "fmt"

func GeneratePrices(prices []float64) <-chan float64 {

	priceChan := make(chan float64)

	go func() {
		defer close(priceChan)
		// k , v
		for _, p := range prices {
			priceChan <- p
		}
	}()

	return priceChan
}

func ApplyTax(priceChan <-chan float64) <-chan float64 {

	taxChan := make(chan float64)

	go func() {
		defer close(taxChan)
		for f := range priceChan {
			taxChan <- f * 1.2
		}
	}()
	return taxChan
}

func main() {
	hamFiyatlar := []float64{100.0, 200.0, 50.0}

	priceStream := GeneratePrices(hamFiyatlar)

	withTaxStream := ApplyTax(priceStream)

	for p := range withTaxStream {
		fmt.Printf("KDV'li Fiyat: %.2f\n", p)
	}
}
