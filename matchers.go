package shiremock

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"regexp"
	"strings"
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
		log.Print("Cannot create a new regex matcher ", err)
		return nil, err
	}
	ret := RegexMatcher{Regex: re}
	return &ret, nil
}

func (matcher *RegexMatcher) Matches(str string) bool {
	return matcher.Regex.MatchString(str)
}

// The object should be of the type you want the JSON to match
// This project also contains a JsonMatcherWithAssertion that allows you to run an assertion after parsing the JSON
// You can also add jsonschema:"required" to any field in the object to make it required
//
// Example:
//
//	type MyObject struct {
//		RequiredField string `json:"required_field" shiremock:"required"`
//		OptionalField string `json:"optional_field"`
//	}
func NewJsonMatcher[T any]() *JsonMatcher {
	ret := JsonMatcher{matchesFunc: func(jsonObject string) bool {
		var objectToAssert T = *new(T)

		err := json.Unmarshal([]byte(jsonObject), &objectToAssert)
		if err != nil {
			log.Print("Cannot match json ", err)
			return false
		}

		err = shiremockJsonTagValidator(objectToAssert)
		if err != nil {
			log.Print("Cannot match json ", err)
			return false
		}

		return true
	}}
	return &ret
}

type JsonMatcher struct {
	matchesFunc func(jsonObject string) bool
}

// Runs validation of the shiremock json tags
func shiremockJsonTagValidator[T any](object T) error {
	fields := reflect.ValueOf(object)
	for i := 0; i < fields.NumField(); i++ {
		tags := fields.Type().Field(i).Tag.Get("shiremock")
		if strings.Contains(tags, "required") && fields.Field(i).IsZero() {
			return errors.New("Required field is missing")
		}
	}
	return nil
}

func (matcher *JsonMatcher) Matches(str string) bool {
	return matcher.matchesFunc(str)
}

// Matches JSON then runs an assertion on it, use NewJsonMatcherWithAssertion to make an object
type JsonMatcherWithAssertion struct {
	objectToMatch any
	assertion     func(jsonObject string) bool
}

// See JsonMatcher for more docs on required fields
func NewJsonMatcherWithAssertion[T any](assertionFunc func(jsonObject T) bool) *JsonMatcherWithAssertion {
	ret := JsonMatcherWithAssertion{objectToMatch: *new(T),
		assertion: func(jsonObj string) bool {
			var objectToAssert T = *new(T)

			err := json.Unmarshal([]byte(jsonObj), &objectToAssert)
			if err != nil {
				log.Print("Cannot match json ", err)
				return false
			}

			err = shiremockJsonTagValidator(objectToAssert)
			if err != nil {
				log.Print("Cannot match json ", err)
				return false
			}

			return assertionFunc(objectToAssert)
		}}
	return &ret
}

func (matcher *JsonMatcherWithAssertion) Matches(str string) bool {
	return matcher.assertion(str)
}
