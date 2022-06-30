package jira

import (
	"encoding/base64"
	"net/http"
)

// BasicAuthRoundTripper adds basic authentication to an http.Request and delegates to the provided Next http.RoundTripper
type BasicAuthRoundTripper struct {
	Next         http.RoundTripper
	EncodedCreds string
}

// NewBasicAuthRoundTripper creates a BasicAuthRoundTripper, base64 encoding the provided username and password. Requests
// made to the Next http.RoundTripper will have a basic authentication header added to them.
func NewBasicAuthRoundTripper(next http.RoundTripper, username string, password string) *BasicAuthRoundTripper {
	return &BasicAuthRoundTripper{Next: next, EncodedCreds: basicAuth(username, password)}
}

func (b BasicAuthRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Add("Authorization", "Basic "+b.EncodedCreds)
	return b.Next.RoundTrip(request)
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
