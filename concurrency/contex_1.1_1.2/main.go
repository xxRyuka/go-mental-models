package main

import (
	"context"
	"fmt"
	"time"
)

func GetUserProfile(ctx context.Context, id int) (string, error) {

	chanDB := make(chan string)

	go PullDb(chanDB)
	select {
	case res := <-chanDB:
		fmt.Printf("islem basarili Response : '%v' \n", res)
		return res, nil
	case <-ctx.Done():
		fmt.Println("Zaman Asimi !! ")
		return "", ctx.Err()
	}
}

func PullDb(chanDb chan<- string) {
	time.Sleep(2 * time.Second)
	chanDb <- "ryuka"

}

func main() {
	ctxB, canc := context.WithTimeout(context.Background(), 3*time.Second)
	defer canc() // Bunu Defer ile 2. kez cagirmakta sorun yok
	fmt.Printf("ctx : %+v ,\n\ncanc: %+v\n\n\n\n", ctxB, canc)

	fmt.Println("Veritabanına istek atılıyor (Max süre: 3sn)...")

	result, err := GetUserProfile(ctxB, 5)
	canc() // Bunu Defer ilede cagırabilirdik
	// ama fonskyion devam edecek olursa +15+20 saniye orneğin deferi beklemesine gerek kalmayacak.
	//fonksiyon tamamlandıktan sonra contexti bitirecek !
	if err != nil {
		fmt.Printf("Err: %e \n", err.Error())
		return
	}
	fmt.Printf("Veri iste burda : %v \n", result)

}
