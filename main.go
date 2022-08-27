package main

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "math"
	// "net/http"
	// "os"

	// "golang.org/x/exp/slices"
	// "golang.org/x/net/html/charset"
	// "Workflow/Url"
	"Parser/Workflow"
)
// go mod init <module_name> and then just import "<module_name>/<pkg_name>
// var (
	// body                                                                                           []byte
	// textJson                                                                                       map[string]interface{}
	// url, urlChapters, mainPage                                                                     string
	// numberPageChapters                                                                             float64 = 1
	// lastPageNumber                                                                                 float64
	// counterDescription, countPagesImages, countRepeatGetImages, countBrenchId, checkCountChapters  int
	// arrayIdChapters, arrayApiTitles, arrayBrenchId, rus_names, description, numberChapters, Images []string
	// numberPage                                                                                     = 1
	// totalPages                                                                                     = 3
// )

func main() {
Workflow.Start()



	// getMainPage()
	// changeUrlApiTitles()
	// changeUrlChapters()
	// getImages()
}

// func getMainPage() {
	// url = ""
	// if numberPage != totalPages {
	// 	mainPage = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=" + fmt.Sprintf("%v", numberPage) + "&count=30"
	// 	url = mainPage
	// }
	// JsonParse()

	// for _, item := range textJson["content"].([]interface{}) {
	// 	rus_names = append(rus_names, fmt.Sprintf("%v", item.(map[string]interface{})["rus_name"])+",")
	// 	arrayApiTitles = append(arrayApiTitles, "https://api.remanga.org/api/titles/"+fmt.Sprintf("%v", item.(map[string]interface{})["dir"]))
	// }

	// if len(rus_names) == 30*numberPage {
	// 	fmt.Println("Page " + fmt.Sprintf("%v", numberPage) + " " + "Have all names")
	// }
	// createJson()
	// for key, item := range textJson["props"].(map[string]interface{}) {
	// 	switch key {
	// 	case "total_pages":
	// 		var i int
	// 		fmt.Sscanf(fmt.Sprint(item), "%d", &i)
	// 		totalPages = i
	// 	}
	// }


	// changeUrlApiTitles()
// }

// func changeUrlApiTitles() { 
// 	url = arrayApiTitles[counterDescription]
// 	JsonParse()
// 	Backend()
// 	createJson()
// 	if counterDescription != 29 {
// 		counterDescription++
// 		changeUrlApiTitles()
// 	} else {
// 		fmt.Println("Page " + fmt.Sprintf("%v", numberPage) + " " + "Have all descriptions")
// 		// changeUrlChapters()
// 	}
// }

// func changeUrlChapters() { 
// 	urlChapters = "https://api.remanga.org/api/titles/chapters/?branch_id=" + arrayBrenchId[countBrenchId] + "&count=60&ordering=-index&page=" + fmt.Sprint(numberPageChapters) + "&user_data=1"
// 	url = urlChapters
// 	JsonParse()
// 	if numberPageChapters != lastPageNumber {
// 		doUrlChapter()
// 		changePageChapters()
// 		createJson()
// 		changeUrlChapters()
// 	} else {
// 		if countBrenchId != len(arrayBrenchId)-1 {
// 			countBrenchId++
// 			numberPageChapters = 0
// 			if slices.Contains(numberChapters, "Глава 1,") {
// 				checkCountChapters++
// 			}
// 			if checkCountChapters == len(arrayBrenchId)-1 {
// 				fmt.Println("Page " + fmt.Sprintf("%v", numberPage) + " " + "Have all chapters")
// 			}
// 			changeUrlChapters()
// 		} else {
// 			countBrenchId = 0
// 			numberPageChapters = 0
// 			checkCountChapters = 0
			// getImages()
		// }
	// }
// }

// func getImages() { 
// 	if url != arrayIdChapters[len(arrayIdChapters)-1] {
// 		url = arrayIdChapters[countRepeatGetImages]
// 	}
// 	JsonParse()
// 	getJsonImages()
// 	countPagesImages = 0
// 	createJson()
// 	if url != arrayIdChapters[len(arrayIdChapters)-1] {
// 		countRepeatGetImages++
// 		getImages()
// 	} else {
// 		fmt.Println("Page " + fmt.Sprintf("%v", numberPage) + " " + "Have all images")
// 		numberPage++
// 		countRepeatGetImages = 0
// 		counterDescription = 0
// 		url = ""
// 		arrayApiTitles = nil
// 		arrayIdChapters = nil
// 		arrayBrenchId = nil
// 		main()
// 	}
// }

// func getJsonImages() {
	// for key, value := range textJson["content"].(map[string]interface{}) {
// 		switch key {
// 		case "pages":
// 			for key, value := range value.([]interface{}) {
// 				switch key {
// 				case countPagesImages:
// 					countPagesImages++
// 					if fmt.Sprintf("%T", value) == "[]interface {}" {
// 						for _, value := range value.([]interface{}) {
// 							Images = append(Images, fmt.Sprintf("%v", value.(map[string]interface{})["link"])+",")
// 						}
// 					} else {
// 						for key, value := range value.(map[string]interface{}) {
// 							switch key {
// 							case "link":
// 								Images = append(Images, fmt.Sprintf("%v", value)+",")
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func JsonParse() {
// 	resp, _ := http.Get(url)
// 	defer resp.Body.Close()

// 	utf8, _ := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
// 	body, _ = io.ReadAll(utf8)

// 	json.Unmarshal(body, &textJson)
// }

// func doUrlChapter() {
// 	if fmt.Sprintf("%T", textJson["content"]) == "[]interface {}" {
// 		for _, item := range textJson["content"].([]interface{}) {
// 			numberChapters = append(numberChapters, "Глава "+fmt.Sprintf("%v", item.(map[string]interface{})["chapter"])+",")
// 			arrayIdChapters = append(arrayIdChapters, "https://api.remanga.org/api/titles/chapters/"+fmt.Sprintf("%v", item.(map[string]interface{})["id"]))
// 		}
// 	}
// }

// func changePageChapters() {
// 	var i int64
// 	fmt.Sscan(numberChapters[0][11:len(numberChapters[0])-1], &i)
// 	lastPageNumber = math.Round(float64(i)/60) + 1
// 	if numberPageChapters < lastPageNumber {
// 		numberPageChapters++
// 	}
// }

// func Backend() {
// 	for key, value := range textJson["content"].(map[string]interface{}) {
// 		switch key {
// 		case "description":
// 			description = append(description, fmt.Sprintf("%v", value)+",")
// 		case "branches":
// 			for _, item := range value.([]interface{}) {
// 				arrayBrenchId = append(arrayBrenchId, fmt.Sprintf("%v", item.(map[string]interface{})["id"]))
// 			}
// 		}
// 	}
// }

// func createJson() {

// 	parserData, _ := os.Create("parserData.json")
// 	parserData.Close()

// 	f, _ := os.OpenFile("parserData.json", os.O_CREATE|os.O_RDWR, 0777)

// 	f.WriteString("{\"Name\":" + "\"" + fmt.Sprintf("%v", rus_names) + "\"")
// 	f.WriteString("\"Description\":" + "\"" + fmt.Sprintf("%v", description) + "\"" + ",")
// 	f.WriteString("\"Chapters\":" + "\"" + fmt.Sprintf("%v", numberChapters) + "\"" + ",")
// 	f.WriteString("\"Images\":" + "\"" + fmt.Sprintf("%v", Images) + "\"}")
// 	f.Close()
// }
