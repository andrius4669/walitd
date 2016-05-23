package dbacc

import (
	cfg "../configmgr"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// opens sql connection with parameters from config
func OpenSQL() *sql.DB {
	dbinfo := ""
	hst, _ := cfg.GetOption("sql.host")
	if hst != "" {
		dbinfo = dbinfo + fmt.Sprintf("host=%s ", hst)
	}
	prt, _ := cfg.GetOption("sql.port")
	if prt != "" {
		dbinfo = dbinfo + fmt.Sprintf("port=%s ", prt)
	}
	usr, _ := cfg.GetOption("sql.user")
	if usr != "" {
		dbinfo = dbinfo + fmt.Sprintf("user=%s ", usr)
	}
	pwd, _ := cfg.GetOption("sql.password")
	if pwd != "" {
		dbinfo = dbinfo + fmt.Sprintf("password=%s ", pwd)
	}
	dbn, _ := cfg.GetOption("sql.dbname")
	dbinfo = dbinfo + fmt.Sprintf("dbname=%s sslmode=disable", dbn)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}
