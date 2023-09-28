package shiremock

import (
	"log"
	"testing"
)

func Test_RegexMatcherInterfaceMatch(_t *testing.T) {
	var _ Matcher = &RegexMatcher{}
	log.SetFlags(log.Flags() | log.Lshortfile)
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

func Test_RegexMatcherNew_Fail(t *testing.T) {
	_, err := NewRegexMatcher("(+")
	if err == nil {
		log.Print("(+ is an invalid regex so should fail")
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
	var _ Matcher = &JsonMatcher{}
}

type PooPooTestObject struct {
	PeePee string `json:"peepee"`
}

func Test_JsonMatcherMatchesProperly(t *testing.T) {
	matcher := NewJsonMatcher[PooPooTestObject]()
	result := matcher.Matches(`{
    "peepee": "testing 123"
  }`)

	if !result {
		log.Println("Cannot match the JSON object :(")
		t.Fail()
	}
}

func Test_JsonMatcherMatches_Fail(t *testing.T) {
	matcher := NewJsonMatcher[PooPooTestObject]()
	result := matcher.Matches("poooooooo")

	if result {
		log.Println("The JSON object should not match:(")
		t.Fail()
	}
}

func Test_JsonMatcherWithAssertionsMatchesProperly(t *testing.T) {
	assertionCalled := false
	matcher := NewJsonMatcherWithAssertion[PooPooTestObject](func(obj PooPooTestObject) bool {
		log.Print("Assertion was called")
		assertionCalled = true
		return true
	})

	result := matcher.Matches(`{
    "peepee": "testing 123"
  }`)

	if !result {
		log.Println("Cannot match the JSON object :(")
		t.Fail()
	}

	if !assertionCalled {
		log.Println("Assertion func was not called")
		t.Fail()
	}
}

func Test_JsonMatcherWithAssertionsMatches_Fail(t *testing.T) {
	matcher := NewJsonMatcherWithAssertion[PooPooTestObject](func(obj PooPooTestObject) bool {
		t.Fail()
		return true
	})

	result := matcher.Matches("turds innit mate")

	if result {
		log.Println("The JSON object should not have been matched :(")
		t.Fail()
	}
}

func Test_StringMatcherMatchesTheInterface(_ *testing.T) {
	var _ Matcher = &StringMatcher{}
}

func Test_StringMatcherGood(t *testing.T) {
	matcher := StringMatcher{Str: "test"}
	if !matcher.Matches("test") {
		log.Println("test does not match test, this is quite bad")
		t.Fail()
	}
}

func Test_StringMatcherBad(t *testing.T) {
	matcher := StringMatcher{Str: "test123"}
	if matcher.Matches("test") {
		log.Println("test does match test123, this is quite bad")
		t.Fail()
	}
}

type MyObject struct {
	RequiredField string `json:"required_field" shiremock:"required"`
	OptionalField string `json:"optional_field"`
}

func Test_RequiredJsonField_Pass(t *testing.T) {
	matcher := NewJsonMatcher[MyObject]()
	result := matcher.Matches(`{"optional_field": "value", "required_field": "value"}`)

	if !result {
		log.Print("Cannot match json")
		t.Fail()
	}
}

func Test_RequiredJsonField_Fail(t *testing.T) {
	matcher := NewJsonMatcher[MyObject]()
	result := matcher.Matches(`{"optional_field": "value"}`)

	if result {
		log.Print("JSON matched when it should not have")
		t.Fail()
	}
}
