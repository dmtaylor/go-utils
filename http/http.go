package http

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

const DefaultRetryLimit = uint(5)

var MissingRetryAfter = errors.New("missing retry-after header in request")

type RetryClient struct {
	C          http.Client
	RetryLimit uint
}

// Do executes HTTP request, implementing retries
func (c RetryClient) Do(req *http.Request) (*http.Response, error) {

	var res *http.Response
	var err error
	for i := uint(0); i < c.RetryLimit+1; i++ {
		res, err = c.C.Do(req)
		if err != nil {
			return res, err
		}
		if res.StatusCode == http.StatusTooManyRequests {
			ws := res.Header.Get("Retry-After")
			if ws == "" {
				return nil, MissingRetryAfter
			}
			if waitSecs, err := strconv.Atoi(ws); err == nil {
				time.Sleep(time.Duration(waitSecs) * time.Second)
				continue
			}
			if waitUntil, err := time.Parse(time.RFC1123, ws); err == nil {
				waitDuration := time.Until(waitUntil)
				time.Sleep(waitDuration)
				continue
			}
		}
		return res, err
	}

	return res, err
}
