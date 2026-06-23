package pokeapi

import (
	"net/http"
	"time"

	"github.com/invaderjt/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	newCache := pokecache.NewCache(cacheInterval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: newCache,
	}
}
