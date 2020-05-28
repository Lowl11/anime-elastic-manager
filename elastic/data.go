package elastic

import (
	"elastic-manager/common"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// индексирование данных в актуальный индекс эластика
func IndexData(w *http.ResponseWriter, baseUrl string, request *http.Request) common.JsonResult {
	// TODO: не принимает параметры с POST запроса
	params := mux.Vars(request)
	fmt.Println("Params:", params)
	return common.GetSimpleResult(true, "Индексирование всех записей прошло успешно")
}
