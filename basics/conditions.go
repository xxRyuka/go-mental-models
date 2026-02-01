package basics

import "fmt"

func main() {
	role := "admin"

	switch role {
	case "admin":
		fmt.Println("Adminsin go")
		// break; <-- GEREK YOK! Go otomatik çıkar.

	case "guest":
		fmt.Println("goruntuleme mdu")

	case "editor", "writer":
		fmt.Println("yazabilirsn")

	default:
		fmt.Println("sen burdan gecemen")
	}

	if isTokenValid := true; role == "admin" && isTokenValid {
		fmt.Println("2FA Başarılı! Hoşgeldin Süper Admin")
	} else {
		fmt.Println("2FA Hatası veya Yetki Yok")
	}

}
