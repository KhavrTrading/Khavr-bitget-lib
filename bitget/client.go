package bitget

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/time/rate"
)

// BitgetClient holds the credentials and an HTTP client.
type BitgetClient struct {
	APIKey     string
	APISecret  string
	Passphrase string

	HTTPClient *http.Client
	BaseURL    string // e.g. "https://api.bitget.com"

	ipLimiter *rate.Limiter

	mu       sync.Mutex
	limiters map[string]*rate.Limiter
}

type countingTransport struct {
	base           *http.Transport
	reqCount       uint32
	resetThreshold uint32
}

func newCountingTransport(threshold uint32) *countingTransport {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
		// you can tune these if you want
		MaxIdleConns:    100,
		IdleConnTimeout: 30 * time.Second,
	}
	http2.ConfigureTransport(tr)
	return &countingTransport{
		base:           tr,
		resetThreshold: threshold,
	}
}

func (c *countingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := c.base.RoundTrip(req)
	// increment *after* the request so we get a full stream count
	if cnt := atomic.AddUint32(&c.reqCount, 1); cnt >= c.resetThreshold {
		// gracefully close the HTTP/2 session
		c.base.CloseIdleConnections()
		// reset the counter so we do this periodically
		atomic.StoreUint32(&c.reqCount, 0)
	}
	return resp, err
}

// NewBitgetClient creates a new BitgetClient.
func NewBitgetClient(apiKey, apiSecret, passphrase string) *BitgetClient {

	ct := newCountingTransport(15_000)

	return &BitgetClient{
		APIKey:     apiKey,
		APISecret:  apiSecret,
		Passphrase: passphrase,
		HTTPClient: &http.Client{
			Transport: ct,
			Timeout:   10 * time.Second,
		},
		BaseURL:   "https://api.bitget.com",
		ipLimiter: rate.NewLimiter(rate.Limit(20), 20),
		limiters:  make(map[string]*rate.Limiter),
	}
}

// SignPayload implements the Bitget HMAC SHA256 signing process for a request.
func (c *BitgetClient) SignPayload(method, requestPath, body, timestamp string) string {
	payload := timestamp + method + requestPath + body

	h := hmac.New(sha256.New, []byte(c.APISecret))
	h.Write([]byte(payload))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}

func (c *BitgetClient) GetLimiter(uid string) *rate.Limiter {
	c.mu.Lock()
	defer c.mu.Unlock()
	if l, ok := c.limiters[uid]; ok {
		return l
	}
	l := rate.NewLimiter(rate.Limit(10), 10)
	c.limiters[uid] = l
	return l
}
