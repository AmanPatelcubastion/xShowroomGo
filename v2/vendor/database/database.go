package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	SQL *gorm.DB
)

func Connect(username string, password string, hostname string, port int, dbname string) {
	var err error
	SQL, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, dbname))
	if err != nil {
		fmt.Println("SQL Driver Error", err)
	}
}
