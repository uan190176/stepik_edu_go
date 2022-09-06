/*
На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:
6 8
Sample Output:
10
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(math.Sqrt(float64(a*a + b*b)))
}
