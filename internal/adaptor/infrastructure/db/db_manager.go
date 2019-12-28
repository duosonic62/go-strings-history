package db

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"time"
)

// sqlboilerのセットアップ
func InitDB() {
	dns := "root:pass@/strings_history?charset=utf8mb4&parseTime=true"
	con, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	// connection pool settings
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(300 * time.Second)

	// global connection setting
	boil.SetDB(con)
	boil.DebugMode = true
}
