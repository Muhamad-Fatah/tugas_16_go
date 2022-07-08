package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type masyarakat struct {
	ID      string
	Nama    string
	Umur    int
	Profesi string
}

var data []masyarakat

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/tugasgolang")

	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	getDB()
}

func getDB() {
	// Open db
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// Open table
	rows, err := db.Query("SELECT * FROM masyarakat")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	// For each table
	for rows.Next() {
		var each = masyarakat{}
		var err = rows.Scan(&each.ID, &each.Nama, &each.Umur, &each.Profesi)

		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
		fmt.Println(each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
