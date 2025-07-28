package utils

import "fmt"

func PrintReversedMessage(msg string) {
	msgRune := []rune(msg)
	for i, j := 0, len(msgRune)-1; i < j; i, j = i+1, j+1 {
		msgRune[i], msgRune[j] = msgRune[j], msgRune[i]
	}

	fmt.Println(string(msgRune))
}
