package scraper

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// ✅ Extracts links from a webpage
func extractLinks(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						fmt.Println("Found link:", attr.Val)
					}
				}
			}
		}
	}
}

// ✅ Run Web Scraper
func Run() {
	fmt.Println("Scraping example.com for links...")
	extractLinks("https://example.com")
}

/*
🔹 Explanation:
- `html.NewTokenizer(resp.Body)`: Parses HTML.
- `token.Data == "a"`: Extracts anchor (`<a>` tags).
- `attr.Val`: Retrieves the `href` attribute.
*/
