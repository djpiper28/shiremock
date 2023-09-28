package shiremock

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

type ExampleType struct {
	RequiredField string `json:"required_field" shiremock:"required"`
	OptionalField string `json:"optional_field"`
}

const URL = "localhost:65421"

// This starts the server with an example configuration
func StartServer() {
	regex, err := NewRegexMatcher("/test/[0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	shiremock := Shiremock{
		BindUrl: URL,
		Mocks: []MockEntity{
			{
				Method: http.MethodGet,
				UrlMatcher: &StringMatcher{
					Str: "/test",
				},
				BodyMatcher: &StringMatcher{
					Str: "",
				},
				Response: Response{
					Body: "Hello GoLinuxCloud Members!",
					Code: http.StatusOK,
				},
			},
			{
				Method:      http.MethodPost,
				UrlMatcher:  regex,
				BodyMatcher: NewJsonMatcher[ExampleType](),
				Response: Response{
					Body: "Hello GoLinuxCloud Members!",
					Code: http.StatusCreated,
				},
			},
			{
				Method: http.MethodPost,
				UrlMatcher: &StringMatcher{
					Str: "/other_test",
				},
				BodyMatcher: NewJsonMatcherWithAssertion[ExampleType](func(object ExampleType) bool {
					return object.RequiredField == "test"
				}),
				Response: Response{
					Body: "Hello GoLinuxCloud Members!",
					Code: http.StatusCreated,
				},
			},
		},
	}

	err = shiremock.Start()
	if err != nil {
		log.Println("Error starting server", err)
	}
}

// This runs the server then system tests it
// Yes I used Github Copilot to write these, don't even @ me
func Test_Example(t *testing.T) {
	go StartServer()

	// Wait for 1 seconds
	time.Sleep(1 * time.Second)

	URL := "http://" + URL

	// Test the first
	response, err := http.Get(URL + "/test")
	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != "Hello GoLinuxCloud Members!" {
		t.Errorf("Expected body 'Hello GoLinuxCloud Members!', got '%s'", string(bodyBytes))
	}

	// Test the second one
	response, err = http.Post(URL+"/test/12345", "application/json", strings.NewReader(`{"required_field": "test"}`))

	if response.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", response.StatusCode)
	}

	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != "Hello GoLinuxCloud Members!" {
		t.Errorf("Expected body 'Hello GoLinuxCloud Members!', got '%s'", string(bodyBytes))
	}

	// Test the second one (Error)
	response, err = http.Post(URL+"/test/12345", "application/json", strings.NewReader("doggy input"))

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404 , got %d", response.StatusCode)
	}

	// Test the third one
	response, err = http.Post(URL+"/other_test", "application/json", strings.NewReader(`{"required_field": "test"}`))

	if response.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", response.StatusCode)
	}

	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != "Hello GoLinuxCloud Members!" {
		t.Errorf("Expected body 'Hello GoLinuxCloud Members!', got '%s'", string(bodyBytes))
	}
}
