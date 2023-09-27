package shiremock

import (
	"encoding/json"
	"log"
	"regexp"
)

type Matcher interface {
	Matches(str string) bool
}

type StringMatcher struct {
	Str string
}

func (matcher *StringMatcher) Matches(str string) bool {
	return str == matcher.Str
}

// You can use the NewRegexMatcher function to do the regex funkiness for you
type RegexMatcher struct {
	// Compiled regex to match urls to
	Regex *regexp.Regexp
}

// Compiles the regex and returns a regex matcher
func NewRegexMatcher(regex string) (*RegexMatcher, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		log.Print("Cannot create a new regex matcher", err)
		return nil, err
	}
	ret := RegexMatcher{Regex: re}
	return &ret, nil
}

func (matcher *RegexMatcher) Matches(str string) bool {
	return matcher.Regex.MatchString(str)
}

type JsonMatcher struct {
	ObjectToMatch any
}

func (matcher *JsonMatcher) Matches(str string) bool {
	err := json.Unmarshal([]byte(str), &matcher.ObjectToMatch)
	if err != nil {
		log.Print("Cannot match json", err)
	}
	return err == nil
}

// Matches JSON then runs an assertion on it, use NewJsonMatcherWithAssertion to make an object
type JsonMatcherWithAssertion struct {
	objectToMatch any
	assertion     func(jsonObject string) bool
}

func NewJsonMatcherWithAssertion[T any](assertionFunc func(jsonObject T) bool) *JsonMatcherWithAssertion {
	ret := JsonMatcherWithAssertion{objectToMatch: *new(T),
		assertion: func(jsonObj string) bool {
			var objectToAssert T = *new(T)

			err := json.Unmarshal([]byte(jsonObj), &objectToAssert)
			if err != nil {
				log.Print("Cannot match json", err)
				return false
			}

			return assertionFunc(objectToAssert)
		}}
	return &ret
}

func (matcher *JsonMatcherWithAssertion) Matches(str string) bool {
	return matcher.assertion(str)
}
