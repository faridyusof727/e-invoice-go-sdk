package auth

type Client struct {
	Request *Request
	Url     string
}

type Options struct {
	Request *Request
	Url     string
}

// WithRequest returns a function that sets the Options Request field to the provided Request.
//
// r is the Request to be set in the Options.
// Returns a function that modifies the Options.
func WithRequest(r *Request) func(o *Options) {
	return func(o *Options) {
		o.Request = r
	}
}

// IsSandBox returns a function that sets the Options Url to the sandbox API if y is true.
//
// y determines whether to use the sandbox API.
// Returns a function that modifies the Options.
func IsSandBox(y bool) func(o *Options) {
	return func(o *Options) {
		if y {
			o.Url = "https://preprod-api.myinvois.hasil.gov.my"
		}
	}
}

// New creates a new Authenticator instance.
//
// It takes a variable number of functions that modify the Options.
// Returns an Authenticator instance.
func New(o ...func(o *Options)) Authenticator {
	opts := &Options{}
	opts.Url = "https://api.myinvois.hasil.gov.my"

	for _, f := range o {
		f(opts)
	}

	return &Client{
		Request: opts.Request,
		Url:     opts.Url,
	}
}
