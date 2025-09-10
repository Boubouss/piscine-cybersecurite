package main

import (
  "fmt"
  "flag"
  "spider/internal/types"
)


func main() {
  var opt types.SpiderOption

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

  // fmt.Printf("Download all images from '%s', if it is recursive (%v), the depth is %d and the storage path is '%s'\n", options.url, options.recursive, options.depth, options.path)
}
