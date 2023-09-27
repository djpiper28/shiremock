package shiremock

import (
	"log"
	"net/http"
	"testing"
  "net/url"
)

func Test_MockEntityRequestMatchMethod_Fail(t *testing.T) {
	entity := MockEntity{Method: http.MethodPost}

	req := http.Request{Method: http.MethodGet}

	if entity.MatchesHttpRequest(&req) {
    log.Println("Should not have matched")
		t.Fail()
	}
}

func Test_MockEntityRequestMatchUrl_Fail(t *testing.T) {
  entity := MockEntity{Method: http.MethodGet, UrlMatcher: StringMatcher{Str:"test"}}

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
