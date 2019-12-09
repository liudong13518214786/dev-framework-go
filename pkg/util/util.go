package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func EncodeMD5(signature string) string {
	t := md5.New()
	t.Write([]byte(signature))
	return hex.EncodeToString(t.Sum(nil))
}

func SelectSql(table string, select_item []string, condtion map[string]interface{}, orderBy string, limit, offset string) string {
	sql := fmt.Sprintf("\"%s\"", strings.Join(select_item, "\",\""))
	var where string
	sql = fmt.Sprintf("select %s from %s", sql, table)
	if condtion != nil {
		where = " WHERE"
		for k, v := range condtion {
			where += fmt.Sprintf(" \"%s\"='%v' AND", k, v)
		}
		where = where[:len(where)-3]
	}

	res := sql + where
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
	sql := fmt.Sprintf("INSERT INTO %s%s VALUES(%s)", table, insertKey, insertValue[:len(insertValue)-1])
	return sql
}

func UpdateSql(table string, update_item map[string]interface{}, condtion map[string]interface{}) string {
	var updateStr string
	for k, v := range update_item {
		updateStr += fmt.Sprintf("\"%s\"='%v',", k, v)
	}
	sql := fmt.Sprintf("UPDATE %s SET %s", table, updateStr[:len(updateStr)-1])
	if condtion != nil {
		var cond string
		for k, v := range condtion {
			cond += fmt.Sprintf(" \"%s\"='%v' AND", k, v)
		}
		sql += fmt.Sprintf(" WHERE %s", cond[:len(cond)-3])
	}
	return sql

}
