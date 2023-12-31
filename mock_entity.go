package shiremock

import (
	"io"
	"log"
	"net/http"
)

// The response for the mocked API call
type Response struct {
	// Body
	Body string
	// Headers for the response
	Headers map[string]string
	// Response code, see net/http for consts
	Code int
	// Method type (GET/POST, etc...), see net/http for consts
	Method string
}

// A request to mock, you can use NewMockEntity to help you
type MockEntity struct {
	// Method type (GET/POST, etc...), see net/http for consts
	Method      string
	UrlMatcher  Matcher
	BodyMatcher Matcher
	// What to send back to the caller
	Response Response
}

// A helper to make a new mock entity so you don't forget any of the parts lol
func NewMockEntity(method string, urlMatcher Matcher, bodymatcher Matcher, response Response) *MockEntity {
	MockEntity := MockEntity{Method: method,
		UrlMatcher:  urlMatcher,
		BodyMatcher: bodymatcher,
		Response:    response}
	return &MockEntity
}

func printRequest(req *http.Request) {
	log.Printf(`Requst: '%s' '%s'
Host: %s
IP: %s
Headers: %s`, req.Method,
		req.URL.EscapedPath(),
		req.Host,
		req.RemoteAddr,
		req.Header)
}

// Tests if a HTTP request matches this entity's definition
func (entity *MockEntity) MatchesHttpRequest(req *http.Request) bool {
	printRequest(req)

	if entity.Method != req.Method {
		log.Printf("Incoming request has method '%s', expected '%s'", req.Method, entity.Method)
		return false
	}

	url := req.URL.EscapedPath()
	if !entity.UrlMatcher.Matches(url) {
		log.Printf("Incoming request has URL '%s', expected to match via '%#v'", url, entity.UrlMatcher)
		return false
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print("Cannot read body", err)
		return false
	}

	body := string(bodyBytes)
	if !entity.BodyMatcher.Matches(body) {
		log.Printf("Incoming request has Body '%s', expected to match via '%#v'", body, entity.BodyMatcher)
		return false
	}

	return true
}
