package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.chinanews.com.cn/scroll-news/news2.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "HMF_CI=d1176a59b64c058a6918d1c736459e57cbda476f83f1304a713f7c5b87406638cd089efc863076854ce042328ff54cbb1034e558865bde521ab27681100ddf4b4a; HBB_HC=585b38620b0e1524a3d05748e052a0e1de45349456f87931cd6461e78597ab15d919459e7b41ca5924e1a89a9b6dfabcd9")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://www.chinanews.com.cn/scroll-news/news1.html")
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

	//fmt.Printf("%s\n", bodyText)
	// 获取新闻标题和网址
	foodName := doc.Find(".content_list li")
	foodName.Each(func(i int, s *goquery.Selection) {
		kind := s.Find(".dd_lm a").Text()
		title := s.Find(".dd_bt a").Text()
		link, exists := s.Find(".dd_bt a").Attr("href")
		if exists {
			// 打印新闻标题和网址
			fmt.Printf("分类%d: %s 标题 %d: %s\n网址 %d: https://www.chinanews.com.cn/%s\n", i+1, kind, i+1, title, i+1, link)
		}
		baseURL := "https://www.chinanews.com.cn/"
		url := baseURL + link
		fmt.Printf(url)
		
	})
}
