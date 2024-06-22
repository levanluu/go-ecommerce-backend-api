package setting

type Config struct {
	Mysql MySQLSetting `mapstructure:"mysql"`
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
