package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseUrl = "https://www.fmkorea.com/index.php?mid=best&page=5"

func main() {
	getPages()
}

func getPages() {
	res, err := http.Get(baseUrl)
	if err != nil || res.StatusCode > 400 {
		log.Fatalln("Request failed with status code", res.StatusCode)
	}
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln("Error loading HTTP response body. ", err)
	}
	defer res.Body.Close()
	
	doc.Find(".bd_pg").Each(func(i int, s *goquery.Selection) {
		log.Println(s.Html())
	})
}