// mocky is a tiny program to convert an input text into a MoCkInG tExT.
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input := buildInputFromOsArgs()
	fmt.Println(mock(input))
}

func buildInputFromOsArgs() string {
	if len(os.Args) < 2 {
		fmt.Println("error: no input given")
		os.Exit(1)
	}

	var builder strings.Builder
	for _, v := range os.Args[1:] {
		builder.WriteString(v + " ")
	}
	str := builder.String()
	return str[:len(str)-1]
}

func mock(input string) string {
	var builder strings.Builder
	uppercase := true

	for _, r := range input {
		targetRune := r
		if unicode.IsLetter(r) {
			switch uppercase {
			case true:
				targetRune = unicode.ToUpper(r)
			case false:
				targetRune = unicode.ToLower(r)
			}
			uppercase = !uppercase
		}

		builder.WriteRune(targetRune)
	}
	return builder.String()
}
