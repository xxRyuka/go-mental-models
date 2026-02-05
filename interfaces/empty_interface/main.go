package main

import "fmt"

type Resulution int

const (
	r1024x524 Resulution = iota
	r1080x720
	r1920x1280
)

type ServerStatus string

const (
	Down    ServerStatus = "down"
	Pending ServerStatus = "pending"
	Up      ServerStatus = "up"
)

type Server struct {
	port string
	ip   string
}

type ServerConfig struct {
	fullScreen bool
	MaxPlayers int
	Resulution Resulution
	status     ServerStatus
	info       Server
}

func main() {
	config := map[string]any{
		"tamEkran":   true,
		"maxPlayer":  5,
		"cozunurluk": r1024x524,
		"status":     Pending,
		"info": Server{
			port: "8080",
			ip:   "192.168.1.1",
		},
	}

	for key, value := range config {
		fmt.Printf("key : %v | valueType : %T | value : %+v \n", key, value, value)
		//key : info | valueType : main.Server | value : {port:8080 ip:192.168.1.1}
		//key : tamEkran | valueType : bool | value : true
		//key : maxPlayer | valueType : int | value : 5
		//key : cozunurluk | valueType : main.Resulution | value : 0
		//key : status | valueType : main.ServerStatus | value : pending
	}

	// any oldugu için type assertion yapmadan işlem yapamayız
	//limit := config["maxPlayer"] + 5  invalid operation: config["maxPlayer"] + 5 (mismatched types any and int)

	for key, value := range config {
		switch v := value.(type) {
		case int:
			config[key] = v * 2
			fmt.Println("Yeni oyuncu limiti : ", v*2)
		case bool:
			if v {
				fmt.Println("Tam ekran modu acık")
			} else {
				fmt.Println("tam erkan modu kapalı")
			}
		case Server:
			fmt.Printf("Connecting %v:%v\n", v.ip, v.port)

		case ServerStatus:
			fmt.Printf("Sunucu durumu : %v\n", v)
		default:
			fmt.Printf("kaydedilmemiş tip : %T", v)
		}

	}

	for key, value := range config {
		fmt.Printf("NEW ! || key : %v | valueType : %T | value : %+v \n", key, value, value)
		//NEW ! || key : maxPlayer | valueType : int | value : 10

	}
}
