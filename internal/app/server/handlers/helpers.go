package web_app

import (
	"astrologist/internal/app/models"
	"github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

func natalInputParse(params url.Values, log *logrus.Logger) (models.NatalCardInput, error) {
	const path = "handlers.helpers.natalInputParse"
	var input models.NatalCardInput
	input.FirstName = params.Get("fn")
	birthDay, err := strconv.Atoi(params.Get("fd"))
	input.BrithDay = birthDay
	if err != nil {
		log.Errorf("%s : Ошибка получения даты рождения: %v", path, err.Error())
		return models.NatalCardInput{}, err
	}
	birthMonth, err := strconv.Atoi(params.Get("fm"))
	input.BrithMonth = birthMonth
	if err != nil {
		log.Errorf("%s : Ошибка получения месяца рождения: %v", path, err.Error())
		return models.NatalCardInput{}, err
	}
	birthYear, err := strconv.Atoi(params.Get("fy"))
	input.BrithYear = birthYear
	if err != nil {
		log.Errorf("%s : Ошибка получения года рождения: %v", path, err.Error())
		return models.NatalCardInput{}, err
	}
	birthHour, err := strconv.Atoi(params.Get("fh"))
	input.BrithHour = birthHour
	if err != nil {
		log.Errorf("%s : Ошибка получения часа рождения: %v", path, err.Error())
		return models.NatalCardInput{}, err
	}
	birthMinute, err := strconv.Atoi(params.Get("fmn"))
	input.BrithMinute = birthMinute
	if err != nil {
		log.Errorf("%s : Ошибка получения минут рождения: %v", path, err.Error())
		return models.NatalCardInput{}, err
	}
	timeZoneID, err := strconv.Atoi(params.Get("ttz"))
	if err != nil {
		log.Errorf("%s : Ошибка получения часового пояса: %v", path, err.Error())
		return models.NatalCardInput{}, err
	}
	input.TimeZoneID = timeZoneID
	input.Latitude = params.Get("lt")
	input.Longitude = params.Get("ln")
	if input.FirstName == "" || input.Latitude == "" || input.Longitude == "" {
		log.Errorf("%s : Ошибка получения данных: %v", path, "Не все данные были получены")
		return models.NatalCardInput{}, err
	}
	input.City = params.Get("c1")
	return input, nil
}
