package parser

import (
	"astrologist/internal/app/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	mainURL        = "https://geocult.ru/"
	natalURL       = mainURL + "natalnaya-karta-onlayn-raschet"
	aspectPagesURL = mainURL + "natalnaya-karta/"
)

func GetNatalChart(input models.NatalCardInput) (models.NatalCardOutput, error) {
	urlParams := fmt.Sprintf("?fn=%v&fd=%v&fm=%v&fy=%v&fh=%v&fmn=%v&ttz=%v&lt=%v&ln=%v&hs=P&as=1&sb=1",
		input.FirstName, input.BrithDay, input.BrithMonth, input.BrithYear, input.BrithHour, input.BrithMinute, input.TimeZoneID, input.Latitude, input.Longitude)
	req, err := http.NewRequest("GET", natalURL+urlParams, nil)
	if err != nil {
		return models.NatalCardOutput{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.NatalCardOutput{}, err
	}
	defer resp.Body.Close()
	return ParseNatalDetails(resp.Body)
}

func getURLImage(imageUrl string) io.ReadCloser {
	response, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	return response.Body
}

func ParseNatalDetails(body io.ReadCloser) (models.NatalCardOutput, error) {
	var result models.NatalCardOutput
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return result, err
	}
	// Картинки с чертежами домой
	mainImage, bottomImage := getNatalChartImages(doc)
	result.MainImage = mainImage
	result.BottomImage = bottomImage

	// 3 таблицы со знаками
	tables, err := extractTables(doc)
	if err != nil {
		return result, err
	}
	result.PlanetTables = tables

	dataSUN, err := extractAstroData(doc, "sunnat")
	if err != nil {
		return result, err
	}
	dataMOON, err := extractAstroData(doc, "moonnat")
	dataMOON.Description = "\nЛуна – передатчик ощущений, эмоций, настроения. Она царит в сферах души и подсознания, врожденных рефлексов и привычек, идущих из детства. Луна — хозяйка Рака. Местом ее возвышения считают знак Тельца. Козерог считается местом ее падения, Скорпион — местом изгнания (плена). Если Луна в гороскопе сильна по статусу и имеет благоприятные взаимодействия с другими планетами натальной карты, то субъект ощущает гармонию в душе, и даже если возникают какие-то сложности в жизни, то он с легкостью преодолевает периоды неприятностей.\n"
	if err != nil {
		return result, err
	}
	dataMERCURY, err := extractAstroData(doc, "mercnat")
	if err != nil {
		return result, err
	}
	dataVENUS, err := extractAstroData(doc, "venunat")
	if err != nil {
		return result, err
	}
	dataMARS, err := extractAstroData(doc, "marsnat")
	if err != nil {
		return result, err
	}
	dataJUPITER, err := extractAstroData(doc, "jupinat")
	if err != nil {
		return result, err
	}
	dataSATURN, err := extractAstroData(doc, "satunat")
	if err != nil {
		return result, err
	}
	dataURAN, err := extractAstroData(doc, "urannat")
	if err != nil {
		return result, err
	}
	dataNEPTUNE, err := extractAstroData(doc, "neptnat")
	if err != nil {
		return result, err
	}
	dataPLUTO, err := extractAstroData(doc, "plutnat")
	if err != nil {
		return result, err
	}
	dataHIRON, err := extractAstroData(doc, "hironat")
	if err != nil {
		return result, err
	}
	result.Planet.Sun = dataSUN
	result.Planet.Moon = dataMOON
	result.Planet.Mercury = dataMERCURY
	result.Planet.Venus = dataVENUS
	result.Planet.Mars = dataMARS
	result.Planet.Jupiter = dataJUPITER
	result.Planet.Saturn = dataSATURN
	result.Planet.Uran = dataURAN
	result.Planet.Neptune = dataNEPTUNE
	result.Planet.Pluto = dataPLUTO
	result.Planet.Hiron = dataHIRON
	karma, err := extractKarmaData(doc)
	if err != nil {
		log.Printf("Ошибка получения кармы: %v", err.Error())
	}
	result.Karma = karma
	moonNodes, err := extractMoonNodes(doc)
	if err != nil {
		log.Printf("Ошибка получения Лунных узлов: %v", err.Error())
	}
	result.MoonNodes = moonNodes
	selenaData, err := extractSelenaData(doc)
	if err != nil {
		log.Printf("Ошибка получения Селены: %v", err.Error())
	}
	result.SelenaData = selenaData
	lilitData, err := extractLilitData(doc)
	if err != nil {
		log.Printf("Ошибка получения Селены: %v", err.Error())
	}
	result.LilitData = lilitData

	return result, nil
}

func getNatalChartImages(doc *goquery.Document) (string, string) {
	var mainImage, bottomImage string

	// Находим все элементы с id="r660" и id="r705"
	doc.Find("a[id=r660], a[id=r705]").Each(func(i int, s *goquery.Selection) {
		// Получаем значение атрибута href
		src, exists := s.Attr("href")
		if exists {
			// Добавляем ссылку в список
			if mainImage == "" {
				mainImage = src
			} else {
				bottomImage = src
			}
		}
	})
	return mainImage, bottomImage
}

func GetAspectDetailedPage(name string) (models.AspectDetailedPage, error) {
	resp, err := http.Get(aspectPagesURL + name)
	if err != nil {
		return models.AspectDetailedPage{}, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return models.AspectDetailedPage{}, err
	}
	return parseAspectDetailed(doc), nil
}

func parseAspectDetailed(doc *goquery.Document) models.AspectDetailedPage {
	var page models.AspectDetailedPage
	var elements []models.AspectDetailedPageElement
	var element *models.AspectDetailedPageElement
	var descriptionBuilder strings.Builder
	var elementDescriptionBuilder strings.Builder

	doc.Find("h1.entry-title.fl-l").Each(func(i int, s *goquery.Selection) {
		page.Header = s.Text()
	})

	doc.Find("#tr-content").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(i int, s *goquery.Selection) {
			switch goquery.NodeName(s) {
			case "p":
				text := s.Text()
				if element == nil {
					if descriptionBuilder.Len() > 0 {
						descriptionBuilder.WriteString("\n")
					}
					descriptionBuilder.WriteString(text)
				} else {
					if elementDescriptionBuilder.Len() > 0 {
						elementDescriptionBuilder.WriteString("\n")
					}
					elementDescriptionBuilder.WriteString(text)
				}
			case "h2", "h3":
				if element != nil {
					element.Description = elementDescriptionBuilder.String()
					elements = append(elements, *element)
					elementDescriptionBuilder.Reset()
				}
				element = &models.AspectDetailedPageElement{
					Header: s.Text(),
				}
			}
		})
	})

	if element != nil {
		element.Description = elementDescriptionBuilder.String()
		elements = append(elements, *element)
	}

	page.Description = descriptionBuilder.String()
	page.AspectParagraphs = elements
	return page
}
