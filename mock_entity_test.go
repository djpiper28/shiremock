package shiremock

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func Test_MockEntityRequestMatchMethod_Fail(t *testing.T) {
	entity := MockEntity{Method: http.MethodPost}

	url, err := url.Parse("http://test.com/testing123")
	if err != nil {
		log.Print("Cannot make mock url", err)
		t.Fail()
	}

	req := http.Request{Method: http.MethodGet, URL: url}

	if entity.MatchesHttpRequest(&req) {
		log.Println("Should not have matched")
		t.Fail()
	}
}

func Test_MockEntityRequestMatchUrl_Fail(t *testing.T) {
	entity := MockEntity{Method: http.MethodGet, UrlMatcher: &StringMatcher{Str: "/test"}}

	url, err := url.Parse("http://test.com/testing123")
	if err != nil {
		log.Print("Cannot make mock url", err)
		t.Fail()
	}

	req := http.Request{Method: http.MethodGet, URL: url}

	if entity.MatchesHttpRequest(&req) {
		log.Println("Should not have matched")
		t.Fail()
	}
}

func Test_MockEntityRequstMatchBody_Fail(t *testing.T) {
	entity := MockEntity{Method: http.MethodGet, UrlMatcher: &StringMatcher{Str: "/test"}, BodyMatcher: &StringMatcher{Str: "test"}}

	url, err := url.Parse("http://test.com/testing123")
	if err != nil {
		log.Print("Cannot make mock url", err)
		t.Fail()
	}

	body := io.NopCloser(strings.NewReader("Hello GoLinuxCloud Members!"))

	req := http.Request{Method: http.MethodGet, URL: url, Body: body}

	if entity.MatchesHttpRequest(&req) {
		log.Println("Should not have matched")
		t.Fail()
	}
}

func Test_MockEntityRequstMatch_Pass(t *testing.T) {
	entity := MockEntity{Method: http.MethodGet, UrlMatcher: &StringMatcher{Str: "/test"}, BodyMatcher: &StringMatcher{Str: "test"}}

	url, err := url.Parse("http://test.com/test")
	if err != nil {
		log.Print("Cannot make mock url", err)
		t.Fail()
	}

	body := io.NopCloser(strings.NewReader("test"))

	req := http.Request{Method: http.MethodGet, URL: url, Body: body}

	if !entity.MatchesHttpRequest(&req) {
		log.Println("Should have matched")
		t.Fail()
	}
}
