package main

import "fmt"

type Character struct {
	name    string
	health  int
	stamina int
}

func (c Character) HasarAl_ValueReceiver(dmg int) int {
	c.health -= dmg
	fmt.Printf("%v karakteri %v hasari aldi !", c.name, dmg)
	return c.health
}

func (c *Character) HasarAl_PointerReceiver(dmg int) int {
	c.health -= dmg
	fmt.Printf("%v karakteri %v hasari aldi ! ", c.name, dmg)
	return c.health
}

// Burda pointer receiver'da kullanabilirdik
// ama bi veri değişikliği söz konusu olmadıgı için ve structure kucuk oldugu için bu sekildede kullanilabilir
func (c Character) GetInfo() { // ÖNERİLEN DURUM HEPSİNDE POINTER KULLANMAKTIR
	fmt.Printf("\nCurrent HP(from getInfo func)  : %v\n", c.health)
}
func main() {

	c1 := Character{
		name:    "sam",
		health:  100,
		stamina: 20,
	}

	fmt.Printf("%v cani : %v ", c1.name, c1.health)
	exceptedHp := c1.HasarAl_ValueReceiver(80)
	fmt.Println("Value Receiver Func out: ")
	fmt.Printf("\nCurrent HP : %v\nExcepted Hp : %v", c1.health, exceptedHp)
	c1.GetInfo()

	fmt.Println("----")

	exceptedHp2 := c1.HasarAl_PointerReceiver(30)
	fmt.Println("Pointer Receiver Func out: ")
	fmt.Printf("\nCurrent HP : %v\nExcepted Hp : %v", c1.health, exceptedHp2)
	c1.GetInfo()

}
