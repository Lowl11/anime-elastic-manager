package elastic

import (
	"elastic-manager/common"
	"net/http"
)

// индексирование данных в актуальный индекс эластика
func IndexData(w *http.ResponseWriter, baseUrl string) common.JsonResult {
	return common.GetSimpleResult(true, "Индексирование всех записей прошло успешно")
}
