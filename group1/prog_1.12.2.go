/*
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:
5
41 -231 24 49 6
Sample Output:
49
*/
package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	a := make([]int, n, n+1)
	if n >= 4 {
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		fmt.Println(a[3])
	}
}
