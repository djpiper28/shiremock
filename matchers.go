package shiremock

import (
	"encoding/json"
	"log"
	"regexp"
	"sync"
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

// / You can use the NewRegexMatcher function to do the regex funkiness for you
type RegexMatcher struct {
	/// Compiled regex to match urls to
	Regex *regexp.Regexp
	lock  sync.Mutex
}

// / Compiles the regex and returns a regex matcher
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
	matcher.lock.Lock()
	defer matcher.lock.Unlock()
	return matcher.Regex.MatchString(str)
}

type JsonMatcher struct {
	ObjectToMatch any
}

func (matcher *JsonMatcher) Matches(str string) bool {
	err := json.Unmarshal([]byte(str), &matcher.ObjectToMatch)
	return err == nil
}
