package sqlstore

import (
	"astrologist/internal/app/models"
	"astrologist/internal/app/parser"
	"encoding/json"
)

// NatalChartRepository ...
type NatalChartRepository struct {
	store *Store
}

func (r NatalChartRepository) create(key string, input models.NatalCardInput) (models.NatalCardOutput, error) {
	result, err := parser.GetNatalChart(input)
	if err != nil {
		return models.NatalCardOutput{}, err
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return models.NatalCardOutput{}, err
	}
	return result, r.store.db.QueryRow(
		"INSERT INTO natal_chart_result (key, json_result, date_created) VALUES ($1, $2, NOW()) RETURNING key",
		key, jsonResult,
	).Scan(&key)
}

func (r NatalChartRepository) find(key string) (*models.NatalCardOutput, error) {
	var jsonResult string
	err := r.store.db.QueryRow(
		"SELECT json_result FROM natal_chart_result WHERE key = $1",
		key,
	).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}
	var result models.NatalCardOutput
	err = json.Unmarshal([]byte(jsonResult), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r NatalChartRepository) GetChart(key string, input models.NatalCardInput) (models.NatalCardOutput, error) {
	foundResult, err := r.find(key)
	if err == nil {
		return *foundResult, nil
	}
	result, err := r.create(key, input)
	if err != nil {
		return models.NatalCardOutput{}, err
	}
	return result, nil
}
