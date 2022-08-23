package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"os"
	"strings"
)

var body []byte
var textJson map[string]interface{}

var dirs string
var rus_names string
var chapter string
var url string
var api string
var siteDirs string
var description string

var arrayDirs []string
var arrayRus_names []string
var arraySiteDirs []string
var arrayApi []string

func main() {
	check()
	parser()
	valuesJson()
	createJson()
}

func check() {

	if len(arrayRus_names) == 34 {
		url = arrayApi[0]
	} else {
		url = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=1&count=30"
	}
}

func parser() {
	// fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	utf8, _ := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))

	body, _ = io.ReadAll(utf8)

}

func valuesJson() {

	json.Unmarshal(body, &textJson)

	if url != "https://api.remanga.org/api/titles/solo-leveling" {

		for _, item := range textJson["content"].([]interface{}) {
			dirs += fmt.Sprintf("%v", item.(map[string]interface{})["dir"]) + ","
			rus_names += fmt.Sprintf("%v", item.(map[string]interface{})["rus_name"]) + ","
		}
	}
	arrayDirs = strings.Split(dirs, ",")
	arrayRus_names = strings.Split(rus_names, ",")
	if url == "https://api.remanga.org/api/titles/solo-leveling" {

		mapJson := textJson["content"].(map[string]interface{})
		for key, value := range mapJson {
			switch key {
			case "description":
				description += fmt.Sprintf("%v", value)
			case "first_chapter":
				for key, value2 := range value.(map[string]interface{}) {
				    if key == "id" {
				        chapter = fmt.Sprintf("%v", value2)
				    }
				}
			}
		}
	}
	chapter = "https://remanga.org/manga/" + arrayDirs[0] + "/" + "ch" + chapter
    fmt.Println(chapter)
	for i := range arrayDirs {
		siteDirs += "https://remanga.org/manga/" + arrayDirs[i] + ","
		api += "https://api.remanga.org/api/titles/" + arrayDirs[i] + ","
	}

	arraySiteDirs = strings.Split(siteDirs, ",")
	arrayApi = strings.Split(api, ",")

}

func createJson() {
	parserData, _ := os.Create("parserData.json")
	parserData.Close()

	f, _ := os.OpenFile("parserData.json", os.O_CREATE|os.O_RDWR, 0777)

	f.WriteString("{\"Name\":" + "\"" + rus_names + "\" ")
	f.WriteString("\"Description\":" + "\"" + description + "\"")
	// f.WriteString("\"Images\":" + "\"" + rus_names + "\"}")
	f.Close()
	if url != arrayApi[0] {
		defer main()
	}
}
