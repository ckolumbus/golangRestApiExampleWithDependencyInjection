/**
 * File: db.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package db

import (
	"database/sql"
	"fmt"
	"log"
)

// SetupDB creates the connection to give sqlite3 database
func SetupDB(dbType string, dbConnectString string) *sql.DB {
	var err error
	db, err := sql.Open(dbType, dbConnectString)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
	return db
}
