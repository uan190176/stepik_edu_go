/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений:
"n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).
Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

Sample Input:
10
Sample Output:
10 korov
*/
package main

import (
	"fmt"
)

func main() {
	var n int
	var cow int

	fmt.Scan(&n)
	if n >= 11 && n <= 14 {
		fmt.Printf("%v korov", n)
	} else {
		cow = n % 10
		if cow == 0 || (cow >= 5 && cow <= 9) {
			fmt.Printf("%v korov", n)
		} else if cow == 1 {
			fmt.Printf("%v korova", n)
		} else if cow >= 2 && cow <= 4 {
			fmt.Printf("%v korovy", n)
		}
	}
}
