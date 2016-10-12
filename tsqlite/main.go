package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Println(err)
	}

	//tx, err := db.Begin()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()

	//// Insert
	//for i := 0; i < 100; i++ {
	//	_, err = stmt.Exec(i, fmt.Sprintf("No. %d", i))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		log.Println(id, name)
	}

	var selectid, order, from int
	var detail string
	err = db.QueryRow("EXPLAIN QUERY PLAN select id, name from foo where name='No. 3'").Scan(&selectid, &order, &from, &detail)
	if err != nil {
		log.Panic(err)
	}
	log.Println(selectid, order, from, detail)

}
