package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html/charset"
)

var (
	body     []byte
	textJson map[string]interface{}

	mainPage           = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=1&count=30" 
	rus_names, urlImagesTitle, url, description, Images, urlChapters, brenchId, numberChapters string
	numberPageChapters float64 = 1
	lastPageNumber     float64
	counterDescription int
	arrayIdChapters, arrayApiTitles	[]string
)

func main() {
	// мен mainpage defer main()
	// for countRepetitions := 1; countRepetitions < 32; countRepetitions++ {
	for countRepetitions := 1; countRepetitions < 7; countRepetitions++ {
		urlGet()
		parser()
		valuesJson()
		createJson()
		fmt.Println(url)
	}
}

func urlGet() {
	switch url {
	case "":
		url = mainPage
	case mainPage:
		changeUrlApiTitles()
	case changeUrlApiTitles():
		changeUrlChapters()
	case changeUrlChapters():
		getImages()
	}
}

func getImages() string { // разбить на функ без повт
	url = arrayIdChapters[0]
	parser()
	json.Unmarshal(body, &textJson)

	for key, value := range textJson["content"].(map[string]interface{}) {
		switch key {
		case "pages":
			for key, value := range value.([]interface{}) {
				switch key {
				// 0-6 for
				case 0:
					for _, value := range value.([]interface{}) {
						Images += fmt.Sprintf("%v",value.(map[string]interface{})["link"])
					}
				}
			}
		}
	}
	return url
}

// textJson["content"].(map[string]interface{}) switch deepJson   textJson["content"].([]interface{}) переменные getJsonValues
// func getJsonValues() 



	// for key, value := range textJson["content"].(map[string]interface{}) {
	// 	switch key {
	// 	case "pages":
	// 		for key, value := range value.([]interface{}) {
	// 			switch key {
	// 			// 0-6 for
	// 			case 0:
	// 				for _, value := range value.([]interface{}) {
	// 					Images += fmt.Sprintf("%v",value.(map[string]interface{})["link"])
	// 				}
	// 			}
	// 		}
	// 	}
	// }
// }



func changeUrlApiTitles() string {
	// if description != "" && counterDescription != 29{
		// counterDescription++
	// }
	url = arrayApiTitles[counterDescription]
	return url
}

func changeUrlChapters() string{
	if url == urlChapters {
		changePageChapters()
	}
	urlChapters = "https://api.remanga.org/api/titles/chapters/?branch_id=" + brenchId + "&count=60&ordering=-index&page=" + fmt.Sprint(numberPageChapters) + "&user_data=1"
	url = urlChapters
	return url
}


func parser() {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	utf8, _ := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	body, _ = io.ReadAll(utf8)
}

func valuesJson() {
	json.Unmarshal(body, &textJson)
	Frontend()
	Backend()
}

func Frontend() {
	doUrlMainPage()
	doUrlChapter()
}

func doUrlMainPage() {
	if url == mainPage {
		for _, item := range textJson["content"].([]interface{}) {
			rus_names += fmt.Sprintf("%v", item.(map[string]interface{})["rus_name"]) + ","
			arrayApiTitles = append(arrayApiTitles, "https://api.remanga.org/api/titles/"+fmt.Sprintf("%v", item.(map[string]interface{})["dir"]))
		}
	}
}

func doUrlChapter() {
	if url == urlChapters {
		for _, item := range textJson["content"].([]interface{}) {
			numberChapters += "Глава " + fmt.Sprintf("%v", item.(map[string]interface{})["chapter"]) + ","
			arrayIdChapters = append(arrayIdChapters, "https://api.remanga.org/api/titles/chapters/"+fmt.Sprintf("%v", item.(map[string]interface{})["id"]))
		}
		changePageChapters()
	}
}

func changePageChapters() {
	var i int64
	fmt.Sscanf(numberChapters[11:strings.Index(numberChapters, ",")], "%d", &i)
	lastPageNumber = math.Round(float64(i) / 60)
	if numberPageChapters < lastPageNumber {
		numberPageChapters++
	}
}

func Backend() {
	if url == arrayApiTitles[counterDescription] || url == urlImagesTitle {
		for key, value := range textJson["content"].(map[string]interface{}) {
			switch key {
			case "description":
				description += fmt.Sprintf("%v", value)
			case "branches":
				for _, item := range value.([]interface{}) {
					brenchId = fmt.Sprintf("%v", item.(map[string]interface{})["id"])
				}
			}
		}
	}
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
}

