package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate() error {
	migrations := &migrate.FileMigrationSource{
		Dir: "mysql/migrations",
	}

	db, err := sql.Open("mysql", "root:pass@/strings_history?parseTime=true")
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Apply migrate start!")
	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
