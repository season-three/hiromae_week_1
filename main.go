package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	url := "https://news.google.com/topstories?hl=ja&gl=JP&ceid=JP:ja"

	//構造体の作成
	type title struct {
		id    int
		title string
		URL   string
	}

	//httpアクセス
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	//タイトル名とURLを取得。Eachを使いたい。
	doc.Find(".ipQwMb a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})

	doc.Find(".ipQwMb a").Each(func(s *Selection) Text() string {
		text, _ := s.Text()
		fmt.Println(text)

	})

}
