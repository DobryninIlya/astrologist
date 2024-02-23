package models

type NatalCardOutput struct {
	MainImage    string   `json:"main_image"`
	BottomImage  string   `json:"bottom_image"`
	PlanetTables []string `json:"planet_tables"`
	Planet
	Karma      KarmaData
	MoonNodes  MoonNodes
	SelenaData SelenaData
	LilitData  LilitData
}

type AstroData struct {
	Title           string
	Power           float64
	Harmony         float64
	Description     string
	Position        string
	Degree          string
	PositionMeaning string
	AspectName      string
	AspectDesc      string
	AspectList      []AspectList
}

type AspectList struct {
	AspectName string
	Link       string
}

type Planet struct {
	Sun     AstroData
	Moon    AstroData
	Mercury AstroData
	Venus   AstroData
	Mars    AstroData
	Jupiter AstroData
	Saturn  AstroData
	Uran    AstroData
	Neptune AstroData
	Pluto   AstroData
	Hiron   AstroData
}

type KarmaData struct {
	Description         string
	Header              string
	DetailedDescription string
}

type MoonNodes struct {
	Description         string
	Header              string
	DetailedDescription string
	AspectDescription   string
	AspectList          []AspectList
}

type SelenaData struct {
	Description         string
	Header              string
	DetailedDescription string
}

type LilitData struct {
	Description         string
	Header              string
	DetailedDescription string
}

type AspectDetailedPage struct {
	Header           string
	Description      string
	AspectParagraphs []AspectDetailedPageElement
}

type AspectDetailedPageElement struct {
	Header      string
	Description string
}
