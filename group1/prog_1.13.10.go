/*
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).
Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.

Sample Input:
100
500
Sample Output:
497
*/
package main

import (
	"fmt"
)

func main() {
	var a, b, i int

	fmt.Scan(&a)
	fmt.Scan(&b)

	for i = b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}
	}

	fmt.Println("NO")
}
