package chat

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SendChat(query string) string {
	return "ok"
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
