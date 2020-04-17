package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Barang struct {
	Id          int
	NamaBarang  string
	HargaBarang string
}

func koneksi() (*sql.DB, error) {
	var username string = "root"
	var password string = ""
	var host string = "localhost"
	var database string = "balajar_go"
	db, err := sql.Open("mysql", fmt.Sprintf("%s@%stcp(%s:3306)/%s", username, password, host, database))
	if err != nil {
		return nil, err
	}
	return db, nil

}

func sql_tampil() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from tbl_barang")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Barang

	for rows.Next() {
		var each = Barang{}
		var err = rows.Scan(&each.Id, &each.NamaBarang, &each.HargaBarang)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, row := range result {
		fmt.Println(row.NamaBarang)
	}
}

func main() {
	sql_tampil()
}
