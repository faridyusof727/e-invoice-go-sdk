package auth

import (
	"github.com/faridyusof727/e-invoice-go-sdk/config"
)

type AuthImpl struct {
	clientConfig  *config.Client
	serviceConfig *config.Service
}

func New(opts ...func(*AuthImpl)) Client {
	authImpl := &AuthImpl{
		serviceConfig: &config.Service{
			Environment: config.EnvironmentProduction, // default to production
		},
	}

	for _, opt := range opts {
		opt(authImpl)
	}

	return authImpl
}

// WithClientConfig returns an option function that sets the client configuration for the AuthImpl struct.
// It allows for dependency injection of the client configuration during the initialization of the AuthImpl.
//
// Parameter:
//   - clientConfig: The client configuration to be used by the AuthImplementation.
//
// Returns:
//   - A function that takes a pointer to AuthImpl and sets its clientConfig field.
func WithClientConfig(clientConfig *config.Client) func(*AuthImpl) {
	return func(i *AuthImpl) {
		i.clientConfig = clientConfig
	}
}

// WithServiceConfig returns an option function that sets the service configuration for the AuthImpl struct.
// It allows for dependency injection of the service configuration during the initialization of the AuthImpl.
//
// Parameter:
//   - serviceConfig: The service configuration to be used by the AuthImplementation.
//
// Returns:
//   - A function that takes a pointer to AuthImpl and sets its serviceConfig field.
func WithServiceConfig(serviceConfig *config.Service) func(*AuthImpl) {
	return func(i *AuthImpl) {
		i.serviceConfig = serviceConfig
	}
}
