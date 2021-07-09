package main

import (
	"fmt"

	"github.com/badoux/goscraper"
)

func Crawl(url string) <-chan string {
	res := make(chan string)

	go func() {
		// defer close(res)
		res <- url
	}()

	return res
}

func GetImg(url *string) <-chan interface{} {
	res := make(chan interface{})
	go func() {
		s, err := goscraper.Scrape(*url, 5)
		if err != nil {
			fmt.Println(err)
			res <- nil
		}
		res <- s.Preview.Images[0]
	}()
	return res
}

func main() {
	// c1, c2 := <-Crawl("t"), <-Crawl("x")
	// fmt.Println(c1, c2)

	// urls := []string{"a", "b"}
	// c := make([]string, len(urls))
	// for k, v := range urls {
	// 	c[k] = <-Crawl(v)
	// }
	// fmt.Println(c)

	urls := []string{"https://www.yomiuri.co.jp/national/20210628-OYT1T50097/", "http://www.asahi.com/articles/ASP6X4640P6XUNHB001.html"}
	c := make([]interface{}, len(urls))
	for k, v := range urls {
		c[k] = <-GetImg(&v)
	}
	fmt.Println(c)
}
