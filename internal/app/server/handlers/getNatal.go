package web_app

import (
	"astrologist/internal/app/models"
	"astrologist/internal/app/store/sqlstore"
	pageAssembler "astrologist/internal/app/templates/page_assembler"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewNatalHandler(store sqlstore.StoreInterface, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const path = "handlers.getNatal.NewNatalHandler"
		var input models.NatalCardInput
		params := r.URL.Query()
		rawQuery := r.URL.RawQuery
		input, err := natalInputParse(params, log)
		if err != nil {
			log.Errorf("%s : Ошибка получения данных: %v", path, err.Error())
			RespondAPI(w, r, http.StatusBadRequest, "Ошибка получения данных")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(pageAssembler.GetNatalResult()))
		chart, err := store.NatalChart().GetChart(rawQuery, input)
		if err != nil {
			log.Errorf("%s : Ошибка получения натальной карты: %v", path, err.Error())
			RespondAPI(w, r, http.StatusBadRequest, "Ошибка получения натальной карты")
			return
		}
		fmt.Println(chart)
		return
	}
}
