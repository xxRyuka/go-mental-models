package main

import "fmt"

//C#'ta metotların başına public yazmazsan varsayılan olarak private olur ya?
//Go'da da fonksiyonun baş harfini küçük yazarsan (greeting), o fonksiyon private (unexported) olur.

func greeting(name, surname string) string {
	return "Hi" + name + " " + surname
}

func greeting2(name string, surname string) string {
	return "Hi" + name + " " + surname
}

func greeting3(param string) {
	fmt.Println(param)
}
func main() {
	ifade := greeting("John", "Mary")
	ifade2 := greeting2("Jane", "Doe")
	fmt.Println(ifade)
	fmt.Println(ifade2)
	greeting3("Hello World")
}
