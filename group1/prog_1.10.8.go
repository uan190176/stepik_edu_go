/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
Выходные данные
Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.

Sample Input:
564 8954
Sample Output:
5 4
*/
package main

import "fmt"

func main() {
	var x int
	var y int
	var k int
	fmt.Scan(&x)
	fmt.Scan(&y)
	for j := 1000; x/j >= 0; j = (j / 10) {
		if x/j == 0 {
			continue
		}
		k = (x / j) % 10
		for i := 1000; y/i >= 0; i = (i / 10) {
			if y/i == 0 {
				continue
			}
			if k == (y/i)%10 {
				fmt.Print(k, " ")
			}
			if i == 1 {
				break
			}
		}
		if j == 1 {
			break
		}
	}
}
