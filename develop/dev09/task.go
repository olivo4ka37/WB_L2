package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

// normalizeLink преобразовывает ссылку, добавляя к ней схему и хост, если они отсутствуют
func normalizeLink(link, scheme, host string) string {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = scheme + "://" + host + link
	}
	return link
}

// extractLinks возвращает массив ссылок из html.Token
func extractLinks(token html.Token) []string {
	var links []string

	switch token.Data {
	case "link":
		for _, attr := range token.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	case "script":
		for _, attr := range token.Attr {
			if attr.Key == "src" {
				links = append(links, attr.Val)
			}
		}
	}
	return links
}

// getLinks возвращает массив ссылок из html страницы
func getLinks(r io.Reader) []string {
	var links []string

	tokenizer := html.NewTokenizer(r)

	for {
		token := tokenizer.Next()

		switch token {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			links = append(links, extractLinks(tokenizer.Token())...)
		}
	}
}

// createFile создает файл в папке
func createFile(path, filename string, data []byte) error {
	curDir, _ := os.Getwd()
	dir := curDir + path

	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	file, err := os.Create(dir + "/" + filename)
	if err != nil {
		return err
	}

	if _, err = fmt.Fprintln(file, string(data)); err != nil {
		return err
	}

	return nil
}

// downloadLink скачивает файл по ссылке
func downloadLink(url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return fmt.Errorf("status code: %d", r.StatusCode)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = createFile(
		path.Dir(r.Request.URL.Path),
		path.Base(r.Request.URL.Path),
		data,
	); err != nil {
		return err
	}

	return nil
}

// wget скачивает страницу целиком
func wget(dir, url string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0755); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return fmt.Errorf("status code: %d", r.StatusCode)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = createFile(
		path.Dir(r.Request.URL.Path),
		"index.html",
		data,
	); err != nil {
		return err
	}

	links := getLinks(bytes.NewReader(data))
	for _, link := range links {
		normLink := normalizeLink(link, r.Request.URL.Scheme, r.Request.URL.Host)
		if err := downloadLink(normLink); err != nil {
			fmt.Printf("%s: %s\n", link, err.Error())
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: ./wget [url]")
		os.Exit(1)
	}

	err := wget("web", os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
