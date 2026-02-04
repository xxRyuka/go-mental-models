package main

func main() {
	notificationManager := NotificationManager{}
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

}
