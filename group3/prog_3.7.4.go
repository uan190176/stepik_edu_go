/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).
Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:
12 мин. 13 сек.
Sample Output:
Fri May 15 19:28:18 UTC 2020
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const now = 1589570165

func stoi(s string) (min, sec int) {
	ss := strings.Split(s, ".")
	//fmt.Println(ss)
	s1, s2 := ss[0], ss[1]
	//fmt.Println(s1, s2)
	s1 = strings.Replace(s1, " мин", "", 1)
	s2 = strings.Replace(s2, " сек", "", 1)
	//fmt.Println(s1, s2)
	s1 = strings.Trim(s1, " ")
	s2 = strings.Trim(s2, " ")
	//fmt.Println(s1, s2)
	min, _ = strconv.Atoi(s1)
	sec, _ = strconv.Atoi(s2)
	return
}
func main() {
	var s string
	t := time.Unix(now, 0).UTC()
	//fmt.Println(t)

	rd := bufio.NewScanner(os.Stdin)
	rd.Scan()
	s = rd.Text()
	min, sec := stoi(s)
	//fmt.Println(min, sec)

	t = t.Add(time.Minute*time.Duration(min) + time.Second*time.Duration(sec))
	fmt.Println(t.Format(time.UnixDate))
}
