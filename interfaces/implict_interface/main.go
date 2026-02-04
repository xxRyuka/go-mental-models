package main

func main() {
	//Artık bu kullanılamaz çünkü panic: assignment to entry in nil map hatası alırız
	//notificationManager := NotificationManager{} //
	notificationManager := NewManager() // constuctor fonksiyonunu cagiriyoruz
	emailSender1 := &EmailSender{
		"sa", 50,
	}
	slackSender1 := SlackSender{Channel: "toplanti odası"}
	smsSender1 := SmsSender{PhoneNumber: "55555"}

	notificationManager.AddProvider(FailingSender{})

	notificationManager.AddProvider(emailSender1)
	notificationManager.AddProvider(slackSender1)
	notificationManager.AddProvider(smsSender1)

	notificationManager.Broadcast("sirketi kapatiyoruz")

	notificationManager.SendByProviderId(smsSender1.GetProvider(), "selam")

	notificationManager.SendByProviderId("TEST OLMAYAN PROVİDER", "das")

}
