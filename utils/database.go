package utils

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"blog/const"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() *sql.DB {
	// if db != nil {
	// 	return db
	// }
		//fmt.Println("init")
	connectionString := fmt.Sprintf("%s:%s@/%s",
		constants.DBUSER,
		constants.DBPASSWORD,
		constants.DBNAME,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	db = InitDB()
	// if db == nil {
	// 	db = InitDB()
	// }
	return db
}

// CloseDB closes the database connection
func CloseDB() {
	//fmt.Println("close:", db)
	if db != nil {
		db.Close()
	}

}
