/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1
Sample Output:
Ok
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.ToLower(strings.TrimSuffix(s, "\n"))
	p := `^[a-zA-Z0-9]{5,}$`
	t, _ := regexp.MatchString(p, s)

	if len(s) > 4 && t {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
