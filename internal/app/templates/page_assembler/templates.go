package page_assembler

import (
	"astrologist/internal/app/models"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

func GetNatalResult(input models.NatalCardInput, output models.NatalCardOutput) template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_result.html"))
	return template.HTML(fmt.Sprintf(data,
		input.FirstName,
		fmt.Sprintf("%02d-%02d-%d", input.BrithDay, input.BrithMonth, input.BrithYear),
		fmt.Sprintf("%02d:%02d", input.BrithHour, input.BrithMinute),
		input.City,
		strings.Replace(output.MainImage, "\"", "&quot;", -1),
		//url.QueryEscape(output.MainImage),
	))
}

func GetNatalPlanetsCoordinatesResult(chart models.NatalCardOutput) template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_planet_coordinates.html"))
	table := template.HTML(chart.PlanetTables[0])
	fmt.Println(chart.PlanetTables[1])
	return template.HTML(fmt.Sprintf(data, table))
}

func GetNatalPlanetsAspectsResult(chart models.NatalCardOutput) template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_planet_aspects.html"))
	return template.HTML(fmt.Sprintf(data, chart.Planet.Sun.Power, chart.Planet.Sun.Harmony,
		chart.Planet.Moon.Power, chart.Planet.Moon.Harmony,
		chart.Planet.Mercury.Power, chart.Planet.Mercury.Harmony,
		chart.Planet.Venus.Power, chart.Planet.Venus.Harmony,
		chart.Planet.Mars.Power, chart.Planet.Mars.Harmony,
		chart.Planet.Jupiter.Power, chart.Planet.Jupiter.Harmony,
		chart.Planet.Saturn.Power, chart.Planet.Saturn.Harmony,
		chart.Planet.Uran.Power, chart.Planet.Uran.Harmony,
		chart.Planet.Neptune.Power, chart.Planet.Neptune.Harmony,
		chart.Planet.Pluto.Power, chart.Planet.Pluto.Harmony,
	))
}

func GetNatalKarmaResult() string {
	data := readFile(filepath.Join(htmlPath, "natal_karma.html"))
	return data
}

func GetPlanetDetailed(planet models.AstroData, chart models.NatalCardOutput, planetCase string) string {
	data := readFile(filepath.Join(htmlPath, "natal_planet_detailed.html"))
	var aspectList []string
	for _, val := range planet.AspectList {
		aspectList = append(aspectList, getAspectListElement(val.Link, val.AspectName))
	}
	aspectBlock := strings.Join(aspectList, "\n")

	planetAspectDescription := getPlanetAspectDescription(planet.Position, planet.Degree, planet.PositionMeaning)
	planetCaseTranslated := translatePlanetCase(planetCase)
	data = fmt.Sprintf(data,
		planetCaseTranslated, planetCase, planetCase, planetCaseTranslated,
		planet.Power, planet.Harmony,
		planet.Description,
		planetAspectDescription,
		planet.AspectDesc,
		aspectBlock,
	)
	return data
}

func GetPlanetHarmony(chart models.NatalCardOutput) template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_planet_harmony.html"))
	table := template.HTML(chart.PlanetTables[1])
	return template.HTML(fmt.Sprintf(data, table))
}

func GetPlanetOrbit(chart models.NatalCardOutput) template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_planet_orbits.html"))
	table := template.HTML(chart.PlanetTables[2])
	return template.HTML(fmt.Sprintf(data, table))
}

func GetDescriptionPlanet(chart models.NatalCardOutput, planetCase string) template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_description_planet.html"))
	if planetCase == "lilit" {
		return template.HTML(fmt.Sprintf(data, "Лилит", "darkmoon", chart.LilitData.Description, chart.LilitData.Header, chart.LilitData.DetailedDescription))
	} else if planetCase == "selena" {
		return template.HTML(fmt.Sprintf(data, "Селена", "whitemoon", chart.SelenaData.Description, chart.SelenaData.Header, chart.SelenaData.DetailedDescription))
	} else if planetCase == "moon_nodes" {
		return template.HTML(fmt.Sprintf(data, "Лунные узлы", "moonnodes", "Лунные узлы "+chart.MoonNodes.Description, chart.MoonNodes.Header, chart.MoonNodes.DetailedDescription))
	}
	return template.HTML(fmt.Sprintf(data, ""))
}

func GetNatalForm() template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_form.html"))
	return template.HTML(data)
}

func GetAspectDetailed(name models.AspectDetailedPage) template.HTML {
	data := readFile(filepath.Join(htmlPath, "natal_aspect_detailed.html"))
	var resultList strings.Builder
	for _, val := range name.AspectParagraphs {
		resultList.WriteString(getAspectParagraph(val.Header, val.Description))
	}

	return template.HTML(fmt.Sprintf(data, strings.Split(name.Header, ".")[0], name.Header,
		name.Description,
		resultList.String()))

}
