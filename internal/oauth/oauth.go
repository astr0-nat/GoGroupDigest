package oauth

// Local web server to receive authorization code

type authorizationCode struct {
	// Authorization code.
	Code string
	// State of request.
	State string
}

type codeReceiver struct {
	// Channel to receive an authorization code.
	Result chan *authorizationCode
	//Channel to receive an error message
	Error chan error
}

// newCodeReceiver creates a set of channels
func newCodeReceiver() *codeReceiver {
	return &codeReceiver{
		Result: make(chan *authorizationCode, 1),
		Error:  make(chan error, 1),
	}
}
