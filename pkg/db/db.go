package db

import (
	"database/sql"
	"dev-framework-go/conf"
	"fmt"
	"github.com/wonderivan/logger"
)
import _ "github.com/lib/pq"

var dbPool *sql.DB
var err error

func InitDatabasePool() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASS, conf.DB_NAME)
	dbPool, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	dbPool.SetMaxOpenConns(conf.DB_MaxOpenConns)
	dbPool.SetMaxIdleConns(conf.DB_MaxIdleConns)
	defer dbPool.Close()
	err = dbPool.Ping()
	if err != nil {
		panic(err)
	}
	logger.Debug("[INIT DB POOL SUCCESS]")
}
