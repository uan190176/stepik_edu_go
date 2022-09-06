/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

Входные данные
Задано единственное число N
Выходные данные
Необходимо вывести требуемое представление числа N.

Sample Input:
12
Sample Output:
1100
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, i, r int
	var res = ""

	fmt.Scan(&a)

	for i = 0; a > 0; i++ {
		r = a % 2
		a = a / 2
		res = strconv.Itoa(r) + res
	}

	fmt.Print(res)
}
