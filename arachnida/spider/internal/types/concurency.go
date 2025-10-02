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

type Counter struct {
  mutex   sync.Mutex
  counter int 
}

func (c *Counter) Add() {
  c.mutex.Lock()
  c.counter++
  c.mutex.Unlock()
}

func (c *Counter) Dec() {
  c.mutex.Lock()
  c.counter--
  c.mutex.Unlock()
}

func (c *Counter) Get() (int) {
  c.mutex.Lock() 
  count := c.counter
  c.mutex.Unlock()
  return count
}
