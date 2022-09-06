/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A.
Если А не является числом Фибоначчи, выведите число -1.

Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.

Sample Input:
8
Sample Output:
6
*/
package main

import (
	"fmt"
)

func main() {
	var a, i, val int
	var fib []int

	fmt.Scan(&a)

	val = 0

	// init fib numbers
	fib = append(fib, 0)
	fib = append(fib, 1)

	if a == 0 {
		fmt.Print(0)
		return
	}

	if a == 1 {
		fmt.Print(1)
		return
	}

	for i = 2; val < a; i++ {
		val = fib[i-2] + fib[i-1]
		fib = append(fib, val)
		//fmt.Println(fib[i])
		if fib[i] == a {
			fmt.Print(i)
			return
		}
	}

	fmt.Print(-1)
}
