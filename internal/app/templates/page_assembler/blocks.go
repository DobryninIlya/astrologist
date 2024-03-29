package page_assembler

import (
	"fmt"
	"path/filepath"
	"strings"
)

var (
	aspectListElement            string
	planetAspectDescriptionBlock string
	planetAspectParagraphBlock   string
)

func getAspectListElement(link, aspectName string) string {
	if aspectListElement == "" {
		aspectListElement = readFile(filepath.Join(blocksPath, "aspect_list_element.html"))
	}
	aspectParts := strings.Split(aspectName, " ")
	aspect := strings.Join(aspectParts[1:], " ")
	return fmt.Sprintf(aspectListElement, getSuffix(link), aspectParts[0], aspect)
}

func getPlanetAspectDescription(name, degree, description string) string {
	if name == "" || degree == "" || description == "" {
		return ""
	}
	if planetAspectDescriptionBlock == "" {
		planetAspectDescriptionBlock = readFile(filepath.Join(blocksPath, "planet_aspect_description.html"))
	}
	fmt.Println(fmt.Sprintf(planetAspectDescriptionBlock, name, degree, description))
	return fmt.Sprintf(planetAspectDescriptionBlock, name, degree, description)
}

func getAspectParagraph(name, description string) string {
	if planetAspectParagraphBlock == "" {
		planetAspectParagraphBlock = readFile(filepath.Join(blocksPath, "planet_aspect_paragraph.html"))
	}
	return fmt.Sprintf(planetAspectParagraphBlock, name, description)
}

func getSuffix(url string) string {
	return strings.TrimPrefix(url, "https://geocult.ru/natalnaya-karta/")
}
