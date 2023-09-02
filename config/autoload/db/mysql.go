package db

type Mysql struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`                            // 服务器地址:端口
	Port        string `mapstructure:"port" json:"port" yaml:"port"`                            //:端口
	Config      string `mapstructure:"config" json:"config" yaml:"config"`                      // 高级配置
	Dbname      string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                   // 数据库名
	Username    string `mapstructure:"username" json:"username" yaml:"username"`                // 数据库用户名
	Password    string `mapstructure:"password" json:"password" yaml:"password"`                // 数据库密码
	Prefix      string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                      // 全局表前缀，单独定义TableName则不生效
	Singular    bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                // 是否开启全局禁用复数，true表示开启
	Engine      string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`     // 数据库引擎，默认InnoDB
	MaxIdleConn int    `mapstructure:"max-idle-conn" json:"max-idle-conn" yaml:"max-idle-conn"` // 空闲中的最大连接数
	MaxOpenConn int    `mapstructure:"max-open-conn" json:"max-open-conn" yaml:"max-open-conn"` // 打开到数据库的最大连接数
	LogMode     string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                // 是否开启Gorm全局日志
	//LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
