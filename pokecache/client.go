package pokecache

import (
	"net/http"
	"time"
)

type Client struct {
	Cache         Cache
	EnounterCache Cache
	PokeCache     Cache
	Http          http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache:         NewCache(cacheInterval),
		EnounterCache: NewCache(cacheInterval),
		PokeCache:     NewCache(cacheInterval),
		Http: http.Client{
			Timeout: timeout,
		},
	}
}
