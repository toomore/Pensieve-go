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
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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

func SelectOne(id int64) {
	var d = &user{}
	if err := db.QueryRow("select name,mail from user where id=?", id).Scan(&d.Name, &d.Mail); err == nil {
		log.Println(d)
	} else {
		log.Println(err)
	}
}

func Insert() {
	stmt, err := db.Prepare("Insert into user(name, mail, age) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if err == nil {
		if result, err := stmt.Exec("Toomore", "toomore0929@gmail.com", 30); err == nil {
			log.Println(result.LastInsertId())
			log.Println(result.RowsAffected())
		}
	}
}

func Update() {
	stmt, err := db.Prepare("update user set mail=? where id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if err == nil {
		if result, err := stmt.Exec("me@toomore.net", 121); err == nil {
			log.Println(result.LastInsertId())
			log.Println(result.RowsAffected())
		}
	}

}

func main() {
	defer db.Close()
	Select()
	SelectOne(99)
	Insert()
	Update()
}
