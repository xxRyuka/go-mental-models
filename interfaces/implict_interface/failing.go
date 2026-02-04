package main

type FailingSender struct{}

func (FailingSender) Send(msg string) bool {
	return false
}
