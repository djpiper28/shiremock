package shiremock

import (
	"log"
	"testing"
	"time"
)

func Test_ServerStart_Fail(t *testing.T) {
	s := Shiremock{BindUrl: "localhost:9981",
		Mocks: []MockEntity{},
	}

	go s.Start()
	time.Sleep(1)

	var err error = nil
	go func() {
		err = s.Start()
	}()

	time.Sleep(1)
	if err == nil {
		log.Print("Test should have failed due to binding on the same port twice")
		t.Fail()
	}
}
