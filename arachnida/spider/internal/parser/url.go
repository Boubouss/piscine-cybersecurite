package parser

import "strings"

// To handle relative path in link
func GetDomain(url string) (string) {
  
  pos := strings.Index(url, "http")
  
  if pos == 0 {
  
    pos = strings.Index(url, "://")
    
    if pos == 4 || pos == 5 {
    
      end := strings.Index(url[pos + 3:], "/")
      if end == -1 {
        end = len(url)
      }
      
      url = url[:end]
    
    }
  }

  return url
}

// Construct final url of a relative path or return url
func FormatUrl(url string, domain string) (string) {
  
  url = strings.Split(url, "?")[0]

  if strings.HasPrefix(url, "/") { return domain + url }
  return url
}

