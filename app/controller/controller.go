package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Username struct {
	Id   int
	Name string
}

func (u Username) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s", u.Id, u.Name)
}

var tablename string = "username"

/**
 * コネクションを開く。これだけは共通なので、関数にしておく。
 * closeはしていないので、呼び出し元でdefer closeすること。
 */
func open(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	return db
}

func SelectAll(filepath string) []Username {
	db := open(filepath)
	defer db.Close()

	sqlStmt := "SELECT id, name FROM " + tablename

	rows, err := db.Query(sqlStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	usernameList := []Username{}
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		usernameList = append(usernameList, Username{id, name})
	}
	return usernameList
}

func Insert(filepath string, name string) {
	db := open(filepath)
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", tablename)
	_, err = tx.Exec(sql, name)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func DeleteAll(filepath string) {
	db := open(filepath)
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf("DELETE FROM %s", tablename)
	_, err = tx.Exec(sql)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
