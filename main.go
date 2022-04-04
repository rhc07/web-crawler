package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

var Links []string

func main() {
	uri, _ := GetURI()
	_, err := GetAnchorLinks(uri)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	for _, v := range Links {
		fmt.Println(v)
	}
}

func GetURI() (string, error) {
	fmt.Println("Enter a valid url: ")
	var uri string
	_, err := fmt.Scanln(&uri)
	if err != nil {
		return "", fmt.Errorf("error: cannot scan user input")
	}
	return uri, nil
}

func GetAnchorLinks(uri string) ([]string, error) {
	log.Printf("fetching: %s", uri)
	response, err := http.Get(uri)
	if err != nil {
		return []string{}, errors.New("no body in the response")
	}
	body := response.Body
	z := html.NewTokenizer(body)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return Links, nil
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						Links = append(Links, attr.Val)
					}
				}
			}
		}
	}
}
