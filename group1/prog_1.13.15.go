/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.
Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3
Sample Output:
801272
*/
package main

import (
	"fmt"
)

func main() {
	var a, n, res string
	var i int

	fmt.Scan(&a)
	fmt.Scan(&n)

	for i = 0; i < len(a); i++ {
		if string(a[i]) != n {
			res += string(a[i])
		}
	}

	fmt.Print(res)
}

/*
Без строки
package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    result := 0
    for i := 1; a > 0; {
        if a % 10 != b {
            result += a % 10 * i
            i *= 10
        }
        a /= 10
    }
    fmt.Println(result)
}

Без строки с дефером
деферы выполняются в обратном порядке

package main

import "fmt"

func main() {
	var num, n int
	fmt.Scan(&num, &n)
	for ; num > 0; num /= 10 {
		if num % 10 != n {
			defer fmt.Print(num % 10)
		}
	}
}
*/
