package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/egusa_go_api/news_api/livedoor"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hogehoge, _ := main2()

	// fmt.Println(hogehoge)

	// res, err := json.Marshal(hogehoge)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(hogehoge))
}

func main() {
	http.HandleFunc("/news", handler)
	http.ListenAndServe(":8000", nil)
}

func main2() (string, error) {
	response, httpErr := http.Get("http://news.livedoor.com/topics/rss.xml")

	if httpErr != nil {
		fmt.Printf("Error GET: %v\n", httpErr)
		return "", httpErr
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	ld := livedoor.NewRSS(body)

	parseResult, err := ld.Parse()

	if err != nil {
		return "", err
	}

	JSONCreator := livedoor.NewJSONCreator(parseResult)

	jsonResult, err := JSONCreator.Output()

	if err != nil {
		return "", err
	}

	return jsonResult, nil
}
