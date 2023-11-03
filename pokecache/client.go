package pokecache

import (
	"net/http"
	"time"
)

type Client struct {
	Cache Cache
	Http  http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: NewCache(cacheInterval),
		Http: http.Client{
			Timeout: timeout,
		},
	}
}
