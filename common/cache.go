package common

import (
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
func (c *CacheCtrl) Items() map[string]gocache.Item {
	return cache.Items()
}
