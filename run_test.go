package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestCurl(t *testing.T) {
	resp, err := http.Get("http://service-golang.default.svc.cluster.local:8080")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	if `Hello, World!` != sb {
		t.Errorf("fail")
	}
}
