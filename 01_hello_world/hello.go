package main

const enHelloPrefix = "Hello, "

func Hello(receiver string) string {
	if receiver == "" {
		receiver = "world"
	}
	return enHelloPrefix + receiver
}
