/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:
5
38 24 800 8 16

Sample Output:
40
*/
package main

import "fmt"

func main() {
	var n, a, s int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > 9 && a < 100 && a%8 == 0 {
			s += a
		}
	}
	fmt.Println(s)
}