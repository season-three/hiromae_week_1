package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://go-tour-jp.appspot.com/welcome/1"

	doc, err := goquery.NewDocument(url) //①HTTPリクエスト行を入れる
	if err != nil {
		fmt.Printf("%s\n", err) //③万が一のエラーのためのコードを入れる
	}

	title, err := doc.Find("title").Text()
	if err != nil {
		fmt.Printf("%s\n", err) //③万が一のエラーのためのコードを入れる
	}

	fmt.Println("ページタイトル:" + title)

}
