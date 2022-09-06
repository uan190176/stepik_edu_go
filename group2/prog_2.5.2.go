/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:
топот
Sample Output:
Палиндром
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.ToLower(strings.TrimSuffix(text, "\n"))

	rev := Reverse(text)
	if text == rev {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
