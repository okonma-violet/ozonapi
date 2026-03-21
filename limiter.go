package ozonapi

import (
	"context"
	"errors"
	"time"
)

var errLimiterClosed error = errors.New("ratelimiter closed")

// rate-generation cancels with ctx, so its good to do not use main ctx
func (c *OzonClient) applyRateLimit(ctx context.Context, limit, burst int) {
	if c.ratelimiter != nil {
		panic("already ratelimited")
	}
	c.ratelimiter = newLimiter(ctx, limit, burst)
}

// func (c *OzonClient) newClientWithRateLimit(ctx context.Context, limit, burst int) *OzonClient {
// 	return &OzonClient{baseClient: c.baseClient, ratelimiter: newLimiter(ctx, limit, burst)}
// }

func default_burst(request_ratelimit int) int {
	b := request_ratelimit / 10
	if b == 0 {
		b = 1
	}
	return b
}

type limiter struct {
	tokens_ch chan struct{}
	cancel    context.CancelFunc // потенциальный deadlock
	closed    bool
}

// tokens per minute
func newLimiter(ctx context.Context, limit int, burst int) *limiter {
	if limit < 1 || burst > limit {
		panic("limit cant be less than 1, burst cant be greater than limit")
	}
	ch := make(chan struct{}, burst)
	for i := 0; i < burst; i++ {
		ch <- struct{}{}
	}
	ctxx, cancel := context.WithCancel(ctx)

	lmtr := &limiter{
		tokens_ch: ch,
		cancel:    cancel,
	}
	go lmtr.run(ctxx, limit)

	return lmtr
}

func (lmtr *limiter) run(ctx context.Context, limit int) {
	ticker := time.NewTicker(time.Minute / time.Duration(limit))
	ctx.Err()
loop:
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break loop
		default:
			select {
			case <-ctx.Done():
				ticker.Stop()
				break loop
			case <-ticker.C:
				select {
				case lmtr.tokens_ch <- struct{}{}:
					continue loop
				default:
					continue loop
				}
			}
		}
	}
	close(lmtr.tokens_ch)
}

// return false if limiter closed
func (lmtr *limiter) wait() bool {
	<-lmtr.tokens_ch
	return !lmtr.closed
}

func (lmtr *limiter) close() {
	lmtr.closed = true
	lmtr.cancel()
}

// func newLimiterStubbed() *limiter {
// 	lmtr := &limiter{
// 		cancel:    func() {},
// 		tokens_ch: make(chan struct{}),
// 	}
// 	close(lmtr.tokens_ch)
// 	return lmtr
// }
