package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	dbpwd  = os.Getenv("dbpwd")
	dbuser = os.Getenv("dbuser")
)

func init() {
	var err error
	if db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/user", dbuser, dbpwd)); err != nil {
		log.Fatal(err)
	}
}

type user struct {
	Id   int64
	Mail string
	Name string
}

func Select() {
	rows, err := db.Query("select id,name,mail from user;")
	log.Println(err)
	for rows.Next() {
		d := &user{}
		rows.Scan(&d.Id, &d.Name, &d.Mail)
		log.Println(d)
	}
}

func main() {
	Select()
}
