package main

type FailingSender struct{}

func (s FailingSender) GetProvider() string {
	return "failProvider"
}

func (FailingSender) Send(msg string) bool {
	return false
}
