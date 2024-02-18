package web_app

import (
	"astrologist/internal/app/store/sqlstore"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func NewAutocompleteHandler(store sqlstore.StoreInterface, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const path = "handlers.getAutocomplete.NewAutocompleteHandler"
		url := "https://geocult.ru/swetest/app/form/fetch/fetch_autocomplete.php"
		method := "POST"
		termParam := r.URL.Query().Get("term")
		payload := strings.NewReader(fmt.Sprintf("term=%s", termParam))

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}
}
