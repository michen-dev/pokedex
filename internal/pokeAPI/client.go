package pokeAPI


import (
	"net/http"
	"github.com/michen-dev/pokedex/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(interval),
	}
}