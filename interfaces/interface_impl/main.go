package main

import (
	"fmt"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func EkranaYaz(w Writer, mesaj string) {
	w.Write([]byte(mesaj))
}

type customWriteStruct struct {
	version string
}

// * koyunca neden
// Cannot use 'cws1' (type customWriteStruct) as the type WriterType does not implement 'Writer' as the 'Write' method has a pointer receiver
// hatasi veriyor ?
func (cws *customWriteStruct) Write(p []byte) (n int, err error) {
	fmt.Printf("\n %v Custom Write Struct write fonksiyonu calısıyor, %v", cws.version, string(p))
	return len(p), nil
}

func main() {
	obj1 := os.Stdout
	EkranaYaz(obj1, "anlamiyorum iste verdiğin bu ornekte naptik simdi ya ben her paket eklediğimde kaynak koda mı gitcem ? ayrıca stdout kaynak koduna gittiğimde write fonksiyonnu aldiği parametreyi goremedim bile ??")

	cws1 := customWriteStruct{version: "80"}

	EkranaYaz(&cws1, "bos")
}
