package web_app

import (
	"astrologist/internal/app/store/sqlstore"
	pageAssembler "astrologist/internal/app/templates/page_assembler"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewNatalFormHandler(store sqlstore.StoreInterface, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const path = "handlers.getNatal.NewNatalHandler"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(pageAssembler.GetNatalForm()))

		return
	}
}
