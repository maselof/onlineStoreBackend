package constants

type Postgres struct {
	Host     string
	Port     []string
	User     string
	DBName   string
	Password string
}

var PostgresData = Postgres{
	Host:     "localhost",
	Port:     []string{"9998", "9997"},
	User:     "asafronov",
	DBName:   "online_store",
	Password: "",
}
