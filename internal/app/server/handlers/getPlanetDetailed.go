package web_app

import (
	"astrologist/internal/app/models"
	"astrologist/internal/app/store/sqlstore"
	pageAssembler "astrologist/internal/app/templates/page_assembler"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewPlanetDetailedHandler(store sqlstore.StoreInterface, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
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
		chart, err := store.NatalChart().GetChart(rawQuery, input)
		if err != nil {
			log.Errorf("%s : Ошибка получения натальной карты: %v", path, err.Error())
			RespondAPI(w, r, http.StatusBadRequest, "Ошибка получения натальной карты")
			return
		}
		planetCase := chi.URLParam(r, "planet")
		planet := getPlanetData(planetCase, chart)
		//fmt.Println(planet)
		//fmt.Println(chart)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(pageAssembler.GetPlanetDetailed(planet, chart, planetCase)))
		return
	}
}

func getPlanetData(planet string, chart models.NatalCardOutput) models.AstroData {
	var data models.AstroData
	switch planet {
	case "sun":
		data = chart.Sun
	case "moon":
		data = chart.Moon
	case "mercury":
		data = chart.Mercury
	case "venus":
		data = chart.Venus
	case "mars":
		data = chart.Mars
	case "jupiter":
		data = chart.Jupiter
	case "saturn":
		data = chart.Saturn
	case "uranus":
		data = chart.Uran
	case "neptune":
		data = chart.Neptune
	case "pluto":
		data = chart.Pluto
	case "charon":
		data = chart.Hiron
	}
	return data
}
