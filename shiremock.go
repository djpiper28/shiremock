package shiremock

import (
	"errors"
	"log"
	"net/http"
)

// The shiremock object that contains the state and configuration, use ShiremockBuilder to create
// an instance
type Shiremock struct {
	// Url to bind the server to, default is http://localhost:8080
	BindUrl string
	// An ordered list of the mocked endpoints to be used
	Mocks []MockEntity
}

func (s *Shiremock) Start() error {
	PrintSplashScreen()

	log.SetFlags(log.LUTC | log.Lshortfile | log.Lmsgprefix)
	log.Print("Starting Shiremock on url ", s.BindUrl)

	// Write a function to handle all http requests
	err := http.ListenAndServe(s.BindUrl, s)
	if err != nil {
		log.Print(err)
	}

	return errors.New("Shiremock server stopped")
}

func (s *Shiremock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Loop through all the mocks and find the first one that matches the request
	for _, mock := range s.Mocks {
		if mock.MatchesHttpRequest(r) {
			w.WriteHeader(mock.Response.Code)
			for key, value := range mock.Response.Headers {
				w.Header().Set(key, value)
			}
			w.Write([]byte(mock.Response.Body))
			return
		}
	}
	// If no mocks match, return a 404
	http.NotFound(w, r)
}
