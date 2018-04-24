package common

import (
	"fmt"
	"strings"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var cache *gocache.Cache

type CacheCtrl struct {
}

func init() {
	cache = gocache.New(1*time.Minute, 1*time.Minute)
}
func Cache() *CacheCtrl {
	return &CacheCtrl{}
}

//Log return a LogCtrl对象
func (c *BaseCtrl) Cache() *CacheCtrl {
	return &CacheCtrl{}
}
func (c *CacheCtrl) SetConfineTime(key string, val interface{}) {
	cache.Set(key, val, gocache.DefaultExpiration)
}
func (c *CacheCtrl) SetAlwaysTime(key string, val interface{}) {
	cache.Set(key, val, gocache.NoExpiration)
}
func (c *CacheCtrl) Get(key string) (interface{}, bool) {
	val, found := cache.Get(key)
	return val, found
}
func (c *CacheCtrl) Del(key string) {
	cache.Delete(key)
}
func (c *CacheCtrl) ScanDel(key string) {
	clearCache := cache.Items()
	for k, _ := range clearCache {
		if strings.Contains(k, key) {
			cache.Delete(k)
		}
	}
}
func (c *CacheCtrl) ScanKey(key string) {
	clearCache := cache.Items()
	for k, _ := range clearCache {
		if key == "" {
			fmt.Println(k)
		} else {
			if strings.Contains(k, key) {
				fmt.Println(k)
			}
		}
	}
}
func (c *CacheCtrl) ScanValue(key string) {
	clearCache := cache.Items()
	for k, v := range clearCache {
		if key == "" {
			fmt.Println(v)
		} else {
			if strings.Contains(k, key) {
				fmt.Println(v)
			}
		}
	}
}
func (c *CacheCtrl) Items() map[string]gocache.Item {
	return cache.Items()
}
