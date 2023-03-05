package main

import (
	"fmt"

	"github.com/numbertheory/chatgpt-go/chat"
)

const usage = `Usage of ChatGPT:
`

func main() {
	for {
		userInput := chat.StringPrompt(">")
		if userInput == "exit" {
			break
		}
		fmt.Printf(userInput + "\n")
	}

}
