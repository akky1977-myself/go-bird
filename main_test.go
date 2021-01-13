package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(Home)

	hf.ServeHTTP(recorder, req)

	status := recorder.Code

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Brird Encyclopedia is running."
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("Handler return unexpeted body: got %v wnat %v", actual, expected)
	}
}

func TestRouter(t *testing.T) {

	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/okay")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code shoud be 200, get %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err1 := ioutil.ReadAll(resp.Body)

	if err1 != nil {
		t.Fatal(err1)
	}

	respString := string(b)

	expected := "okay"

	if respString != expected {
		t.Errorf("Expected %v, but got %v", expected, respString)
	}

}

func TestRouterForNonExistentRout(t *testing.T) {
	r := newRouter()
	mockerServer := httptest.NewServer(r)
	resp, err := http.Post(mockerServer.URL+"/okay", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Respond code shoud be %d, but get %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err1 := ioutil.ReadAll(resp.Body)

	if err1 != nil {
		t.Fatal(err1)
	}

	respString := string(b)

	expected := ""

	if respString != expected {
		t.Errorf("Expected %v, but got %v", expected, respString)
	}
}

func TestFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/assets/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Respond code should be %d, but get %d", http.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if contentType != expectedContentType {
		t.Errorf("Respond content type should be  %v, but got %v.", expectedContentType, contentType)
	}
}
