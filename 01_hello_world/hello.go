package main

import "fmt"

func Hello(receiver string) string {
	return fmt.Sprintf("Hello, %s!", receiver)
}
