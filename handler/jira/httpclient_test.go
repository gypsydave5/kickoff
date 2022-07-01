package jira

import (
	"encoding/base64"
	"github.com/gypsydave5/kickoff/test/random"
	"net/http"
	"strings"
	"testing"
)

func TestBasicAuthRoundTripper(t *testing.T) {
	spyRoundTripper := SpyRoundTripper{}
	username := random.String16()
	password := random.String16()

	bart := NewBasicAuthRoundTripper(&spyRoundTripper, username, password)

	request, err := http.NewRequest("", "", strings.NewReader(""))
	_, err = bart.RoundTrip(request)
	if err != nil {
		t.Errorf("Error in roundtripper: %s", err.Error())
	}

	expectedHeader := "Basic: " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

	header := spyRoundTripper.Request.Header.Get("Authorization")

	if header != header {
		t.Errorf("Got %v wanted %v", header, expectedHeader)
	}
}

type SpyRoundTripper struct {
	Response *http.Response
	Request  *http.Request
}

func (s *SpyRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	s.Request = request
	return nil, nil
}
