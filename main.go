package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	url := "https://news.google.com/topstories?hl=ja&gl=JP&ceid=JP:ja"

	//構造体の作成
	type titlelist struct {
		id    int
		title string
		url   string
	}

	//httpアクセス
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	//タイトル名とURLを取得。Eachを使いたい。
	doc.Find(".ipQwMb a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		title := s.Text()
		fmt.Println(title)
		fmt.Println(url)
	})

	//Goでスクレイピングしてきたデータをcsvファイルに保存

	//CSVファイル作成
	file, err := os.Create("firstgo.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//作ったCSVファイルにデータを書き出す

	//CSVファイル開く(なかったら作成する)
	file2, err := os.OpenFile("firstgo.csv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)

	//構造体の設定
	v := titlelist{title, url}

	//ループ処理
	for _, v := range titlelist { //rangeはある構造体配列をポインタの参照渡しで別の構造体配列にコンバートする際に使う
		//配列
		content := []string{v.title, v.url}

		writer.Writer(content)
	}
}
