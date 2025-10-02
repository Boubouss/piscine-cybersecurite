package types

import (
	"slices"
	"strings"
)

type Option struct {
  Recursive bool
  MaxDepth  int
  Path      string
}

type Page struct {
  URL     string
  DOMAIN  string  
  OPT     *Option
  Images  []string
  Links   []string
  Depth   int
}

type Image struct {
  URL string
  OPT *Option
}

func (p *Page) AddLink(link string) {
  if slices.Contains(p.Links, link) || !strings.HasPrefix(link, "http") { return }
  p.Links = append(p.Links, link)
}

func isGoodFormat(image string) (bool) {
  formats := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp" }
  
  for _, f := range formats {
    if strings.HasSuffix(image, f) { return true }  
  }
  return false
}

func (p *Page) AddImage(image string) {
  if slices.Contains(p.Images, image) || !isGoodFormat(image) { return }
  p.Images = append(p.Images, image)
}
