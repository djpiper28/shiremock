package shiremock

import (
	"log"
	"regexp"
	"sync"
)

type StringMatcher interface {
	Matches(str string) bool
}

type RegexMatcher struct {
	/// Compiled regex to match urls to
	Regex *regexp.Regexp
	Lock  sync.Mutex
}

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
	matcher.Lock.Lock()
	defer matcher.Lock.Unlock()
	return matcher.Regex.MatchString(str)
}
