package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"bufio"
	"io"
	"strings"
)

var (
	dbhost = "10.25.18.9"
	dbusername = "root"
	dbpassword = "p@ssw0rd"
	dbname = "csdn"
)
func openconnection() *sql.DB {
	dataSourceName := dbusername + ":" + dbpassword + "@tcp(" + dbhost + ":3306" + ")/" + dbname + "?charset=utf8mb4"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func insert(db *sql.DB, s1 string, s2 string, s3 string){
	defer func() {
		if err:=recover();err!=nil {
			fmt.Println("catch the exception")
		}

	}()

	stmtIns, err := db.Prepare("INSERT userinfo SET username=?, password=?, email=?")
	if err != nil{
		panic(err.Error())
	}
	defer stmtIns.Close()
	res, err1 := stmtIns.Exec(s1, s2, s3);

	if err1 != nil {
		panic(err1.Error())
	}


	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)
}
func main()  {

	fi, err := os.Open("D:\\www.csdn.net.sql.txt")
	if err != nil{
		fmt.Println("Error: %s\n", err)
		return
	}
	defer fi.Close()


	db := openconnection()
	defer db.Close()
	br := bufio.NewReader(fi)

	for {
		line, _, err := br.ReadLine()

		if err == io.EOF{
			break
		}
		fmt.Println(string(line))

		s := strings.Split(string(line), " # ")

		insert(db, s[0], s[1], s[2])
	}
}


