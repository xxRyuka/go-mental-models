package basics

import (
	"fmt"
	"math"
)

func main() {
	var a int = 3
	var b int = 4
	// HATA: math.Sqrt fonksiyonu parametre olarak 'float64' ister.
	// C# olsa 'int'i otomatik 'double'a çevirirdi. Go yapmaz!
	// k := math.Sqrt(a*a + b*b)
	k := math.Sqrt(float64(a*a + b*b))

	// f şu an float64 (5.0). Bunu uint'e çevirmek istiyoruz.
	// z := uint(f) // Yine açıkça yazmalısın.
	var z uint = uint(k)

	fmt.Printf("%v", z)
}
