package main

import (
	"fmt"
	"time"

	"github.com/AngelicaNice/TTLCache/cache"
)

func main() {
	cache := cache.New()
	cache.Set("key_int", 4, time.Second)
	cache.Set("key_string", "some_string", time.Minute)
	for i := 1; i < 10; i++ {
		func() {
			go fmt.Println(cache.Get("key_int"))
			cache.Set("key_int", 5, time.Second)
		}()
	}
	time.Sleep(time.Second)
	for i := 1; i < 10; i++ {
		func() {
			fmt.Println(cache.Get("key_string"))
		}()
	}
	time.Sleep(time.Second * 10)
}
