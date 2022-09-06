/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:

	{
	    "ID":134,
	    "Number":"ИЛМ-1274",
	    "Year":2,
	    "Students":[
	        {
	            "LastName":"Вещий",
	            "FirstName":"Лифон",
	            "MiddleName":"Вениаминович",
	            "Birthday":"4апреля1970года",
	            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
	            "Phone":"+7(948)709-47-24",
	            "Rating":[1,2,3]
	        },
	        {
	            // ...
	        }
	    ]
	}

В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:

	{
	    "Average": 14.1
	}
*/
package main

import (
	"encoding/json"
	"os"
)

type Group struct {
	Students []struct {
		Rating []int
	}
}

func main() {
	var group Group
	if err := json.NewDecoder(os.Stdin).Decode(&group); err != nil {
		panic(err)
	}

	var avg float64
	for _, st := range group.Students {
		avg += float64(len(st.Rating))
	}
	avg /= float64(len(group.Students))

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "    ")
	e.Encode(struct{ Average float64 }{avg})
}
