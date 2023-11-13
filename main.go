package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://microsoft.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkSite(site, c)
	}

	for i := 0; i < len(sites); i++ {
		fmt.Println(<-c)
	}
}

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is unrecheable! ", err)
		c <- "unrecheable! "
		return
	}
	fmt.Println(site, "is OK!")
	c <- "OK!"
}
