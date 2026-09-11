package main

import (
	"fmt"
	"unicode"
)

// Создайте функцию capitalizeWords(s string) string,
// которая преобразует каждое слово в строке так, чтобы первая буква была заглавной,
// а остальные — строчными. Например: "привет мир" → "Привет Мир".

func main() {
	text := "hello world ! h h  hhhhh"
	fmt.Println(capitalizeWords(text))
}

func capitalizeWords(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		if runes[i-1] == ' ' && runes[i] != ' ' {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}
	return string(runes)
}
