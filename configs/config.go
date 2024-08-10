package configs

type Config struct {
	Url          string
	ClientID     string
	ClientSecret string
}

func (c *Config) Prod() {
	c.Url = "https://api.myinvois.hasil.gov.my"
}

func (c *Config) SandBox() {
	c.Url = "https://preprod-api.myinvois.hasil.gov.my"
}
