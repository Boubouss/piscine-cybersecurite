package types

import (
	"sync"
)

type Visited struct {
  mutex sync.Mutex
  Urls  map[string]bool
}

func (v *Visited) IsVisited(url string) (bool) {
  v.mutex.Lock()
  if v.Urls[url] { 
    v.mutex.Unlock()
    return true 
  }
  v.Urls[url] = true;
  v.mutex.Unlock()
  return false
}

