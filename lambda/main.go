package main

import "fmt"

type MyEvent struct {
	Username string `json:"username"`
}

// * Take in a payload and do something with it
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username is empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {

}
