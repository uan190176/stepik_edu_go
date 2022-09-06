/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot
Sample Output:
hello
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.ToLower(strings.TrimSuffix(text, "\n"))

	for i := 0; i < len(text); i++ {
		if i%2 != 0 {
			fmt.Print(text[i : i+1])
		}
	}
}
