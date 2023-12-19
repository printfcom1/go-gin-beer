package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func InitDatabaseMariaDB() *sqlx.DB {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		panic(err.Error())
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", envMap["DB_USER"],
		envMap["DB_PASSWORD"],
		envMap["DB_HOSTNAME"],
		envMap["DB_PROT"],
		envMap["DB_NAME"])

	driver := envMap["DB_DREIVER"]
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
