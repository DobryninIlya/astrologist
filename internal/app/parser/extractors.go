package parser

import (
	"astrologist/internal/app/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func extractAstroData(doc *goquery.Document, id string) (models.AstroData, error) {

	var data models.AstroData
	page := doc.Find("div.entry-content.article")
	page.Children().Find("*").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "Читать далее...") || strings.Contains(s.Text(), "Читать далее") || strings.Contains(s.Text(), "Подробнее") {
			s.Remove()
		}
	})

	page.Children().Each(func(i int, s *goquery.Selection) {
		target := s.Find("#" + id)
		if target.Length() > 0 {
			target = target.Parent().Parent()
			s = target.NextAll()
			var stopIteration bool
			s.Each(func(i int, s *goquery.Selection) {
				if stopIteration {
					return
				}
				if s.Is("p") {
					// <p> - <strong> - <span>, <span> СИЛА, ГАРМОНИЯ
					if s1 := s.Find("strong"); s1.Length() > 0 {
						if s2 := s1.Find("span"); s2.Length() == 2 {
							data.Power, _ = strconv.ParseFloat(strings.TrimSpace(s2.Eq(0).Text()), 64)
							data.Harmony, _ = strconv.ParseFloat(strings.TrimSpace(s2.Eq(1).Text()), 64)
						}
						//} else if s1 := s.Find("img"); s1.Length() > 0 { // <p2> Text()
						//	data.Description = s.Text()
					} else if attr, ok := s.Attr("style"); attr == "text-align: justify;" && ok && data.AspectName == "" {
						data.Description += s.Text()
						//} else if s1 := s.Find("b"); s1.Length() > 0 { // <p> - <b> Text()
					} else if attr, ok := s.Attr("style"); attr == "text-align: justify;" && ok && data.AspectName != "" {
						data.AspectDesc = s.Text()
					}
				} else if s.Is("font") { // <font>
					if s1 := s.Find("h3"); s1.Length() > 0 { // <font> - <h3> Text()
						content := s1.Text()
						contentParts := strings.Split(content, "\n")
						nameParts := strings.Split(contentParts[0], "  ")
						if len(nameParts) != 2 {
							return
						}
						data.Position = nameParts[1]
						degreeParts := strings.Split(contentParts[1], " ")
						data.Degree = degreeParts[2]
					}
					if s1 := s.Find("p"); s1.Length() > 0 { // <font> - <img> alt="Text"

						data.PositionMeaning = strings.ReplaceAll(s1.Text(), "Читать далее...", " ")
						//data.PositionMeaning = s1.Text()
					}

				} else if s.Is("div") { // <div class="block"> - <div class="name"> - <li> - <a href=...> Text()
					if attr, ok := s.Attr("class"); attr == "block" && ok {
						if data.AspectList == nil {
							data.AspectList = make([]models.AspectList, 0)
						}
						s.Find("li").Each(func(i int, s *goquery.Selection) {
							var aspect models.AspectList
							if s1 := s.Find("a"); s1.Length() > 0 {
								aspect.AspectName = s1.Text()
								aspect.Link, _ = s1.Attr("href")
								data.AspectList = append(data.AspectList, aspect)
							}
						})
					}
				} else if s.Is("a") {
					if attr, ok := s.Attr("href"); attr == "#show-more11" && ok {
						stopIteration = true
						return
					}
				} else if s.Is("h3") {
					data.AspectName = s.Text()
				}

			})
		}
	})
	return data, nil
}

func extractKarmaData(doc *goquery.Document) (models.KarmaData, error) {
	var data models.KarmaData
	page := doc.Find("div.entry-content.article")
	page.Children().Each(func(i int, s *goquery.Selection) {
		target := s.Find("#carmanat")
		if target.Length() > 0 {
			target = target.Parent()
			s = target.NextAll()
			var stopIteration bool
			s.Each(func(i int, s *goquery.Selection) {
				if stopIteration {
					return
				}
				if s.Is("p") {
					if attr, ok := s.Attr("style"); attr == "text-align: justify;" && ok {
						if data.Description == "" {
							data.Description = s.Text()
						} else if data.DetailedDescription == "" {
							data.DetailedDescription = s.Text()
						}
					}
				} else if s.Is("h3") {
					if data.Header == "" {
						data.Header = s.Text()
					}
				} else if s.Is("a") {
					if attr, ok := s.Attr("href"); attr == "#show-more11" && ok {
						stopIteration = true
						return
					}
				}
			})
		}
	})
	return data, nil
}

func extractTables(doc *goquery.Document) ([]string, error) {
	var tables []string

	// Находим все таблицы в документе
	doc.Find("table").Each(func(i int, tableSelection *goquery.Selection) {
		// Получаем HTML-код таблицы
		tableHTML, _ := tableSelection.Html()
		// Добавляем HTML-код таблицы в список
		tables = append(tables, tableHTML)
	})

	if len(tables) < 4 {
		return nil, fmt.Errorf("not enough planet tables found")
	} // Возвращаем только первые четыре таблицы
	return tables[1:4], nil
}

