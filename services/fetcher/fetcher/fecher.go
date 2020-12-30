package fetcher

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/temoto/robotstxt"
	"golang.org/x/net/html"
)

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func scanHTML(node *html.Node) (string, bool) {
	if isTitleElement(node) {
		return node.FirstChild.Data, true
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if result, ok := scanHTML(c); ok {
			return result, ok
		}
	}
	return "", false
}

func getTitleFromHTML(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", false
	}
	return scanHTML(doc)
}

func isAllowedByRobotstxt(urlrequest string) bool {
	pu, err := url.Parse(urlrequest)
	if err != nil {
		return false
	}
	robotstxtURL := strings.Join([]string{"https://", pu.Host, "/robots.txt"}, "")

	resp, err := http.Get(robotstxtURL)
	if err != nil {
		return true
	}
	defer resp.Body.Close()
	robotsData, err := robotstxt.FromResponse(resp)
	if err != nil {
		return false
	}
	group := robotsData.FindGroup("*")

	return group.Test(pu.Path)
}

// Fetch title form url.
func Fetch(ctx context.Context, requestURL string) (string, error) {
	if !isAllowedByRobotstxt(requestURL) {
		return requestURL, nil
	}
	resp, err := http.Get(requestURL)
	if err != nil {
		return requestURL, err
	}
	defer resp.Body.Close()

	if title, ok := getTitleFromHTML(resp.Body); ok {
		return title, nil
	}
	// titleが取得できない場合はurlを返しておきます
	return requestURL, nil
}
