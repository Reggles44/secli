package cache

import (
	"encoding/json"
	"fmt"

	"github.com/Reggles44/secli/internal/request"
)

type Cache[T any] struct {
	URL      string
	FileName string
	Duration int
	Zip      bool

	memcache map[string]*T
}

// var cachedDir string = "/tmp/secli/"
//
// func (c *Cache[T]) path() (string, error) {
// 	url, err := url.Parse(c.URL)
// 	return path.Join(cachedDir, url.Path), err
// }

func (c *Cache[T]) Read(a ...any) (*T, error) {
	if c.memcache == nil {
		c.memcache = make(map[string]*T)
	}

	var obj *T
	url := fmt.Sprintf(c.URL, a...)
	obj, ok := c.memcache[url]
	if ok {
		return obj, nil
	}

	data, err := request.Do("GET", fmt.Sprintf(c.URL, a...))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	c.memcache[url] = obj
	return obj, err
}
