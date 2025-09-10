package scraper

import (
  "fmt"
  "spider/internal/types"
  "spider/internal/storage"
)

func fetchHTML(url string) (string, error) {
  return url, nil
}

func parseHTML(content string, page *types.Page) {

}

func Spider(url string, option *types.Option, recursion int) {
  fmt.Printf("SPIDER : %v, %v\n", url, *option)
  page := types.Page{ URL: url }

  content, err := fetchHTML(page.URL)

  if err != nil {
    fmt.Errorf("Error %v while fetching url : %v\n", err, page.URL)
    return
  }

  parseHTML(content, &page)

  storage.SaveImages(&page.Images, option)

  if option.Recursive && recursion < option.Depth {
    for _, target := range page.Links {
      Spider(target, option, recursion + 1)
    }
  }
}
