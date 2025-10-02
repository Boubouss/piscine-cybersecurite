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
  counter := types.Counter{}

  for range make([]struct{}, 5) {
    wg.Go(func () {
      for current := range pages {
        if visited.IsVisited(current.URL) { 
          counter.Dec()
          continue 
        }
        
        counter.Add()
        logs <- fmt.Sprintf("Fetch page %v\n", current.URL)
        
        content, err := fetchHTML(current.URL)

        if err != nil {
          counter.Dec()
          continue 
        }

        err = parser.ParseHTML(content, current)

        if err != nil { 
          counter.Dec()
          continue 
        }

        for _, image := range current.Images {
          if !downloaded.IsVisited(current.URL) { 
            counter.Add()
            images <- &types.Image{ URL: image, OPT: opt }
          }
        }

        if opt.Recursive && current.Depth < opt.MaxDepth {
          for _, link := range current.Links {
            counter.Add()
            pages <- &types.Page{ URL: link, DOMAIN: parser.GetDomain(link), OPT: opt, Depth: current.Depth + 1 }
          }
        }
        counter.Dec()
        if counter.Get() <= 0 {
          close(pages)
          close(images)
          close(logs)
        }
      }
    })
  }

  for range make([]struct{}, 5) {
    wg.Go(func() {
      for current := range images {
        err := storage.SaveImage(current)

        if err != nil {
          counter.Dec()
          continue
        }

        counter.Add()
        logs <- fmt.Sprintf("Image succefully downloaded from %v\n", current.URL)
        counter.Dec()
      }
    })
  }

  for range make([]struct{}, 5) {
    wg.Go(func() {
      for log := range logs {
        fmt.Println(log)
        counter.Dec()
      }
    })
  }

  counter.Add()
  pages <- &types.Page{ URL: url, DOMAIN: parser.GetDomain(url), OPT: opt, Depth: 0 }

  wg.Wait()
}
