package test

// MockResponse general response mock
type MockResponse struct {
	Response interface{}
	Err      error
}

// MockResponseWithParameters MockResponse general response mock
type MockResponseWithParameters struct {
	Params   []interface{}
	Response interface{}
	Err      error
}
