package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/health")
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	// _ := string(body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
