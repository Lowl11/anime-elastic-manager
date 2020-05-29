package elastic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lazy-owl/elastic-manager/src/common"
	"net/http"
)

// индексирование данных в актуальный индекс эластика
func IndexData(w *http.ResponseWriter, baseUrl string, request *http.Request) common.JsonResult {
	receiveData, _ := ioutil.ReadAll(request.Body) // считываем в байтах

	// json структура для гоу выглядит именно так
	// ключ всегда строка а значение может быть чем угодно
	var parsedJson map[string]interface{}

	// ну и парсим соответственно
	json.Unmarshal([]byte(receiveData), &parsedJson)

	animeList := parsedJson["anime_list"]
	// TODO следующий шаг не понятен: мы получили мапу в которой лежит все что нам нужно но фиг пойми как вытащить глубже
	fmt.Println("Parsed aninme list: ", animeList)
	return common.GetSimpleResult(true, "Индексирование всех записей прошло успешно")
}
