/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.

Sample Input:
50
Sample Output:
1 2 4 8 16 32
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, i, pow float64

	fmt.Scan(&a)

	for i = 0; pow < a; i++ {
		pow = math.Pow(2, i)
		if pow > a {
			break
		}
		fmt.Print(pow)
		fmt.Print(" ")
	}
}
