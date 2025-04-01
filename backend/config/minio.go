package config

type Minio struct {
	Endpoint  string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKey string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	UseSSL    bool   `mapstructure:"use-ssl" json:"use-ssl" yaml:"use-ssl"`
}
