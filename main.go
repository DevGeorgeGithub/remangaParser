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

	mainPage                                                                  = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=1&count=30"
	rus_names, url, description, Images, urlChapters, numberChapters          string
	numberPageChapters                                                        float64 = 1
	lastPageNumber                                                            float64
	counterDescription, countPagesImages, countRepeatGetImages, countBrenchId int
	arrayIdChapters, arrayApiTitles, arrayBrenchId                            []string
)

func main() {
	// мен mainpage defer main()
	for countRepetitions := 1; countRepetitions < 7; countRepetitions++ {
		urlGet()
	}
}

func urlGet() {
	switch url {
	case "":
		getMainPage()
	case mainPage:
		changeUrlApiTitles()
	case arrayApiTitles[counterDescription]:
		changeUrlChapters()
	case urlChapters:
		getImages()
		// case arrayIdChapters[countRepeatGetImages]:
		// 	url = ""
	}
}

func getMainPage() {
	url = mainPage
	parser()
	for _, item := range textJson["content"].([]interface{}) {
		rus_names += fmt.Sprintf("%v", item.(map[string]interface{})["rus_name"]) + ","
		arrayApiTitles = append(arrayApiTitles, "https://api.remanga.org/api/titles/"+fmt.Sprintf("%v", item.(map[string]interface{})["dir"]))
	}
	createJson()
}

func getImages() {
	// if countRepeatGetImages < len(arrayIdChapters) && Images != "" {
	// 	countRepeatGetImages++
	// }
	url = arrayIdChapters[countRepeatGetImages]
	parser()

	for key, value := range textJson["content"].(map[string]interface{}) {
		switch key {
		case "pages":
			for key, value := range value.([]interface{}) {
				switch key {
				case countPagesImages:
					countPagesImages++
					for _, value := range value.([]interface{}) {
						Images += fmt.Sprintf("%v", value.(map[string]interface{})["link"]) + ","
					}
				}
			}
		}
	}
	createJson()
	countPagesImages = 0
	if countRepeatGetImages < len(arrayIdChapters) { // index out of range [59] with length 59
		countRepeatGetImages++
		defer getImages()
	}
}

func changeUrlApiTitles() {
	url = arrayApiTitles[counterDescription]
	parser()
	Backend()
	createJson()
	if counterDescription != 29 {
		counterDescription++
		defer changeUrlApiTitles()
	}
}

func changeUrlChapters() {
	urlChapters = "https://api.remanga.org/api/titles/chapters/?branch_id=" + arrayBrenchId[countBrenchId] + "&count=60&ordering=-index&page=" + fmt.Sprint(numberPageChapters) + "&user_data=1"
	url = urlChapters
	parser()
	doUrlChapter()
	changePageChapters()
	createJson()

	if numberPageChapters != lastPageNumber {
		defer changeUrlChapters()
	} else {
		if countBrenchId != 31 {
			countBrenchId++
			numberPageChapters = 0
			changeUrlChapters()
		}
	}
}

func parser() {
	// fmt.Println(url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	utf8, _ := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	body, _ = io.ReadAll(utf8)

	json.Unmarshal(body, &textJson)
}

func doUrlChapter() {
	for _, item := range textJson["content"].([]interface{}) {
		numberChapters += "Глава " + fmt.Sprintf("%v", item.(map[string]interface{})["chapter"]) + ","
		arrayIdChapters = append(arrayIdChapters, "https://api.remanga.org/api/titles/chapters/"+fmt.Sprintf("%v", item.(map[string]interface{})["id"]))
	}
}

func changePageChapters() {
	var i int64
	fmt.Sscanf(numberChapters[11:strings.Index(numberChapters, ",")], "%d", &i)
	lastPageNumber = math.Round(float64(i)/60) + 1
	if numberPageChapters < lastPageNumber {
		numberPageChapters++
	}
}

func Backend() {
	for key, value := range textJson["content"].(map[string]interface{}) {
		switch key {
		case "description":
			description += fmt.Sprintf("%v", value)
		case "branches":
			for _, item := range value.([]interface{}) {
				arrayBrenchId = append(arrayBrenchId, fmt.Sprintf("%v", item.(map[string]interface{})["id"]))
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
