package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestCurl(t *testing.T) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
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
	if `{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}` != sb {
		t.Errorf("fail")
	}
}
