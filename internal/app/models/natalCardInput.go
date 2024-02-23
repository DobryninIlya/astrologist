package models

type NatalCardInput struct {
	FirstName   string `json:"first_name"`
	BrithDay    int    `json:"brith_day"`
	BrithMonth  int    `json:"brith_month"`
	BrithYear   int    `json:"brith_year"`
	BrithHour   int    `json:"brith_hour"`
	BrithMinute int    `json:"brith_minute"`
	TimeZoneID  int    `json:"time_zone_id"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	City        string `json:"city"`
}
