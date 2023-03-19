package chat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/chzyer/readline"
)

type ChatGPTChoices struct {
	Index        int    `json:"index"`
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
	Logprobs     string `json:"logprobs"`
}

type ChatGPTUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Answer struct {
	Id      string           `json:"id"`
	Object  string           `json:"object"`
	Created int              `json:"created"`
	Model   string           `json:"model"`
	Choices []ChatGPTChoices `json:"choices"`
	Usage   ChatGPTUsage     `json:"usage"`
}

func SendChat(query string, token string) string {
	var answer Answer
	var url = "https://api.openai.com/v1/completions"
	var jsonBody = []byte(`{
		"model": "text-davinci-003",
		"prompt": "This is a conversation with a human being. Provide the next message that would be said as a reply to the last line of text.\n` + query + `\n",
		"temperature": 0.7,
		"max_tokens": 2048
	  }`)

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(body), &answer)
	// return string(body)
	return answer.Choices[0].Text

}

func StringPrompt(label string) string {
	var s string
	rl, err := readline.New(label + " ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	s, _ = rl.Readline()
		
	return strings.TrimSpace(s)
}
