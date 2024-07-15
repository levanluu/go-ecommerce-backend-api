package setting

type Config struct {
	Mysql  MySQLSetting   `mapstructure:"mysql"`
	Logger LoggerStetting `mapstructure:"logger"`
	Redis  RedisSetting   `mapstructure:"redis"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database int    `mapstructrue:"database"`
	Password string `mapstructrue:"password"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructrue:"username"`
	Password        string `mapstructrue:"password"`
	Dbname          string `mapstructrue:"dbname"`
	MaxIdleConns    int    `mapstructrue:"maxIdleConns"`
	MaxOpenConns    int    `mapstructrue:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructrue:"connMaxLifetime"`
}

type LoggerStetting struct {
	Log_level     string `nampstructure:"log_level"`
	File_log_name string `nampstructure:"file_log_name"`
	Max_backups   int    `nampstructure:"max_backups"`
	Max_age       int    `nampstructure:"max_age"`
	Max_size      int    `nampstructure:"max_size"`
	Compress      bool   `nampstructure:"compress"`
}
