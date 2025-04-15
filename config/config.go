package config

type Client struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

type Environment string

const (
	EnvironmentProduction Environment = "production"
	EnvironmentSandbox    Environment = "sandbox"
)

type Service struct {
	Environment Environment `json:"environment"`
}

func (s *Service) GetBaseUrl() string {
	switch s.Environment {
	case EnvironmentProduction:
		return "https://api.myinvois.hasil.gov.my"
	case EnvironmentSandbox:
		return "https://preprod-api.myinvois.hasil.gov.my"
	default:
		return "https://api.myinvois.hasil.gov.my" // Fallback to a safe default URL
	}
}