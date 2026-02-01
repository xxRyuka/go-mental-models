package main

import "strings"

// Named Return Values
func SplitCoordinates(location string) (city string, country string) {
	city, country, found := strings.Cut(location, "-")
	if found {
		return //city, country yazmamıza gerek yok çünkü named return values kullandık

	} else {
		return
	}
}

func main() {
	city, country := SplitCoordinates("Istanbul-Turkey")
	println("City:", city)
	println("Country:", country)
}
