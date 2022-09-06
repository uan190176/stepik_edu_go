/*
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
Если подстроки S нет в строке X - вывести -1

Sample Input:
awesome
es
Sample Output:
2
*/
package main

import (
	"fmt"

	"strings"
)

func main() {

	x, s := "", ""

	fmt.Scan(&x)

	fmt.Scan(&s)

	res := strings.Index(x, s)

	fmt.Println(res)

}
