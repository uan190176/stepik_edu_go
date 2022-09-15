/*
Данная задача ориентирована на реальную работу с данными в формате JSON. Для решения мы будем использовать справочник
ОКВЭД (Общероссийский классификатор видов экономической деятельности), который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json, а сами данные,
которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории
на github.com.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа записать
сумму полей global_id всех элементов, закодированных в файле.
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var items []struct {
		ID uint64 `json:"global_id"`
	}
	f, _ := os.Open("data-20190514T0100.json")
	json.NewDecoder(f).Decode(&items)
	var sum uint64
	for _, item := range items {
		sum += item.ID
	}
	fmt.Println(sum)
}