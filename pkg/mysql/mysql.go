package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xiam/to"
	"log"
	"os"
)

var Db *sql.DB

func InitMySQL() {
	db := os.Getenv("DB_DATABASE_PASSPORT")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	maxConn := to.Int(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn := to.Int(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))

	var err error
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, db))
	if err != nil {
		log.Panicln("数据库打开出现了问题：", err.Error())
	}
	err = Db.Ping()
	if err != nil {
		log.Panicln("数据库连接出现了问题：", err.Error())
		return
	}

	Db.SetMaxOpenConns(maxConn)
	Db.SetMaxIdleConns(maxIdleConn)
}