func extractMoonNodes(doc *goquery.Document) (models.MoonNodes, error) {
	var data models.MoonNodes
	page := doc.Find("div.entry-content.article")
	page.Children().Each(func(i int, s *goquery.Selection) {
		target := s.Find("#nodemoonat")
		if target.Length() > 0 {
			target = target.Parent().Parent()
			s = target.NextAll()
			var stopIteration bool
			s.Each(func(i int, s *goquery.Selection) {
				if stopIteration {
					return
				}
				if s.Is("p") {
					s.Find("a").Each(func(i int, s *goquery.Selection) {
						if val, ok := s.Attr("target"); val == "_blank" && ok {
							s.Remove()
						}
					})
					if attr, ok := s.Attr("style"); attr == "text-align: justify;" && ok {
						if data.Description == "" {
							data.Description = s.Text()
						} else if data.DetailedDescription == "" {
							data.DetailedDescription = s.Text()
						} else {
							data.DetailedDescription += s.Text()
						}
					}
					if s1 := s.Find("b"); s1.Length() > 0 {
						if data.AspectDescription == "" {
							data.AspectDescription = s1.Text()
						}
					}
				} else if s.Is("h3") {
					if data.Header == "" {
						data.Header = s.Text()
					}
				} else if s.Is("a") {
					if attr, ok := s.Attr("href"); attr == "#show-more11" && ok {
						stopIteration = true
						return
					}
				} else if s.Is("div") { // <div class="block"> - <div class="name"> - <li> - <a href=...> Text()
					if attr, ok := s.Attr("class"); attr == "block" && ok {
						if data.AspectList == nil {
							data.AspectList = make([]models.AspectList, 0)
						}
						s.Find("li").Each(func(i int, s *goquery.Selection) {
							var aspect models.AspectList
							if s1 := s.Find("a"); s1.Length() > 0 {
								aspect.AspectName = s1.Text()
								aspect.Link, _ = s1.Attr("href")
								data.AspectList = append(data.AspectList, aspect)
							}
						})
					}
				}
			})
		}
	})
	return data, nil
}

func extractSelenaData(doc *goquery.Document) (models.SelenaData, error) {
	var data models.SelenaData
	page := doc.Find("div.entry-content.article")
	page.Children().Each(func(i int, s *goquery.Selection) {
		target := s.Find("#selenanat")
		if target.Length() > 0 {
			target = target.Parent().Parent()
			s = target.NextAll()
			var stopIteration bool
			s.Each(func(i int, s *goquery.Selection) {
				if stopIteration {
					return
				}
				if s.Is("p") {
					s.Find("a").Each(func(i int, s *goquery.Selection) {
						if val, ok := s.Attr("target"); val == "_blank" && ok {
							if strings.Contains(s.Text(), "...") {
								s.Remove()
							}
						}
					})
					if attr, ok := s.Attr("style"); attr == "text-align: justify;" && ok {
						if data.Description == "" {
							data.Description = s.Text()
						} else if data.DetailedDescription == "" {
							data.DetailedDescription = s.Text()
						} else {
							data.DetailedDescription += s.Text()
						}
					}
				} else if s.Is("a") {
					if attr, ok := s.Attr("href"); attr == "#show-more11" && ok {
						stopIteration = true
						return
					}
				} else if s.Is("div") {
					if attr, ok := s.Attr("align"); ok && attr == "left" {
						data.Header = s.Text()
					}

				}
			})
		}
	})
	return data, nil
}

func extractLilitData(doc *goquery.Document) (models.LilitData, error) {
	var data models.LilitData
	page := doc.Find("div.entry-content.article")
	page.Children().Each(func(i int, s *goquery.Selection) {
		target := s.Find("#lilitnat")
		if target.Length() > 0 {
			target = target.Parent().Parent()
			s = target.NextAll()
			var stopIteration bool
			s.Each(func(i int, s *goquery.Selection) {
				if stopIteration {
					return
				}
				if s.Is("p") {
					s.Find("a").Each(func(i int, s *goquery.Selection) {
						if val, ok := s.Attr("target"); val == "_blank" && ok {
							if strings.Contains(s.Text(), "...") {
								s.Remove()
							}
						}
					})
					if attr, ok := s.Attr("style"); attr == "text-align: justify;" && ok {
						if data.DetailedDescription == "" && data.Header == "" {
							data.Description += s.Text()
						} else if data.Header != "" {
							data.DetailedDescription += s.Text()
						}
					}
				} else if s.Is("a") {
					if attr, ok := s.Attr("href"); attr == "#show-more11" && ok {
						stopIteration = true
						return
					}
				} else if s.Is("div") {
					if attr, ok := s.Attr("align"); ok && attr == "left" {
						data.Header = s.Text()
					}

				}
			})
		}
	})
	return data, nil
}
