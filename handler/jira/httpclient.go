package jira

import (
	"encoding/base64"
	"net/http"
)

type BasicAuthRoundTripper struct {
	next         http.RoundTripper
	encodedCreds string
}

func NewBasicAuthRoundTripper(next http.RoundTripper, username string, password string) *BasicAuthRoundTripper {
	return &BasicAuthRoundTripper{next: next, encodedCreds: basicAuth(username, password)}
}

func (b BasicAuthRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Add("Authorization", "Basic "+b.encodedCreds)
	return b.next.RoundTrip(request)
}

// NewBasicAuthHTTP is a convenience function that creates a default *http.Client with the BasicAuthRoundTripper applied
// to it. See BasicAuthRoundTripper for details.
func NewBasicAuthHTTP(username string, password string) *http.Client {
	c := &http.Client{
		Transport: NewBasicAuthRoundTripper(http.DefaultTransport, username, password),
	}

	return c
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
