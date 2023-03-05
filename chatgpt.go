package main

import (
	"fmt"
	"os"

	"github.com/numbertheory/chatgpt-go/chat"
)

const usage = `Usage of ChatGPT:
`
const tokenNotSet = `Chat GPT token required.
Visit https://platform.openai.com/account/api-keys to get keys for your account.
Then, make that key an environmental variable in your shell.
Example:
    export CHATGPT_TOKEN=<your-token>
`

func main() {
	token := os.Getenv("CHATGPT_TOKEN")
	if token == "" {
		fmt.Printf(tokenNotSet)
		os.Exit(1)
	}
	for {
		userInput := chat.StringPrompt("Human>")
		if userInput == "exit" {
			break
		}
		fmt.Printf(chat.SendChat(userInput) + "\n")
	}

}
