package pkg

import "net/http"

// Accepts all client that implements HTTP Do method.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
