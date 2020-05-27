package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.crunchbase.com/organization/zoom-video-communications#section-overview"

	resp, err := http.Get(url) //①HTTPリクエスト行を入れる
	if err != nil {
	}
	defer resp.Body.Close()                //③万が一のエラーのためのコードを入れる
	body, err := ioutil.ReadAll(resp.Body) //②ReadAll入れる

	println(string(body)) //④HTMLをStringで取得するコードを入れる

}
