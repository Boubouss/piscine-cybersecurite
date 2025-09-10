package types

type Option struct {
  Recursive bool
  Depth     int
  Path      string
  Target    string
}

type Page struct {
  URL     string
  Images  []string
  Links   []string
}
