package Workflow

import (
	"fmt"
)

var (
	numberPage                                                                                     = 1
	totalPages                                                                                     = 3
	mainPage, url string
)

func Start() {
	if numberPage != totalPages {
		mainPage = "https://api.remanga.org/api/search/catalog/?ordering=-rating&page=" + fmt.Sprintf("%v", numberPage) + "&count=30"
		url = mainPage
	}
	JsonParse()
}
