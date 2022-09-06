/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.
Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:
1234
Sample Output:
1
*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
Здесь 6 вариантов решения, на мой взгляд логичнее запостить решения с if или switch, т.к. они соответствуют теме параграфа
*/
func main() {
	var a int

	fmt.Scan(&a)

	if a < 10000 && a > 0 {
		// 1. Use strconv
		s := strconv.Itoa(a)
		fmt.Println(s[0:1])

		// 2. Use recursive func
		fmt.Println(fn(a))

		// 3. Use switch
		switch {
		case a < 10:
			fmt.Println(a)
		case a < 100:
			fmt.Println(a / 10)
		case a < 1000:
			fmt.Println(a / 100)
		case a < 10000:
			fmt.Println(a / 1000)
		case a < 100000:
			fmt.Println(a / 10000)
		}

		// 4. Use if
		if a <= 9 {
			fmt.Println(a)
		} else if a >= 10 && a <= 99 {
			fmt.Println(a / 10)
		} else if a >= 100 && a <= 999 {
			fmt.Println(a / 100)
		} else if a >= 1000 && a <= 9999 {
			fmt.Println(a / 1000)
		} else {
			fmt.Println(a / 10000)
		}

		// 5. Use math
		fmt.Print(int(float64(a) / math.Pow10(int(math.Log10(float64(a))))))

		// 6. Use for
		for a > 9 {
			a /= 10
		}
		fmt.Println(a)
	}
}

func fn(a int) int {
	if a < 10 {
		return a
	} else {
		return fn(a / 10)
	}
}
