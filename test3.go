package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
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
	for i := 0; i < 3; i++ {
		url := fmt.Sprintf("https://www.meishij.net/fenlei/jiachangcai/p%d", i-1)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cookie", "guard=69f1da56wA9QbkAPloGJLuG5mQLSznLjKw==; guardret=biWCyzWmWVtfeY+92s9A4A==")
		req.Header.Set("Pragma", "no-cache")
		req.Header.Set("Referer", "https://www.meishij.net/fenlei/jiachangcai/")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.41")
		req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		//bodyText, err := io.ReadAll(resp.Body)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("%s\n", bodyText)
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		foodName := doc.Find(".list_s2_content .list_s2_item")

		foodName.Each(func(i int, s *goquery.Selection) {
			title := s.Find(".title").Text()
			rating := s.Find(".sc").Text()
			imgSrc, exists := s.Find(".list_s2_item_img").Attr("style")
			if exists {
				url = strings.TrimPrefix(imgSrc, "background:url(")
				url = strings.TrimSuffix(url, ") center no-repeat;background-size:cover;")
				//fmt.Printf("%s", url)
			}
			fmt.Printf("%d.菜名: %s, 配料:%s,图片链接:%s\n", i+1, title, rating, url)
		})
		fmt.Println("\033[34m下一页。。。\033[0m")
	}
}
