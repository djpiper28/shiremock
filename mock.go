package shiremock

type MockEntry struct {
	/// The Priority of this mocking, lower priority is handled first
	Priority int
	/// Matches a request with a matching url string to this mock
	UrlMatcher StringMatcher
	/// Matches a request with a matching body to this mock
	BodyMatcher StringMatcher
}
