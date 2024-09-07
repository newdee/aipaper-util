package database

// DatabaseConfig 数据库配置
type DatabaseConfig interface {
	GetHost() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDataBase() string
	GetCharset() string
}

// DBConfig 数据库配置信息
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

func (c *DBConfig) GetHost() string {
	if c != nil {
		return c.Host
	}
	return ""
}

func (c *DBConfig) GetPort() int {
	if c != nil {
		return c.Port
	}
	return 0
}

func (c *DBConfig) GetUser() string {
	if c != nil {
		return c.User
	}
	return ""
}

func (c *DBConfig) GetPassword() string {
	if c != nil {
		return c.Password
	}
	return ""
}

func (c *DBConfig) GetDataBase() string {
	if c != nil {
		return c.Database
	}
	return ""
}

func (c *DBConfig) GetCharset() string {
	if c != nil {
		return c.Charset
	}
	return ""
}
