package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
)

func main() {
	uri := GetURI()
	body := GetResponseBody(uri)
	for _, v := range GetAnchorLinks(body) {
		fmt.Println(v)
	}
}

func GetURI() string {
	fmt.Println("Enter a valid url: ")
	var uri string
	fmt.Scanln(&uri)
	return uri
}

func GetResponseBody(uri string) io.Reader {
	log.Printf("fetching: %s", uri)
	response, err := http.Get(uri)
	if err != nil {
		log.Fatalf("error retrieving site `%s`: %s", uri, err)
	}
	return response.Body
}

func GetAnchorLinks(body io.Reader) []string {
	var links []string

	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
