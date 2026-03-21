package ozonapi

import (
	"net/http"
)

type OzonClient struct {
	*baseClient
	ratelimiter *limiter
}

type baseClient struct {
	client  *http.Client
	options map[string]string
}

func NewClient(clientid, apikey string) *OzonClient {
	bc := baseClient{
		client: http.DefaultClient,
		options: map[string]string{
			"Client-Id": clientid,
			"Api-Key":   apikey,
		},
	}
	return &OzonClient{baseClient: &bc}
}
func (c *OzonClient) CopyToNew() *OzonClient {
	return &OzonClient{baseClient: c.baseClient}
}

// not needed if client is not rate limited. rate limit listen ctx, so calling Close() needed for unblocking rate-waiting requests
func (c *OzonClient) Close() {
	if c.ratelimiter != nil {
		c.ratelimiter.close()
	}
}
