package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []Website {
		Google,
		Golang,
		Stackoverflow,
		Amazon,
		Facebook,
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(c, link.String())
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5*time.Second)
			checkLink(c, l)
		}(l) 
	}
}

func checkLink(c chan string, link string) {
	_, err := http.Get(link)
	if(err != nil) {
		fmt.Printf("%s might be down.\n", link)
		c <- link
		return
	}
	fmt.Printf("%s is up.\n", link)
	c <- link 
}