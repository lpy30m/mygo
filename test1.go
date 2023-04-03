package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 设置代理
	proxyURL, err := url.Parse("socks5://127.0.0.1:1080")
	if err != nil {
		log.Fatal(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		//UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 10,
	}

	// 构造请求
	start := 25
	for i := 1; i < 10; i++ {

		url := fmt.Sprintf("https://movie.douban.com/top250?start=%d", (i-1)*start)
		//fmt.Printf(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cookie", `douban-fav-remind=1; ll="108296"; bid=Hu_rxM9jqME; _pk_id.100001.4cf6=04323b04868956ab.1666004379.1.1666005255.1666004379.; gr_user_id=e8cad792-aad0-4a1f-9a38-c187f7ae25bf; viewed="36078199"; ap_v=0,6.0`)
		req.Header.Set("Pragma", "no-cache")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-Site", "none")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.69")
		req.Header.Set("sec-ch-ua", `"Chromium";v="110", "Not A(Brand";v="24", "Microsoft Edge";v="110"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)

		// 爬取 Top 250 榜单页面
		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// 解析 HTML
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 查找电影列表
		movies := doc.Find(".grid_view .item")

		// 遍历电影列表并输出
		movies.Each(func(i int, s *goquery.Selection) {
			title := s.Find(".title").Text()
			rating := s.Find(".rating_num").Text()
			fmt.Printf("电影%d: %s, 评分%s\n", i+1, title, rating)
		})
	}
}
