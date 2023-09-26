package shiremock

import (
	"log"
	"testing"
)

func Test_RegexMatcherInterfaceMatch(_t *testing.T) {
	var _ StringMatcher
	_ = &RegexMatcher{}
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
