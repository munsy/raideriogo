package raideriogo

import (
	"net/http"
	"testing"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	// just in case you want default correct return value
	return &http.Response{}, nil
}

func TestClient(t *testing.T) {
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// do whatever you want
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	request, _ := http.NewRequest("GET", "https://www.reallycoolurl.com/bad_request", nil)
	// as this is a test, we may skip error handling

	response, _ := client.Do(request)
	if response.StatusCode != http.StatusBadRequest {
		t.Error("invalid response status code")
	}
}

func TestGet(t *testing.T) {
	return
}
