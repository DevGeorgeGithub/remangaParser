package Workflow

import (
	"encoding/json"
	"io"
	"net/http"
	"golang.org/x/net/html/charset"
	"Parser/SiteElements"

)

var (
	body     []byte
	textJson map[string]interface{}
)

func JsonParse() map[string]interface{}{
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	utf8, _ := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	body, _ = io.ReadAll(utf8)

	json.Unmarshal(body, &textJson)
	SiteElements.GetNames()
	return textJson
}
