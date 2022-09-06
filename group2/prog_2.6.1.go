/*
Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке.
В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции. Функция divide(a int, b int) (int, error)
получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!

Не забудьте считать два числа которые необходимо поделить (передать в функцию)!

Sample Input:
10 5
Sample Output:
2
*/
package main

import "fmt"

func main() {
	// package main уже объявлен, нужные импорты уже есть
	var a, b int
	fmt.Scan(&a, &b)
	result, err := divide(a, b)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(result)
	}
}

// not include in answer
func divide(a, b int) (r int, e interface{}) {
	return a / b, e
}
