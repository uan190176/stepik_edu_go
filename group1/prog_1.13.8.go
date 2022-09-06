/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.
Выходные данные
Выведите количество минимальных элементов последовательности.

Sample Input:
3
21
11
4
Sample Output:
1
*/
package main

import "fmt"

func main() {
	var n, c, m int
	fmt.Scan(&n)
	if n > 0 {
		a := make([]int, n, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
			if i == 0 {
				m = a[i]
			}
			if a[i] < m {
				m = a[i]
			}
		}
		for i := 0; i < n; i++ {
			if a[i] == m {
				c++
			}
		}
		fmt.Println(c)
	}
}

/*
func main() {
	var n, count, min int
	fmt.Scan(&n)
	for i := n; i > 0; i-- {
		var a int
		fmt.Scan(&a)
		if i == n || a < min {
			min, count = a, 1
		} else if a == min {
			count++
		}
	}
	fmt.Println(count)
}
*/
