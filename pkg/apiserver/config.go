package apiserver


type Config struct {
	HttpConfig
}

type HttpConfig struct {
	BindAddr     string `yaml:"bindAddr,omitempty"`
	InsecurePort int    `yaml:"insecurePort,omitempty"`
	SecurePort   int    `yaml:"securePort, omitempty"`
	GenericPort  int    `yaml:"genericPort,omitempty"`
	TlsCert      string `yaml:"tlsCert,omitempty"`
	TlsKey       string `yaml:"tlsKey,omitempty"`
	CaCert       string `yaml:"caCert,omitempty"`
	CaKey        string `yaml:"caKey,omitempty"`
}

func (c *Config) Validate() []error {
	return nil
}
