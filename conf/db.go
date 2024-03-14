package conf

import "fmt"

type DBEnv struct {
	NAME     string // DB_NAME
	USER     string // DB_USER
	PASSWORD string // DB_PASSWORD
	HOST     string // DB_HOST
	PORT     int    // DB_PORT
}

func (x DBEnv) URL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", x.USER, x.PASSWORD, x.HOST, x.PORT, x.NAME)
}
