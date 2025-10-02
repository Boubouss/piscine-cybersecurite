package parser

import (
	"fmt"
	"spider/internal/types"
	"strings"
	"golang.org/x/net/html"
)


func extractImage(node *html.Node, page *types.Page) {
  for _, attr := range node.Attr {
    if attr.Key == "src" { page.AddImage(FormatUrl(attr.Val, page.DOMAIN)) }
  }
}

func extractLink(node *html.Node, page *types.Page) {
  for _, attr := range node.Attr {
    if attr.Key == "href" { page.AddLink(FormatUrl(attr.Val, page.DOMAIN)) }
  }
}

// Cross document tree recursively to extract images and links
func parseTree(node *html.Node, page *types.Page) {
  if node == nil { return }

  if node.Type == html.ElementNode && node.Data == "img" { extractImage(node, page) } 

  if node.Type == html.ElementNode && node.Data == "a" { extractLink(node, page) }

  for child := node.FirstChild; child != nil; child = child.NextSibling {
    parseTree(child, page)
  }
}


func ParseHTML(content string, page *types.Page) (error) {
  document, err := html.Parse(strings.NewReader(content))
  
  if err != nil {
    return fmt.Errorf("Failed to parse page from %v", page.URL)
  }
  
  parseTree(document, page)

  // fmt.Printf("Images : %v\n", page.Images)
  // fmt.Printf("Liens : %v\n", page.Links)

  return nil
}
