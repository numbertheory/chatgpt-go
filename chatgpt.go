package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"strconv"

	"github.com/numbertheory/chatgpt-go/chat"
	"github.com/TwiN/go-color"
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
	return string(b[1 : len(b)-1])
}

func main() {
	token := os.Getenv("CHATGPT_TOKEN")
	conversation := ""
	response := ""
	turn := 0
	var record_interactions []string
	if token == "" {
		fmt.Printf(tokenNotSet)
		os.Exit(1)
	}
	for {
		userInput := chat.StringPrompt(color.InGreen("(" + strconv.Itoa(turn) +")Human>"))
		if userInput == "exit" {
			break
		}
		record_interactions = append(record_interactions, userInput)
		conversation = conversation + jsonEscape(userInput)
		response = chat.SendChat(conversation, token)
		fmt.Printf(color.InPurple("ChatGPT> ") + strings.TrimSpace(response) + "\n")
		conversation = conversation + jsonEscape(response)
		record_interactions = append(record_interactions, response)
		turn += 1
		// every 20th turn we will keep the first ten messages and the last ten messages to prevent the conversation from 
		// overwhelming the API
		if turn == 10 {
			conversation = ""
			conversation = conversation + jsonEscape(strings.Join(record_interactions[0:2], ""))
			conversation = conversation + jsonEscape(strings.Join(record_interactions[8:10], ""))
			record_interactions = []string{}
			turn = 0
		}
	}

}
