package main

import "fmt"

type SlackSender struct {
	Channel string
	info    string
}

func (ss SlackSender) GetProvider() string {
	return "slackProvider"
}

func (ss SlackSender) Send(msg string) bool {

	fmt.Printf("%v kanalina %v mesaji iletildi\n", ss.Channel, msg)
	return true
}
