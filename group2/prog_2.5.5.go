/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:
zaabcbd
Sample Output:
zcd
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
	s = strings.ToLower(strings.TrimSuffix(s, "\n"))

	tmp := ""
	for i := 0; i < len(s); i++ {
		if strings.Count(s, s[i:i+1]) > 1 {
			if strings.Count(tmp, s[i:i+1]) == 0 {
				tmp += s[i : i+1]
			}
		}
	}
	for i := 0; i < len(tmp); i++ {
		s = strings.ReplaceAll(s, tmp[i:i+1], "")
	}
	fmt.Println(s)
}
