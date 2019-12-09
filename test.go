package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/lib/pq"

type Users struct {
	uuid string
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
	row := db.QueryRow("select uuid from users where uuid=$1", "usr_xiafo5poeszu")
	err = row.Scan(&user.uuid)
	if err == sql.ErrNoRows {
		fmt.Println(123)
	}
	fmt.Println(user)

}
