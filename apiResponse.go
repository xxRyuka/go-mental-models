package main

import (
	"encoding/json"
	"fmt"
)

type UserDTO struct {
	// 1. İsim Değiştirme: Go'da ID büyük, JSON'da "user_id" küçük
	ID int `json:"id"`

	// 2. Standart Kullanım
	Name string `json:"full_name"`

	// 3. Gizli Alan: Şifre struct'ta var ama JSON'da OLMAYACAK
	Password string `json:"-"`

	// 4. OmitEmpty: Eğer email boşsa, JSON'da yer kaplamasın
	Email string `json:"email,omitempty"`

	// 5. Type Coercion: Sayıyı string olarak gönder (Frontend dostu)
	Balance float64 `json:"balance,string"`
}

func main() {
	user1 := UserDTO{
		ID:       1,
		Name:     "frkan",
		Password: "1234",
		Email:    "ensfurkankus@16",
		Balance:  108.5,
	}

	// Marshaling (Struct -> JSON)
	// C#'taki JsonConvert.SerializeObject()
	jsonData1, _ := json.Marshal(user1)
	fmt.Println("--- Dolu Kullanıcı JSON ---")
	fmt.Println(string(jsonData1)) // output : {"id":1,"full_name":"frkan","email":"ensfurkankus@16","balance":"10"}

	user2 := UserDTO{
		ID:       2,
		Name:     "cabbar",
		Password: "3214",
		Email:    "", //omiempty oldugu için jsona yazılmadı
		Balance:  0,
	}
	jsonData2, _ := json.Marshal(user2)
	fmt.Println("\n--- Eksik Kullanıcı JSON (Email Yok) ---")
	fmt.Println(string(jsonData2)) // {"id":2,"full_name":"cabbar","balance":"0"}

}
