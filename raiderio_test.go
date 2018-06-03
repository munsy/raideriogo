package raideriogo

import (
	"net/http"
	"testing"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockHTTPClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

// FIX THIS IT IS WORTHLESS FILTH
func TestGet(t *testing.T) {
	client := &MockHTTPClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	request := http.Get(EndpointCharacter("us", "thrall", "munsy", ""))

	response, _ := client.Do(request)
	if response.StatusCode != http.StatusBadRequest {
		t.Error("Invalid response status code")
	}
}
