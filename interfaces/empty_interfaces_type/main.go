package main

import (
	"fmt"
)

// event structlari
type UserLogin struct {
	uName    string
	succsess bool
}

type Payment struct {
	amount   int
	currency string
}
type SystemErr struct {
	msg  string
	code int
}

type DbErr struct {
	cluster string
}

func WriteLogVariadic(vars ...any) {
	for _, v := range vars {
		switch v := v.(type) {
		case UserLogin:
			fmt.Printf("kullanıcı giris olayi %+v\n", v)

		case Payment:
			fmt.Printf("odeme olayi %+v | tutar %v , doviz %v\n", v, v.amount, v.currency)

		case SystemErr:
			fmt.Printf("error olayi %+v\n", v)
		default:
			fmt.Printf("bilinmeyen tip type : %T , value %+v\n", v, v)

		}
	}
}

func WriteLog(events []any) {

	for _, v := range events {
		switch v := v.(type) {
		case UserLogin:
			fmt.Printf("kullanıcı giris olayi %+v\n", v)

		case Payment:
			fmt.Printf("odeme olayi %+v | tutar %v , doviz %v\n", v, v.amount, v.currency)

		case SystemErr:
			fmt.Printf("error olayi %+v\n", v)
		default:
			fmt.Printf("bilinmeyen tip type : %T , value %+v\n", v, v)

		}
	}
}

func main() {
	eventUL := UserLogin{
		uName:    "cbbr",
		succsess: false,
	}
	eventPayment := Payment{
		amount:   20,
		currency: "USD",
	}
	eventSE := SystemErr{
		msg:  "nullReference ",
		code: 20,
	}

	eventDE := DbErr{
		cluster: "as",
	}
	mixedEvents := []any{eventSE, eventUL, eventPayment, eventDE}
	WriteLog(mixedEvents)

	//error olayi {msg:nullReference  code:20}
	//kullanıcı giris olayi {uName:cbbr succsess:false}
	//odeme olayi {amount:20 currency:USD} | tutar 20 , doviz USD
	//bilinmeyen tip type : main.DbErr , value {cluster:as}

	// fonksiyonu variadic olarak tanımladim ve direk parametre verilebilir hale getirdim her tipten
	WriteLogVariadic(eventSE, eventUL, eventPayment, eventDE, 5, "selam")
}
