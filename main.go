package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"math"
	"strings"
	"golang.org/x/net/html/charset"
)

var (
	body     []byte
	textJson map[string]interface{}

	mainPage           = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=1&count=30"
	dirs               string
	rus_names          string
	urlImagesTitle     string
	url                string
	description        string
	Images             string
	urlChapters        string
	numberChapters     string
	numberPageChapters float64 = 1
	urlChaptersId      string

	arrayApiTitles []string
	arrayDirs      []string
)

func main() {
	urlGet()
	parser()
	valuesJson()
	createJson()
}

func urlGet() {
	if url == "" {
		url = mainPage
	} else if url == mainPage {
		url = arrayApiTitles[0]
	} else if url == arrayApiTitles[0] {
		url = urlImagesTitle
	} else if url == urlImagesTitle {
		url = urlChapters
	} else if url != urlChapters {
		url = urlChapters
	}

}
func parser() {

	fmt.Println(url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	utf8, _ := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	body, _ = io.ReadAll(utf8)
}

func valuesJson() {
	json.Unmarshal(body, &textJson)
	Frontend()
	changePageChapters()
	Backend()
}

func Frontend() {
	if url == mainPage || url == urlChapters {
		for _, item := range textJson["content"].([]interface{}) {
			dirs += fmt.Sprintf("%v", item.(map[string]interface{})["dir"])
			if url == urlChapters {
				numberChapters += "Глава " + fmt.Sprintf("%v", item.(map[string]interface{})["chapter"]) + ","
			} else {
				rus_names += fmt.Sprintf("%v", item.(map[string]interface{})["rus_name"]) + ","
			}

			arrayApiTitles = append(arrayApiTitles, "https://api.remanga.org/api/titles/"+dirs)
			arrayDirs = append(arrayDirs, "https://remanga.org/manga/"+dirs)
		}
	}
}

func changePageChapters() {
	if url == urlChapters {
		var i int64
		var lastPageNumber float64

		fmt.Sscanf(numberChapters[11:strings.Index(numberChapters, ",")], "%d", &i)
		lastPageNumber = math.Round(float64(i) / 60)
		if numberPageChapters < lastPageNumber {
			numberPageChapters++
		}
	}
}

func Backend() {
	if url == arrayApiTitles[0] || url == urlImagesTitle {
		mapJson := textJson["content"].(map[string]interface{})
		for key, value := range mapJson {
			switch key {
				case "description":
					description += fmt.Sprintf("%v", value)
				case "branches":
					for _, item := range value.([]interface{}) {
						urlChaptersId = fmt.Sprintf("%v", item.(map[string]interface{})["id"])
					}
				case "first_chapter":
					for key, value2 := range value.(map[string]interface{}) {
						if key == "id" {
							urlImagesTitle = "https://api.remanga.org/api/titles/chapters/" + fmt.Sprintf("%v", value2)
						}
					}
				case "pages":
					for _, item := range value.([]interface{}) {
						Images += fmt.Sprintf("%v", item.(map[string]interface{})["link"]) + ", "
				  }
			 }
		}
	}
	urlChapters = "https://api.remanga.org/api/titles/chapters/?branch_id=" + urlChaptersId + "&count=60&ordering=-index&page=" + fmt.Sprint(numberPageChapters) + "&user_data=1"
}

func createJson() {

	parserData, _ := os.Create("parserData.json")
	parserData.Close()

	f, _ := os.OpenFile("parserData.json", os.O_CREATE|os.O_RDWR, 0777)

	f.WriteString("{\"Name\":" + "\"" + rus_names + "\"" + ",")
	f.WriteString("\"Description\":" + "\"" + description + "\"" + ",")
	f.WriteString("\"Chapters\":" + "\"" + numberChapters + "\"" + ",")
	f.WriteString("\"Images\":" + "\"" + Images + "\"}")
	f.Close()

	if url != urlChapters {
		defer main()
	}
}
