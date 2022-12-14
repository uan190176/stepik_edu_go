/*
Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
которую внутри функции main в дальнейшем можно будет вызвать по имени fn.
Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют
нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100.
Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

Sample Input:
727178
Sample Output:
28
*/
package main

// fn
/*
fn := func (x uint) uint {
	var rs string
	var r uint
	s := strconv.FormatUint(uint64(x), 10)
	for i := 0; i < len(s); i++ {
		a, _ := strconv.Atoi(s[i : i+1])
		if a != 0 && a%2 == 0 {
			rs += s[i : i+1]
		}
	}
	t, _ := strconv.Atoi(rs)
	r = uint(t)
	if r == 0 {
		r = 100
	}
	return r
}
*/
