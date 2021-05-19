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

func readMysqlEnv() (mysql envMysql, err error) {
	if err = godotenv.Load(); err != nil {
		return envMysql{}, err
	}
	mysql.host = os.Getenv("MYSQL_DB_HOST")
	mysql.port = os.Getenv("MYSQL_DB_PORT")
	mysql.username = os.Getenv("MYSQL_DB_USERNAME")
	mysql.password = os.Getenv("MYSQL_DB_PASSWORD")
	mysql.database = os.Getenv("MYSQL_DB_DATABASE")
	return mysql, nil
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
