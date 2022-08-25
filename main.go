package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html/charset"
)

var (
	body                                                                       []byte
	textJson                                                                   map[string]interface{}
	rus_names, url, description, Images, urlChapters, numberChapters, mainPage string
	numberPageChapters                                                         float64 = 1
	lastPageNumber                                                             float64
	counterDescription, countPagesImages, countRepeatGetImages, countBrenchId  int
	arrayIdChapters, arrayApiTitles, arrayBrenchId                             []string
	numberPage                                                                 = 1
	totalPages                                                                 = 2
	wg                                                                         = sync.WaitGroup{}
)

func main() {
	for countRepetitions := 1; countRepetitions < 7; countRepetitions++ {
		urlGet()
	}
}

// получить Description chapters img 1 манги 2 страницы последней главы   => готово
func urlGet() {
	fmt.Println(url)
	switch url {
	case "":
		getMainPage()
	case mainPage:
		changeUrlApiTitles()
	case arrayApiTitles[counterDescription]:
		changeUrlChapters()
	case urlChapters:
		getImages()
	}
}

func getMainPage() {
	if numberPage != totalPages {
		mainPage = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=" + fmt.Sprintf("%v", numberPage) + "&count=30"
		url = mainPage
	}
	numberPage++
	parser()
	for _, item := range textJson["content"].([]interface{}) {
		rus_names += fmt.Sprintf("%v", item.(map[string]interface{})["rus_name"]) + ","
		arrayApiTitles = append(arrayApiTitles, "https://api.remanga.org/api/titles/"+fmt.Sprintf("%v", item.(map[string]interface{})["dir"]))
	}
	for key, item := range textJson["props"].(map[string]interface{}) {
		switch key {
		case "total_pages":
			var i int
			fmt.Sscanf(fmt.Sprint(item), "%d", &i)
			totalPages = i
		}
		createJson()
		wg.Add(3)
	}
}
func changeUrlApiTitles() {
	url = arrayApiTitles[counterDescription]
	parser()
	Backend()
	go func() {
		createJson()
	}()
	if counterDescription != 29 {
		counterDescription++
		changeUrlApiTitles()
	} else {
		wg.Done()
	}
}

func changeUrlChapters() {
	urlChapters = "https://api.remanga.org/api/titles/chapters/?branch_id=" + arrayBrenchId[countBrenchId] + "&count=60&ordering=-index&page=" + fmt.Sprint(numberPageChapters) + "&user_data=1"
	url = urlChapters
	parser()

	go func() {
		changePageChapters()
		createJson()
	}()

	if numberPageChapters != lastPageNumber {
		doUrlChapter()
		changeUrlChapters()
	} else {
		if countBrenchId != 31 {
			countBrenchId++
			numberPageChapters = 0
			changeUrlChapters()
		} else {
			countBrenchId = 0
			wg.Done()
		}
	}
}

func getImages() {
	// fmt.Println(len(arrayIdChapters))
	url = arrayIdChapters[countRepeatGetImages] //index out of range [2915] with length 2915
	parser()
	for key, value := range textJson["content"].(map[string]interface{}) {
		switch key {
		case "pages":
			for key, value := range value.([]interface{}) {
				switch key {
					case countPagesImages:
						countPagesImages++
						if fmt.Sprintf("%T", value) == "[]interface {}" {
							for _, value := range value.([]interface{}) {
								Images += fmt.Sprintf("%v", value.(map[string]interface{})["link"]) + ","
							}
							} else {
								for key, value := range value.(map[string]interface{}) {
								switch key {
								case "link":
									Images += fmt.Sprintf("%v", value) + ","
								}
							}
						}
					}
				}
			}
		}
		go func() {
		createJson()
	}()
	countPagesImages = 0
	if countRepeatGetImages < len(arrayIdChapters) {
		countRepeatGetImages++
		getImages()
	} else {
			countRepeatGetImages = 0
			counterDescription = 0
			url = ""
			arrayApiTitles = nil
			arrayIdChapters = nil
			arrayBrenchId = nil
			defer main()
			wg.Done()
			wg.Wait()
		}
	}
	// if countRepeatGetImages == len(arrayIdChapters)-1 {
	// 	countRepeatGetImages = 0
	// 	counterDescription = 0
	// 	url = ""
	// 	wg.Done()
	// 	wg.Wait()
	// 	arrayApiTitles = nil
	// 	arrayIdChapters = nil
	// 	arrayBrenchId = nil
	// 	defer main()
	// }
// }

func parser() {
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
