package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"golang.org/x/net/html/charset"
)

var (
	body []byte
	textJson map[string]interface{}
	
	defaultUrl = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=1&count=30"
	dirs string
	rus_names string
	chapter string
	url string
	description string
	Images string
	chaptersUrl string
	numberChapters string

	api []string
	arrayDirs []string
)

func main() {

	urlGet()
	parser()
	valuesJson()
	createJson()
}

func urlGet() {

	if url == "" {
		url = defaultUrl
	} else if url == defaultUrl {
		url = api[0]
	} else if url == api[0] {
		url = chapter
	}else if url == chapter {
		url = chaptersUrl
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

	if url == defaultUrl || url == chaptersUrl{
		for _, item := range textJson["content"].([]interface{}) {
			dirs += fmt.Sprintf("%v", item.(map[string]interface{})["dir"])
			api = append(api, "https://api.remanga.org/api/titles/"+dirs)
			arrayDirs = append(arrayDirs, "https://remanga.org/manga/"+dirs)
	
			if url == defaultUrl {
				rus_names += fmt.Sprintf("%v", item.(map[string]interface{})["rus_name"]) + ","
			}else {
				numberChapters += "Глава " + fmt.Sprintf("%v", item.(map[string]interface{})["chapter"]) + ","
			}
		}
	}

	if url == api[0] || url == chapter {
		mapJson := textJson["content"].(map[string]interface{})
		for key, value := range mapJson {
			switch key {
				case "description":
					description += fmt.Sprintf("%v", value)
				case "branches":
					for _, item := range value.([]interface{}) {
						chaptersUrl = "https://api.remanga.org/api/titles/chapters/?branch_id=" + fmt.Sprintf("%v", item.(map[string]interface{})["id"]) + "&count=60&ordering=-index&page=1&user_data=1"
					}
				case "first_chapter":
					for key, value2 := range value.(map[string]interface{}) {
						if key == "id" {
							chapter = "https://api.remanga.org/api/titles/chapters/" + fmt.Sprintf("%v", value2)
						}
					}
				case "pages":
					for _, item := range value.([]interface{}) {
						Images += fmt.Sprintf("%v", item.(map[string]interface{})["link"]) + ", "
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

	if url != chaptersUrl  {
		defer main()
	}
}

