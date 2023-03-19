package main

import (
	"fmt"
	"os"
	"encoding/json"

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

func jsonEscape(i string) string {
    b, err := json.Marshal(i)
    if err != nil {
        panic(err)
    }
    // Trim the beginning and trailing " character
    return string(b[1:len(b)-1])
}

func main() {
	token := os.Getenv("CHATGPT_TOKEN")
	conversation := ""
	response := ""
	if token == "" {
		fmt.Printf(tokenNotSet)
		os.Exit(1)
	}
	for {
		userInput := chat.StringPrompt("Human>")
		if userInput == "exit" {
			break
		}

		conversation = conversation + jsonEscape(userInput)
		response = chat.SendChat(conversation, token)
		fmt.Printf(response + "\n")
		conversation = conversation + jsonEscape(response)
	}

}
