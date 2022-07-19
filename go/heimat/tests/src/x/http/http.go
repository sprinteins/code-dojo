package http

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"time"
)

// Get configures Send for a GET HTTP Request
func Get(
	url string,
	queries []Query,
) (*http.Response, error) {
	return Send("GET", url, nil, queries)
}

// Post configures Send for a POST HTTP Request
func Post(
	url string,
	queries []Query,
	payload []byte,
) (*http.Response, error) {
	return Send("POST", url, payload, queries)
}

// Delete configures Send for a DELETE HTTP Request
func Delete(
	url string,
) (*http.Response, error) {
	return Send("DELETE", url, nil, nil)
}

// Send can send any HTTP method
func Send(
	method string,
	url string,
	payload []byte,
	queries []Query,
) (*http.Response, error) {
	tr := &http.Transport{
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: 0,
		MaxConnsPerHost:     0,
		IdleConnTimeout:     0,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: time.Minute * 2, Transport: tr}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for _, query := range queries {
		q.Add(query.key, query.value)
	}
	req.URL.RawQuery = q.Encode()

	return client.Do(req)
}

// Query _
type Query struct {
	key   string
	value string
}
