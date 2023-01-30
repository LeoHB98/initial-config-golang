package config

type Email struct {
	From string
	SMTP string
	Port string
	Pass string
	Tls  string
	TO   string
}

type DBConfig struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort int
	DbSid  string
}

type Config struct {
	DbConfig           *DBConfig
	Authorization      string
	EmailLog           *Email
	Email              *Email
	FTP                *FTP
	ParamSigcomIdEmail string
	Port               int
	Company            string
	DevMode            bool
}

type FTP struct {
	FtpHost string
	FtpPort string
	FtpUser string
	FtpPass string
}
