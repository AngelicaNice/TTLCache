package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type value struct {
	mutex  sync.RWMutex
	expire int64
	value  any
}

type Cache struct {
	storage map[string]value
}

func (c Cache) Get(key string) (any, error) {
	val, ok := c.storage[key]
	if ok {
		val.mutex.Lock()
		defer val.mutex.Unlock()
		if c.isExpired(key) {
			delete(c.storage, key)
			return nil, fmt.Errorf("your value by key %s was spoiled", key)
		}
		return val.value, nil
	}
	return nil, errors.New("your value didn't exist")
}

func (c Cache) Set(key string, val any, t time.Duration) {
	c.storage[key] = value{
		expire: time.Now().Unix() + int64(t.Seconds()),
		value:  val,
	}
}

func (c Cache) Delete(key string) {
	delete(c.storage, key)
}

func (c Cache) isExpired(key string) bool {
	t := time.Now().Unix()
	if t > c.storage[key].expire {
		return true
	}
	return false
}

func New() Cache {
	return Cache{
		storage: make(map[string]value),
	}
}
