package main

import (
	"errors"
	"fmt"
	"net/http"
)

type request struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan request)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	//results["gello"] = "Hello"
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

// 웹사이트로 접속(hit)하고 그 결과를 알려주는 함수
// (*hit: 인터넷 웹 서버의 파일 1개에 접속하는 것을 뜻함)

func hitURL(url string, c chan<- request) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- request{url: url, status: status}
}
