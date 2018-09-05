package livedoor

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
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

// JSONItem はjsonの全体
type JSONItem struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Mobile      int    `json:"mobile"`
	PubDate     string `json:"pubDate"`
	GUID        string `json:"guid"`
}

// JSON はjsonの個別
type JSON struct {
	ItemList []JSONItem `json:"items"`
}

// LiveDoor は構造体
type LiveDoor struct {
	XML []byte
}

// NewRSS は コンストラクタ
func NewRSS(livedoorXML []byte) *LiveDoor {
	l := new(LiveDoor)
	l.XML = livedoorXML
	return l
}

// Parse 文字列をparse
func (livedoor *LiveDoor) Parse() (RSS2, error) {

	result := RSS2{}
	unmarshalErr := xml.Unmarshal([]byte(livedoor.XML), &result)
	if unmarshalErr != nil {
		fmt.Printf("error: %v", unmarshalErr)
		return RSS2{}, unmarshalErr
	}

	return result, nil
}

// JSONCreator は構造体
type JSONCreator struct {
	parseResult RSS2
}

func NewJSONCreator(parseResult RSS2) *JSONCreator {
	j := new(JSONCreator)
	j.parseResult = parseResult
	return j
}

// Output 文字列をparse
func (jsonCreator *JSONCreator) Output() (string, error) {
	var jsoon JSON

	result := jsonCreator.parseResult

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
		return "", err
	}
	// fmt.Println(string(bbb))

	out := new(bytes.Buffer)
	// プリフィックスなし、スペース4つでインデント
	json.Indent(out, bbb, "", "    ")
	fmt.Println(out.String())
	return out.String(), nil
}
