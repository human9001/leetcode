package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

const URL = "http://vipchanger.com/res/xml/exprates.xml"

type VipChangerRates struct {
	Rates []VipChangerItem `xml:"rates"`
}
type VipChangerItem struct {
	From      string  `xml:"from"`
	To        string  `xml:"to"`
	In        float64 `xml:"in"`
	Out       float64 `xml:"out"`
	Amount    float64 `xml:"amount"`
	MinAmount int     `xml:"minAmount"`
	MaxAmount int     `xml:"maxAmount"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	request, err := http.NewRequest("GET", URL, nil)
	/*
		request.Header.Add("content-type", "text/xml")
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp")
		request.Header.Add("Accept-Language", "en-GB,en;q=0.5")
		request.Header.Add("Cache-Control", "no-cache")
		request.Header.Add("Connection", "keep-alive")
		request.Header.Add("Host", "vipchanger.com")
		request.Header.Add("DNT", "1")
		request.Header.Add("TE", "trailers")
		request.Header.Add("Pragma", "no-cache")
		request.Header.Add("Upgrade-Insecure-Requests", "1")
		request.Header.Add("Sec-Fetch-Site", "cross-site")
		request.Header.Add("Sec-Fetch-Dest", "document")
		request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:122.0) Gecko/20100101 Firefox/122.0")
	*/
	request.Header.Add("Host", "vipchanger.com")
	request.Header.Add("Cf-Connecting-Ip", " 104.248.247.153")
	request.Header.Add("Sec-Fetch-User", " ?1")
	request.Header.Add("Sec-Fetch-Site", " same-origin")
	request.Header.Add("Sec-Fetch-Mode", " navigate")
	request.Header.Add("Sec-Fetch-Dest", " iframe")
	request.Header.Add("Upgrade-Insecure-Requests", " 1")
	request.Header.Add("Dnt", " 1")
	request.Header.Add("Content-Type", "text/xml")
	request.Header.Add("Accept-Language", " en-GB,en;q=0.5")
	request.Header.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:122.0) Gecko/20100101 Firefox/122.0")
	request.Header.Add("X-Forwarded-Proto", " https")
	request.Header.Add("Content-Length", " 590")
	request.Header.Add("Cf-Ray", " 842e125c89293a49-FRA")
	request.Header.Add("X-Forwarded-For", " 104.248.247.153")
	request.Header.Add("Accept-Encoding", " utf-8")
	request.Header.Add("Cf-Ipcountry", " DE")
	request.Header.Add("Cdn-Loop", " cloudflare")
	request.Header.Add("Cookie", "__cf_bm=1ZKc.8_LrdL0A6Jcyl5tmLeLTD2R9BStcauxRVTxM6M-1704813863-1-AdFsYSe4nNF32KEa6j6zuimfyuXV21x5kCb/D9mK7OPQNqOIbtnxhMuEw102woZgLFGEUg6HhI6iG6heoXofvtM=")
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	var xmlData VipChangerRates
	fmt.Println(string(body))
	err = xml.Unmarshal(body, &xmlData)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
}
