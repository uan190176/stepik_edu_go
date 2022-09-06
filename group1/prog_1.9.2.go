/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.
Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:
237
Sample Output 1:
YES
Sample Input 2:
117
Sample Output 2:
NO
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, f, s, t int

	fmt.Scan(&a)

	if len(strconv.Itoa(a)) < 3 {
		t = a % 10
		s = (a % 100) / 10
		f = a / 100
		if f == s || f == t || s == t {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
