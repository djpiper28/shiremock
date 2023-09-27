package shiremock

import (
	"io"
	"log"
	"net/http"
)

type Response struct {
	/// Body
	Body string
	/// Headers for the response
	Headers map[string]string
	/// Response code, see net/http for consts
	Code int
	/// Method type (GET/POST, etc...), see net/http for consts
	Method string
}

// / A request to mock
type MockEntity struct {
	/// Method type (GET/POST, etc...), see net/http for consts
	Method      string
	UrlMatcher  StringMatcher
	BodyMatcher StringMatcher
	/// What to send back to the caller
	Response Response
}

func printRequest(req *http.Request) {
	log.Printf(`Requst: '%s' '%s'
Host: %s
IP: %s
Headers: %s`, req.Method,
		req.URL,
		req.Host,
		req.RemoteAddr,
		req.Header)
}

func (entity *MockEntity) MatchesHttpRequest(req *http.Request) bool {
	printRequest(req)

	if entity.Method != req.Method {
		log.Printf("Incoming request has method '%s', expected '%s'", req.Method, entity.Method)
		return false
	}

	url := req.URL.String()
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
