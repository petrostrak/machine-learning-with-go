package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {

	// Create a cache with a default expiration time of 5 mins
	// and which purges expired items every 30 seconds
	c := cache.New(5*time.Minute, 30*time.Second)

	// Put a key and value into the cache.
	c.Set("myKey", "myValue", cache.DefaultExpiration)

	// To retrieve the value for myKey out of cache, we just need
	// to use the Get method.
	v, found := c.Get("myKey!")
	if !found {
		fmt.Println("Did not find key with given value.")
	} else {
		fmt.Printf("key: mykey, value: %s\n", v)
	}
}
