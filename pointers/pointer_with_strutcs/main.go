package main

import "fmt"

type Emplooye struct {
	Id     int
	Salary float64
}

func RaiseSalary(e *Emplooye, amount float64) {
	e.Salary += amount
	fmt.Println("maas +++")
}
func main() {
	e1 := &Emplooye{
		Id:     1,
		Salary: 2500,
	}

	// Yani burdan cıkarmam gereken ders:
	//fonksiyona struct verirken değişiklik yapılması gerekiyorsa yine pointer kullan.
	//Çünkü calue olarak verirsen kopyası alınacağı için değişiklik yapılmayacak ?
	fmt.Println(e1.Salary)

	RaiseSalary(e1, 5000)

	fmt.Println(e1.Salary)

}
