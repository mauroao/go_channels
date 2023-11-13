package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://microsoft.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	channel := make(chan string)

	for _, site := range sites {
		go checkSite(site, channel)
	}

	for site := range channel {
		go func(_site string) {
			time.Sleep(5 * time.Second)
			checkSite(_site, channel)

		}(site)
	}
}

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is unrecheable! ", err)
	} else {
		fmt.Println(site, "is ok! ")
	}
	c <- site
}
