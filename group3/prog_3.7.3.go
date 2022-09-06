/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:
24h0m0s
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s string
	//fmt.Scan(&s)
	//s = "2020-05-15 14:00:00"
	f := "02.01.2006 15:04:05"
	rd := bufio.NewScanner(os.Stdin)
	rd.Scan()
	s = rd.Text()
	ss := strings.Split(s, ",")
	s1, s2 := ss[0], ss[1]

	t1, e := time.Parse(f, s1)
	if e != nil {
		panic(e)
	}
	t2, e := time.Parse(f, s2)
	if e != nil {
		panic(e)
	}

	d := time.Since(t1).Round(time.Second) - time.Since(t2).Round(time.Second)
	if d < 0 {
		d = -1 * d
	}

	fmt.Println(d)
}
