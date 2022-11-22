package httpclient

import (
	"errors"
	"net/http"
	"sync"
)

type mock struct {
	resp *http.Response

	err error

	locker sync.Mutex
}

// New Initialize a new httpclient mock
func New() *mock {
	return &mock{}
}

// SetResponse set response for the mock
func (h *mock) SetResponse(resp *http.Response) *mock {
	h.resp = resp

	return h
}

// SetError set custom error for the mock
func (h *mock) SetError(err error) *mock {
	h.err = err

	return h
}

// Get return result based on the given settings
func (h *mock) Get(_ string) (*http.Response, error) {
	if h.err != nil {
		return nil, h.err
	}

	if h.resp == nil {
		return nil, errors.New("mock_response_not_set")
	}

	return h.resp, nil
}
