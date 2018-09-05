package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get livedoor xml to json

// RSS2 は 全体
type RSS2 struct {
	XMLName       xml.Name `xml:"rss"`
	Version       string   `xml:"version,attr"`
	Title         string   `xml:"channel>title"`
	Link          string   `xml:"channel>link"`
	Description   string   `xml:"channel>description"`
	LastBuildDate string   `xml:"channel>lastBuildDate"`
	ItemList      []Item   `xml:"channel>item"`
}

// Item は 個別
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Mobile      int    `xml:"mobile"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

// json用

type JSONItem struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Mobile      int    `json:"mobile"`
	PubDate     string `json:"pubDate"`
	GUID        string `json:"guid"`
}

type JSON struct {
	ItemList []JSONItem `json:"items"`
}

func main() {
	response, httpErr := http.Get("http://news.livedoor.com/topics/rss.xml")

	if httpErr != nil {
		fmt.Printf("Error GET: %v\n", httpErr)
		return
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	result := RSS2{}
	unmarshalErr := xml.Unmarshal([]byte(body), &result)
	if unmarshalErr != nil {
		fmt.Printf("error: %v", unmarshalErr)
		return
	}

	// for _, item := range result.ItemList {
	// 	fmt.Printf("%v\n", item.Title)
	// 	fmt.Printf("%v\n", item.Description)
	// }

	var jsoon JSON

	for _, item := range result.ItemList {
		// fmt.Printf("%v\n", item.Description)

		jsoon.ItemList = append(jsoon.ItemList, JSONItem{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			Mobile:      item.Mobile,
			PubDate:     item.PubDate,
			GUID:        item.GUID,
		})
	}

	bbb, err := json.Marshal(jsoon)
	if err != nil {
		fmt.Println("json err:", err)
	}
	// fmt.Println(string(bbb))
	out := new(bytes.Buffer)
	// プリフィックスなし、スペース4つでインデント
	json.Indent(out, bbb, "", "    ")
	fmt.Println(out.String())
}
