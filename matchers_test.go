package shiremock

import (
	"log"
	"testing"
)

func Test_RegexMatcherInterfaceMatch(_t *testing.T) {
	var _ StringMatcher = &RegexMatcher{}
}

func Test_RegexMatcherNew(t *testing.T) {
	matcher, err := NewRegexMatcher(".*")
	if err != nil {
		log.Println("Cannot create a new regex", err)
		t.Fail()
	}

	if matcher == nil {
		log.Println("Returned a nil instance")
		t.Fail()
	}
}

func Test_RegexMatcherMatchesRegexWithObjectReuse(t *testing.T) {
	matcher, err := NewRegexMatcher("a{5}b+")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	for i := 0; i < 100; i++ {
		result := matcher.Matches("aaaaabbbbbb")
		if !result {
			log.Println("The string should have matched properly")
			t.Fail()
		}
	}
}

func Test_RegexMatcherMatchesRegex(t *testing.T) {
	matcher, err := NewRegexMatcher("a{5}b+")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	result := matcher.Matches("aaaaabbbbbb")
	if !result {
		log.Println("The string should have matched properly")
		t.Fail()
	}
}

func Test_JsonMatcherDoesTheInterfaceThing(_ *testing.T) {
	var _ StringMatcher = &JsonMatcher{}
}

type PooPooTestObject struct {
	PeePee string `json:"peepee"`
}

func Test_JsonMatcherMatchesProperly(t *testing.T) {
	matcher := JsonMatcher{}
	result := matcher.Matches(`{
    "peepee": "testing 123"
  }`)

	if !result {
		log.Println("Cannot match the JSON object :(")
		t.Fail()
	}
}
