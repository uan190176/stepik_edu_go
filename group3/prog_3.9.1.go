//3.9.1
/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.
Функция work() ничего не принимает и не возвращает.
*/
package main

/*Объяснение:

создаём канал пустых структур;
создаём анонимную функцию и сразу же запускаем её в виде горутины;
внутри горутины вызываем work() и закрываем канал;
основной поток всё это время после запуска горутины ждёт закрытия канала. work() может работать сколь угодно долго, и основной поток в любом случае дождётся её завершения.
Верное решение #325424031

done := make(chan struct{})
go func() {
	work()
	close(done)
}()
<-done */

//3.9.2
/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.


	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work()
		}(wg)
	}
	wg.Wait()

*/

//3.9.3
/*
Вам необходимо написать функцию calculator следующего вида:

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.

в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.

После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите предельное время работы.

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
     ch := make(chan int)

    go func (){
            select {
                case num := <- firstChan:
                ch <- num*num
                close(ch)
                return

                case num := <- secondChan:
                ch <- num*3
                close(ch)
                return

                case <-stopChan:
                close(ch)
                return
            }
    } ()

    return ch
}

*/

//3.9.4
/*
Вам необходимо написать функцию calculator следующего вида:

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.

Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу. Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая управление.

Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
    ch := make(chan int)
    sum := 0

    go func (){
        for {
            select {
                case num := <- arguments:
                sum += num

                case <-done:
                ch <- sum
                close(ch)
                return
            }
        }
    } ()

    return ch
}

*/

//3.9.5
/*
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы:

n раз сделать следующее

прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
вычислить f(x1) + f(x2)
записать полученное значение в out
Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.



Формат ввода:

количество итераций передается через аргумент n.
целые числа подаются через аргументы-каналы in1 и in2.
функция для обработки чисел перед сложением передается через аргумент fn.
Формат вывода:

канал для вывода результатов передается через аргумент out.

package main

import (
	"fmt"  // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"time" // пакет используется для проверки выполнения условия задачи, не удаляйте его

)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int) {
		done := make(chan bool, n*2)
		results := make([]*int, n*2)

		go func() {
			nDone := 0
			for <-done {
				for i := nDone; i < n; i++ {
					if results[i] != nil && results[i+n] != nil {
						out <- (*results[i] + *results[i+n])
						if nDone++; nDone == n {
							return
						}
					} else {
						break
					}
				}
			}
		}()
		input := func(ch <-chan int, results []*int) {
			for i := 0; i < n; i++ {
				x := <-ch
				go func(i int, x int) {
					result := f(x)
					results[i] = &result
					done <- true
				}(i, x)
			}
		}
		go input(in1, results[:n])
		go input(in2, results[n:])

	}(fn, in1, in2, out)
}
*/
