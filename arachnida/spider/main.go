package main

import (
  "fmt"
  "flag"
  "spider/internal/types"
  "spider/internal/scraper"
)


func main() {
  var opt types.Option

  flag.BoolVar(&opt.Recursive, "r", false, "Recursively downloads the images in a URL received as a parameter")
  flag.IntVar(&opt.Depth, "l", 5, "Maximum depth level of the recursive download")
  flag.StringVar(&opt.Path, "p", "./data/", "Path where the downloaded files will be saved")
  flag.Parse()
  
  args := flag.Args()
  
  if len(args) != 1 {
    fmt.Println("Wrong number of arguments")
    return
  }

  opt.Target = args[0]

  scraper.Spider(args[0], &opt, 0)
}
