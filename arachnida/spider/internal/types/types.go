package types

type SpiderOption struct {
  Recursive bool
  Depth     int
  Path      string
  Target    string
}

type SpiderData struct {
  Target  string
  Images  []string
  Urls    []string
}

