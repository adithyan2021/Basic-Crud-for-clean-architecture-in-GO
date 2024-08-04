package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shakezidin/internal/config"
)

func ConnectPGDB(cnf config.Config) *sql.DB {
	fmt.Println(cnf)
	postgresURL := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v", cnf.PGDBmsName, cnf.PGPassword, cnf.PGPassword, cnf.PGHost, cnf.PgPort, cnf.PGDBName, cnf.PgSSLMode)
	db, err := sql.Open(cnf.PgDriverName, postgresURL)
	if err != nil {
		log.Fatal(err, err.Error(), "driver name", cnf.PgDriverName, "postgres url", postgresURL)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("not connected to postgres db: ", err.Error())
	}

	log.Println("connected to postgres db successfully!")

	return db
}
