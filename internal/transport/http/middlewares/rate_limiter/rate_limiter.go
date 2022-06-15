package rate_limiter

import (
	"github.com/ulule/limiter/v3"
)

func New(store limiter.Store) *limiter.Limiter {
	rate, _ := limiter.NewRateFromFormatted("15-M")

	return limiter.New(store, rate)
}
