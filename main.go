package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	url := "https://go-tour-jp.appspot.com/welcome/1"

	//httpアクセス
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	//タイトル取ってくる
	title := doc.Find("title").Text()
	titleresult := "ページタイトル:" + title
	fmt.Println(titleresult)

	//DB接続
	db, err := gorm.Open("mysql", "scraper:11194222@tcp(localhost:3306)/scraping?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer db.Close()

	//構造体の作成
	var scraping.go struct {
		id      int
		content string
	}

	//データの追加
	error := db.Create("titleresult")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}

}
