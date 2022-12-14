/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Входные данные
На вход программе подается целое число d (0 < d < 360).
Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.

Sample Input:
90
Sample Output:
It is 3 hours 0 minutes.
************************************************
Решение:
12 часов = 360 градусов, 1 час = 60 минут, тогда 360 градусов = 12 * 60 = 720 минут
отсюда 1 градус = 2 минуты
30 градусов - это один час, поэтому количество целых часов будет = name / 30,
а минуты получаются из формулы, по которой подсчитывается остаток: name % 30.
И этот результат надо умножить на 2, так как один градус - это 2 минтуtы.
*/
package main

import "fmt"

func main() {
	var d int

	fmt.Scan(&d)
	if d > 0 && d < 360 {
		fmt.Println("It is", d/30, "hours", 2*(d%30), "minutes.")
	}
}
