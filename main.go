package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly/v2"
)

var link = []string{"https://cmlabs.co", "https://sequence.day", "https://go.dev"}

func main() {

	// Instantiate default collector
	c := colly.NewCollector()
	c.Async = true
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		HtmlExtractor(r.URL.String())

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err)
	})

	for _, v := range link {
		c.Visit(v)
	}

	c.Wait()

}

func HtmlExtractor(link string) {

	linkSplit := strings.Split(link, "https://")

	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	os.MkdirAll(filepath.Join(rootPath, "result"), os.ModePerm)

	fs, err := os.OpenFile(fmt.Sprintf("%s/%s.html", filepath.Join(rootPath, "result"), linkSplit[1]), os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	defer fs.Close()

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		log.Println(err)
	}

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	byt, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	fs.Write(byt)

}
