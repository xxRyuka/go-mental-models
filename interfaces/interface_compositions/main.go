package main

import "fmt"

// Yetenekler (interfaces)
type Speaker interface {
	Sound(level int)
}

type Light interface {
	Lumen(level int)
}

type Wifi interface {
	Connect(ip string)
}

// Cihazlar (structs)

type Radio struct {
}

func (r Radio) Connect(ip string) {
	fmt.Println("internete baglandi")
}

func (r Radio) Sound(level int) {
	fmt.Println("ses ")
}

type SmartBulb struct {
}

func (s SmartBulb) Lumen(level int) {
	fmt.Println("isik yandi")
}

func (s SmartBulb) Connect(ip string) {
	fmt.Println("internete baglandi")

}

type Laptop struct {
}

func (l Laptop) Connect(ip string) {
	fmt.Println("internete baglandi")

}

func (l Laptop) Sound(level int) {
	fmt.Println("ses ")

}

func (l Laptop) Lumen(level int) {
	fmt.Println("isik yandi")

}

type SmartDevice interface {
	Light
	Speaker
	Wifi
}

func PlayMedia(ses Speaker) {
	fmt.Println("ses calar")
}

func PartyMode(device SmartDevice) {
	fmt.Println("hepsi")
}

func main() {
	fener := SmartBulb{}
	pc := Laptop{}
	mp3 := Radio{}
	fmt.Println(fener)
	PlayMedia(mp3)

	PartyMode(pc)

	PlayMedia(pc)
}
