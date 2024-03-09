package constants

type Postgres struct {
	Host     string
	Port     []string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

var PostgresData = Postgres{
	Host:     "localhost",
	Port:     []string{"9998", "9997"},
	User:     "postgres",
	DBName:   "online_store",
	Password: "postgres",
	SSLMode:  "disable",
}
