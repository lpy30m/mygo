package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://src.sjtu.edu.cn/gift/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "csrftoken=3GcaGOPCRUcj1oM6M8artR7KAExt7QP1bdNZxZWaIqANJNX5MYt9vSYfLvoCR7d5; sessionid=7r8vinlgq42bt2ptj42iiz3l69o0texr")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://src.sjtu.edu.cn/gift/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.43")
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	//	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	giftsUrl := doc.Find(".pic")
	giftsUrl.Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		// if exists {
		//	// 打印新闻标题和网址
		//	fmt.Printf("网址 %d:%s\n", i+1, link)
		// }
		baseURL := "https://src.sjtu.edu.cn"
		url := baseURL + link
		// fmt.Printf("完整网址%s\n", url)
		resp2, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp2.Body.Close()
		// bodyText, err := io.ReadAll(resp2.Body)
		// if err != nil {
		//	log.Fatal(err)
		// }
		// fmt.Printf("%s\n", bodyText)
		html := getHtml(fmt.Sprintf(url))
		parse, _ := htmlquery.Parse(strings.NewReader(html))
		names := htmlquery.Find(parse, "/html/body/div/div/div[1]/div/div")
		for _, name := range names {
			nameitem := htmlquery.Find(name, "/div[1]/div")
			giftname := htmlquery.InnerText(nameitem[1])
			priceitem := htmlquery.Find(name, "/div[2]/div")
			giftprice := htmlquery.InnerText(priceitem[1])
			numitem := htmlquery.Find(name, "/div[3]/div[2]/span/strong")
			giftnum := htmlquery.InnerText(numitem[0])
			pacenum := htmlquery.Find(name, "/div[5]/div")
			giftpace := htmlquery.InnerText(pacenum[1])
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				if r := recover(); r != nil {
					requirenum := htmlquery.Find(name, "/div[7]/div[2]/p")
					giftrequire := htmlquery.InnerText(requirenum[0])
					fmt.Printf("要求 %s", giftrequire)
				}
			}()
			requirenum := htmlquery.Find(name, "/div[6]/div[2]/p")
			giftrequire := htmlquery.InnerText(requirenum[0])
			fmt.Printf("名称:%s,价格:%s,剩余数量:%s,学校是:%s,要求：%s", giftname, giftprice, giftnum, giftpace, giftrequire)
			fmt.Println(strings.Repeat("*", 60))
			fmt.Printf("\n")
		}
		if err != nil {
			log.Fatal(err)
		}
		
	})

}

func getHtml(url_ string) string {
	req, _ := http.NewRequest("GET", url_, nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "csrftoken=3GcaGOPCRUcj1oM6M8artR7KAExt7QP1bdNZxZWaIqANJNX5MYt9vSYfLvoCR7d5; sessionid=7r8vinlgq42bt2ptj42iiz3l69o0texr")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://src.sjtu.edu.cn/gift/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.43")
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && data == nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("%s", data)
}
