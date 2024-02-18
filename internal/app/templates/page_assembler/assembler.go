package page_assembler

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	path       string
	htmlPath   string
	blocksPath string
)

func init() {
	path = filepath.Join("internal", "app", "templates")
	htmlPath = filepath.Join(path, "html")
	blocksPath = filepath.Join(htmlPath, "blocks")
}

func readFile(path string) string {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("file does not exist")
			return ""
		}
	}
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer file.Close()

	var rows []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	return strings.Join(rows, "\n")

}
