package models

import (
	"database/sql"
	"dev-framework-go/pkg/db"
	"fmt"
)

//测试数据库
type Project struct {
	Uuid      sql.NullString
	PjName    sql.NullString
	ShowName  sql.NullString
	Descs     sql.NullString
	BuildTime sql.NullString
	Status    sql.NullString
}

//查询单条
func SelectTags(sqlstring string, parm ...interface{}) *Project {
	var p Project
	fmt.Println(sqlstring)
	rows := db.DBPool.QueryRow(sqlstring, parm...)
	err := rows.Scan(&p.Uuid, &p.PjName, &p.ShowName, &p.Descs, &p.BuildTime, &p.Status)
	if err != nil {

	}
	return &p
}
