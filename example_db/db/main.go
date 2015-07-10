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
	defer rows.Close()

	log.Println(err)
	datas := make([]*user, 0)
	for rows.Next() {
		d := &user{}
		rows.Scan(&d.Id, &d.Name, &d.Mail)
		datas = append(datas, d)
	}
	for _, v := range datas {
		log.Println(v.Id, v)
	}
}

func main() {
	defer db.Close()
	Select()
}
