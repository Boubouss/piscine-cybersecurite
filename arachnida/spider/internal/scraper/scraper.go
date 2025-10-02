package scraper

import (
	"fmt"
	"io"
	"net/http"
	"spider/internal/parser"
	"spider/internal/storage"
	"spider/internal/types"
	"sync"
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


func Spider(url string, opt *types.Option) {
  pages := make(chan *types.Page, 100)
  images := make(chan *types.Image, 100)
  logs := make(chan string, 100)

  var wg sync.WaitGroup

  visited := types.Visited{ Urls: make(map[string]bool) }
  downloaded := types.Visited{ Urls: make(map[string]bool) }

  for range make([]struct{}, 5) {
    go func () {
      for current := range pages {
        if visited.IsVisited(current.URL) { 
          continue 
        }

        wg.Add(1)
        logs <- fmt.Sprintf("Fetch page %v\n", current.URL)
        
        content, err := fetchHTML(current.URL)

        if err != nil {
          continue 
        }

        err = parser.ParseHTML(content, current)

        if err != nil { 
          continue 
        }

        for _, image := range current.Images {
          if !downloaded.IsVisited(current.URL) { 
            wg.Add(1)
            images <- &types.Image{ URL: image, OPT: opt }
          }
        }

        if opt.Recursive && current.Depth < opt.MaxDepth {
          for _, link := range current.Links {
            wg.Add(1)
            pages <- &types.Page{ URL: link, DOMAIN: parser.GetDomain(link), OPT: opt, Depth: current.Depth + 1 }
          }
        }
        wg.Done()
      }
    }()
  }

  for range make([]struct{}, 5) {
    go func() {
      for current := range images {
        err := storage.SaveImage(current)

        if err != nil {
          wg.Done()
          continue
        }

        wg.Add(1)
        logs <- fmt.Sprintf("Image succefully downloaded from %v\n", current.URL)
        wg.Done()
      }
    }()
  }

  for range make([]struct{}, 5) {
    go func() {
      for log := range logs {
        fmt.Println(log)
        wg.Done()
      }
    }()
  }

  wg.Add(1)
  pages <- &types.Page{ URL: url, DOMAIN: parser.GetDomain(url), OPT: opt, Depth: 0 }

  wg.Wait()
  close(pages)
  close(images)
  close(logs)
}
