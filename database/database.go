package database

import (
	"database/sql"
	"log"
	"time"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(env Config) *Database {
	db, err := sql.Open(env.DB_CONNECTION, NewDSN(env).FormatDSN())
	if err != nil {
		log.Fatalf("DB Error: %v", err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return &Database{DB: db}
}

func (db *Database) Close() {
	if err := db.DB.Close(); err != nil {
		log.Fatalf("DB_Error: %v", err)
		return
	}
	log.Println("DB: Connection close")
}

func (db *Database) Ping() {
	if err := db.DB.Ping(); err != nil {
		log.Fatalf("DB_Error: %v", err)
		return
	}
	log.Println("Database connected !!!")
}

func (db *Database) Status() sql.DBStats {
	return db.DB.Stats()
}
