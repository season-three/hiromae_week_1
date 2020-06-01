package main

import (
	"fmt"
	"encoding/csv"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	url := "https://news.google.com/topstories?hl=ja&gl=JP&ceid=JP:ja"

	//構造体の作成
	type title struct {
		id      int
		title string
		URL string
	}

	//httpアクセス
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	title, _ := doc.Find("h3.ipQwMb")
	fmt.Println(title)



