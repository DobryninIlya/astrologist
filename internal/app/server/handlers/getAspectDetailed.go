package web_app

import (
	"astrologist/internal/app/parser"
	"astrologist/internal/app/store/sqlstore"
	pageAssembler "astrologist/internal/app/templates/page_assembler"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewAspectDetailedHandler(store sqlstore.StoreInterface, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const path = "handlers.getAspectDetailed.NewAspectDetailedHandler"
		//var input models.NatalCardInput
		//params := r.URL.Query()
		//rawQuery := r.URL.RawQuery
		//input, err := natalInputParse(params, log)
		//if err != nil {
		//	log.Errorf("%s : Ошибка получения данных: %v", path, err.Error())
		//	RespondAPI(w, r, http.StatusBadRequest, "Ошибка получения данных")
		//	return
		//}
		//chart, err := store.NatalChart().GetChart(rawQuery, input)
		//if err != nil {
		//	log.Errorf("%s : Ошибка получения натальной карты: %v", path, err.Error())
		//	RespondAPI(w, r, http.StatusBadRequest, "Ошибка получения натальной карты")
		//	return
		//}
		name := chi.URLParam(r, "name")
		aspectPage, err := parser.GetAspectDetailedPage(name)
		if err != nil {
			log.Errorf("%s : Ошибка получения данных: %v", path, err.Error())
			RespondAPI(w, r, http.StatusBadRequest, "Ошибка получения данных")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(pageAssembler.GetAspectDetailed(aspectPage)))
		return
	}
}
