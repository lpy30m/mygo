package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func h(n, t string) string {
	var r = len(t)
	var e = []byte(n)
	for i := 0; i < len(e); i++ {
		e[i] ^= t[(i+10)%r]
	}
	return base64.StdEncoding.EncodeToString(e)
}

// base64加密算法
func encodeBase64(message string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	return encoded
}

func main() {
	currentTime := time.Now()
	timestr := currentTime.Format("2006-01-02")
	// fmt.Println("当前日期为：", timestr)
	var t = "xyz517cda96abcd"
	page := 3
	n := 36
	str := "all"
	parce := "cn"
	iphone := "iphone"
	str2 := "@#/rank/indexPlus/brand_id/2@#"
	str3 := "@#3"
	s := rand.Float64() * 10000
	r := time.Now().UnixNano()/int64(time.Millisecond) - int64(s) - 1661224081041
	rs := int(r)
	fmt.Printf("得到的时间是%v:\n", rs)
	result := timestr + strconv.Itoa(page) + strconv.Itoa(n) + str + parce + iphone
	fmt.Println("当前要加密的结果是:", result)
	base64result := encodeBase64(result)
	fmt.Println("base64加密后的结果是:", base64result)
	urlencrypt := base64result + str2 + strconv.Itoa(rs) + str3
	fmt.Println("对url参数要加密的是", urlencrypt)
	allresult := h(urlencrypt, t)
	fmt.Println("analysis加密值是", allresult)
	client := &http.Client{}
	rawurl := "https://api.qimai.cn/rank/indexPlus/brand_id/2?analysis={}&brand=all&device=iphone&country=cn&genre=36&date=2023-04-08&page=3"
	// var encryptedAnalysis = result
	// 短的url
	// rawurl := "https://api.qimai.cn/rank/indexPlus/brand_id/1?analysis={}&brand=all&device=iphone&country=cn&genre=6014"

	url := strings.Replace(rawurl, "{}", allresult, 1)
	fmt.Println("当前请求的url是", url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "Hm_lvt_ff3eefaf44c797b33945945d0de0e370=1680675359; PHPSESSID=88btv622k0ld12v3rglge7gvg6; gr_user_id=47a1e9e0-0657-446b-bc3c-44706d07c1e4; USERINFO=p0BNY9M3IVbdFagrd6sIsv9PbXZc2yTFGbBIYdQBLVj%2Fk7GgqjczZjOBayZYLyJEHYSPObDsAWJuq%2F5IMVfEMJ79r789jFdZBVW9SFEEkCWQ%2B30Idgy9pSsWdx6WwmLOnfsXXN1dFIg0HBfkqXaNug%3D%3D; ada35577182650f1_gr_last_sent_cs1=qm18067544195; aso_ucenter=022efDXKIKVRi2OJNJVAOpP%2BsPXhs4ICgai5RFSnjmfez3qq08qxHCJFEAD3QpwTKek; AUTHKEY=GjyKT4hlp%2BAivnTuTu%2FDz%2BT7mYQb7xKlr0dxP2QrYz2A27Hn%2BBXw5ia%2BaPzOlRUCLv334lo8rDRJ7LgwpNtqVjrWFjH967o%2FGDafH0NRLtnW63at8QscQA%3D%3D; syncd=-878; Hm_lpvt_ff3eefaf44c797b33945945d0de0e370=1680699696; ada35577182650f1_gr_session_id=16a3cf47-72e2-4d26-8076-50d94940109c; ada35577182650f1_gr_last_sent_sid_with_cs1=16a3cf47-72e2-4d26-8076-50d94940109c; ada35577182650f1_gr_cs1=qm18067544195; ada35577182650f1_gr_session_id_16a3cf47-72e2-4d26-8076-50d94940109c=true; synct=1680699724.712; tgw_l7_route=1ed618a657fde25bb053596f222bc44a")
	req.Header.Set("Origin", "https://www.qimai.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://www.qimai.cn/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	str, _ = strconv.Unquote(strings.Replace(strconv.Quote(string(bodyText)), `\\u`, `\u`, -1))
	// 输出中文字符串和其长度
	newStr := strings.Replace(str, "\\", "", -1)
	fmt.Println(newStr)
	// fmt.Printf("%s\n", bodyText)

}
