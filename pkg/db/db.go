package db

import (
	"dev-framework-go/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"strings"
)
import _ "github.com/lib/pq"

var DBPool *gorm.DB
var err error

func InitDatabasePool() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASS, conf.DB_NAME)
	DBPool, err = gorm.Open("postgres", psqlInfo)
	//DBPool, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	DBPool.DB().SetMaxIdleConns(conf.DB_MaxIdleConns)
	DBPool.DB().SetMaxOpenConns(conf.DB_MaxOpenConns)
	//DBPool.SetMaxOpenConns(conf.DB_MaxOpenConns)
	//DBPool.SetMaxIdleConns(conf.DB_MaxIdleConns)
	//defer dbPool.Close()
	err = DBPool.DB().Ping()
	if err != nil {
		panic(err)
	}
	logger.Debug("[INIT DB POOL SUCCESS]")
}

func SelectSql(table string, select_item []string, condtion map[string]interface{}, orderBy string, limit, offset string) string {
	sqlstring := ""
	if select_item == nil {
		sqlstring = "*"
	} else {
		sqlstring = fmt.Sprintf("\"%s\"", strings.Join(select_item, "\",\""))
	}
	var where string
	sqlstring = fmt.Sprintf("select %s from %s", sqlstring, table)
	if condtion != nil {
		where = " WHERE"
		for k, v := range condtion {
			where += fmt.Sprintf(" \"%s\"='%v' AND", k, v)
		}
		where = where[:len(where)-3]
	}

	res := sqlstring + where
	if orderBy != "" {
		res += orderBy
	}
	if limit != "" {
		res += fmt.Sprintf(" LIMIT %s", limit)
	}
	if offset != "" {
		res += fmt.Sprintf(" OFFSET %s", offset)
	}
	return res
}

func InsertSql(table string, insert_item map[string]interface{}) string {
	var insertKey string
	var insertValue string
	for k, v := range insert_item {
		insertKey += fmt.Sprintf("\"%s\",", k)
		insertValue += fmt.Sprintf("'%v',", v)
	}
	insertKey = fmt.Sprintf("(%s)", insertKey[:len(insertKey)-1])
	fmt.Println(insertKey)
	fmt.Println(insertValue)
	sqlstring := fmt.Sprintf("INSERT INTO %s%s VALUES(%s)", table, insertKey, insertValue[:len(insertValue)-1])
	return sqlstring
}

func UpdateSql(table string, update_item map[string]interface{}, condtion map[string]interface{}) string {
	var updateStr string
	for k, v := range update_item {
		updateStr += fmt.Sprintf("\"%s\"='%v',", k, v)
	}
	sqlstring := fmt.Sprintf("UPDATE %s SET %s", table, updateStr[:len(updateStr)-1])
	if condtion != nil {
		var cond string
		for k, v := range condtion {
			cond += fmt.Sprintf(" \"%s\"='%v' AND", k, v)
		}
		sqlstring += fmt.Sprintf(" WHERE %s", cond[:len(cond)-3])
	}
	return sqlstring

}
