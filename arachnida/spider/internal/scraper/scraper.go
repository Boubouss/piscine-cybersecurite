package scraper

import (
  "fmt"
  "spider/internal/types"
  "spider/internal/storage"
  "net/http"
  "io"
)

func fetchHTML(url string) (string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", fmt.Errorf("Error while fetching %v", url)
  }

  defer resp.Body.Close()
  
  if resp.StatusCode != http.StatusOK {
    return "", fmt.Errorf("Status code => %v", resp.StatusCode)
  }
  
  body, err := io.ReadAll(resp.Body)
  
  if err != nil {
    return "", err
  }
  
  return string(body), nil
}

func parseHTML(content string, page *types.Page) {

}

func Spider(url string, option *types.Option, recursion int) {
  fmt.Printf("SPIDER : %v, %v\n", url, *option)
  page := types.Page{ URL: url }

  content, err := fetchHTML(page.URL)

  if err != nil {
    fmt.Printf("Error : %v\n", err)
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
