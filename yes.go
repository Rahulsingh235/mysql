package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type n1 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func connectMysql() {

}
func main() {
	db, err := sql.Open("mysql", "rahul:Rahul@235@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO user VALUES ('1','Yash')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	//	delete, err := db.Query("DELETE FROM USER WHERE name LIKE'Rahul'")
	///	if err != nil {
	//		panic(err.Error())
	//	}
	//	defer delete.Close()
	//result, err := db.Query("SELECT id,name from user")
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	defer result.Close()
	//	var n2 n1
	//	for result.Next() {
	//		err := result.Scan(&n2.Id, &n2.Name)
	//		if err != nil {
	//			panic(err.Error())
	//		}
	//	}
	//	fmt.Print(n2.Name)
}
