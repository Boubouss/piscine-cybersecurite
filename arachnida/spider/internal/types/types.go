package types

import (
	"slices"
	"strings"
)

type Option struct {
  Recursive bool
  Depth     int
  Path      string
}

type Page struct {
  URL     string
  DOMAIN  string  
  OPT     *Option
  Images  []string
  Links   []string
}

func (p *Page) AddLink(link string) {
  if slices.Contains(p.Links, link) || !strings.HasPrefix(link, "http") { return }
  p.Links = append(p.Links, link)
}

func (p *Page) AddImage(image string) {
  if slices.Contains(p.Images, image) { return }
  p.Images = append(p.Images, image)
}
