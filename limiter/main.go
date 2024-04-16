package limiter

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

type IPRateLimiterI struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	s   int
}

func NewIPRateLimiter(r rate.Limit, s int) *IPRateLimiterI {
	return &IPRateLimiterI{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		s:   s,
	}
}

func (i *IPRateLimiterI) AddIPtoMap(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.s)
	i.ips[ip] = limiter
	return limiter
}

func (i *IPRateLimiterI) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()

	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIPtoMap(ip)
	}

	i.mu.Unlock()
	return limiter
}

func LoadEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ENV_LOAD_FAILED, %s", err.Error())
	}
	return os.Getenv(key)
}
