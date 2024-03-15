package main

import (
	"flag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var reqCount = flag.Int("req-count", 1000, "")

func TestScoreBoard(t *testing.T) {
	flag.Parse()

	url := "http://localhost:3000/top_ten"

	for i := 0; i < *reqCount; i++ {
		response, err := http.Get(url)
		if err != nil {
			t.Errorf("error in request %d: %v", i+1, err.Error())
			return
		}

		statusCode := response.StatusCode

		assert.Equal(t, 200, statusCode)

		err2 := response.Body.Close()
		if err2 != nil {
			t.Errorf("error closing body of request %d: %v", i+1, err2.Error())
		}
	}
}
