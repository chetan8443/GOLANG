package part1

import "fmt"

func HelloGreeting(name string) string {
	message := fmt.Sprintf("Hello %v", name)
	return message
}
