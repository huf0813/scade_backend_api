package mysql

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type envMysql struct {
	host     string
	port     string
	username string
	password string
	database string
}

func readMysqlEnv() (mysqlEnv envMysql, err error) {
	if err = godotenv.Load(); err != nil {
		return envMysql{}, err
	}
	return envMysql{
		host:     os.Getenv("MYSQL_DB_HOST"),
		port:     os.Getenv("MYSQL_DB_PORT"),
		username: os.Getenv("MYSQL_DB_USERNAME"),
		password: os.Getenv("MYSQL_DB_PASSWORD"),
		database: os.Getenv("MYSQL_DB_DATABASE"),
	}, nil
}

func NewDriverMysql() (*gorm.DB, error) {
	env, err := readMysqlEnv()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.username,
		env.password,
		env.host,
		env.port,
		env.database,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
