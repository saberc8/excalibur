package config

import "github.com/spf13/viper"

var AquilaConfig Configuration

func InitConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&AquilaConfig); err != nil {
		panic(err)
	}
}

type Configuration struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL  MySQL  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PGSQL  PGSQL  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Minio  Minio  `mapstructure:"minio" json:"minio" yaml:"minio"`
	Qiniu  Qiniu  `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
