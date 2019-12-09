package main

import (
	"database/sql"
	"fmt"
	"strconv"
)
import _ "github.com/lib/pq"

type Users struct {
	uuid         sql.NullString
	company_uuid sql.NullString
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"dbi.mshare.cn", 1094, "dbuser", "dY8*6fN6Z#xSOg$wG9zDATTe", "sxsdb")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	//db.SetConnMaxLifetime()
	var user Users
	row := db.QueryRow("select uuid, company_uuid from users where uuid=$1", "usr_xiafo5poeszu")
	err = row.Scan(&user.uuid, &user.company_uuid)
	if err == sql.ErrNoRows {
		fmt.Println(123)
	}
	//后续处理
	fmt.Println(user.company_uuid.String)
	fmt.Println(user.uuid.String)
}
