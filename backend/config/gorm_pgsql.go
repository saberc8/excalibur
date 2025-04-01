package config

type PGSQL struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 服务器地址:端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               //:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true表示开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // 数据库引擎，默认InnoDB
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log_zap" json:"log_zap" yaml:"log_zap"`                      // 是否通过zap写入日志文件
}

func (p *PGSQL) DSN() string {
	return "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}

func (p *PGSQL) LinkDSN(dbname string) string {
	return "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

func (m *PGSQL) GetLogMode() string {
	return m.LogMode
}
