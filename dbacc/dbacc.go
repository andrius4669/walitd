package dbacc

import (
	cfg "../configmgr"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// opens sql connection with parameters from config
func OpenSQL() *sql.DB {
	usr, _ := cfg.GetOption("sql.user")
	pwd, _ := cfg.GetOption("sql.password")
	dbn, _ := cfg.GetOption("sql.database")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", usr, pwd, dbn)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}
