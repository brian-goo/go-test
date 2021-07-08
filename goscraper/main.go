package main

import (
	"fmt"

	"github.com/badoux/goscraper"
)

func getImg(url string) interface{} {
	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return s.Preview.Images[0]
}

func main() {
	// url := "https://www.yomiuri.co.jp/national/20210628-OYT1T50097/"
	url := "http://www.asahi.com/articles/ASP6X4640P6XUNHB001.html"

	res := make([]interface{}, 2)
	res[0] = getImg(url)

	fmt.Println(res)
	fmt.Printf("%T", res[0])

	// fmt.Printf("Icon : %s\n", s.Preview.Icon)
	// fmt.Printf("Name : %s\n", s.Preview.Name)
	// fmt.Printf("Title : %s\n", s.Preview.Title)
	// fmt.Printf("Description : %s\n", s.Preview.Description)
	// fmt.Printf("Image: %s\n", s.Preview.Images[0])
	// fmt.Printf("Url : %s\n", s.Preview.Link)
}
