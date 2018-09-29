package client

// APIError represents an error returned from the API.
// https://docs.royaleapi.com/#/errors
type APIError struct {
	StatusCode int    `json:"status"`  // http response code
	Message    string `json:"message"` // human readable message explaining the error
	Reason     string `json:"reason"`
}

func (err APIError) Error() string {
	return err.Message
}
