package main

import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) string {
	var words []string
	var spaceNum int
	for i, j := 0, 0; j < len(text); j++ {
		if text[j] == ' ' {
			if text[i] != ' ' {
				words = append(words, text[i:j])
				i = j
			}
			spaceNum++
		} else {
			if text[i] == ' ' {
				i = j
			}
			if j == len(text)-1 {
				words = append(words, text[i:])
			}
		}
	}

	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaceNum)
	} else {
		return strings.Join(words, strings.Repeat(" ", spaceNum/(len(words)-1))) + strings.Repeat(" ", spaceNum%(len(words)-1))
	}
}

func main() {
	fmt.Println(reorderSpaces("  walks  udp package   into  bar a"))
}
