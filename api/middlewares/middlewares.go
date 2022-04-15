package middlewares

import (
	"log"

	limiter "github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiterMiddleware() *limiter.Limiter {
	// Define a limit rate to 20 requests per minute.
	rate, err := limiter.NewRateFromFormatted("20-M")
	if err != nil {
		log.Fatal(err)
	}
	// Create a store with the redis client.
	store := memory.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	return limiter.New(store, rate)
}
