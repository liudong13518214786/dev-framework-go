package models

import (
	"database/sql"
	"dev-framework-go/pkg/db"
	"fmt"
)

//测试数据库
type Project struct {
	Id        sql.NullInt64
	Uuid      sql.NullString
	PjName    sql.NullString
	ShowName  sql.NullString
	Descs     sql.NullString
	BuildTime sql.NullString
	Status    sql.NullString
}

//查询单条
func SelectTags(sqlstring string, return_field []string) *Project {
	var p Project
	fmt.Println(sqlstring)
	rows := db.DBPool.QueryRow(sqlstring)
	err := rows.Scan(&p.Uuid, &p.Descs, &p.BuildTime)
	if err != nil {

	}
	return &p
}
