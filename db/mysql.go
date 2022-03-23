package db

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type config struct {
	User, Password, Port, Host string
}

var Config = &config{}

func init() {
	flag.StringVar(&Config.User, "u", "root", "user")
	flag.StringVar(&Config.Password, "p", "root", "password")
	flag.StringVar(&Config.Host, "Host", "127.0.0.1", "host")
	flag.StringVar(&Config.Port, "Port", "3307", "port")
}

func Open(max int) *sql.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Config.User, Config.Password, Config.Host, Config.Port, "db")
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open db err: %v\n", err)
		return nil
	}
	return DB
}
