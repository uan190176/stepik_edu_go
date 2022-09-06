/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

Входные данные
Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.
Выходные данные
Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:
1112221112
Sample Output:
2
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	b := []byte(s)
	m := b[0]
	for i := 1; i < len(b); i++ {
		if b[i] > m {
			m = b[i]
		}
	}
	fmt.Println(string(m))
}
