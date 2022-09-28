package database_utils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseUtilities struct {
	db      *sql.DB
	Methods *DatabaseMethods
}

type DatabaseMethods struct {
	InsertWarn *sql.Stmt
	ListWarns  *sql.Stmt
	RemoveWarn *sql.Stmt
	InsertMute *sql.Stmt
	ListeMutes *sql.Stmt
	RemoveMute *sql.Stmt
	InsertBan  *sql.Stmt
	ReasonBan  *sql.Stmt
	RemoveBan  *sql.Stmt
}

var Database *DatabaseUtilities

const file string = "digital-overdose.db"

const createWarnsTable string = `
  CREATE TABLE IF NOT EXISTS warns (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  warn_time DATETIME NOT NULL,
	warn_reason VARCHAR(255) NOT NULL
);`

const createBansTable string = `
  CREATE TABLE IF NOT EXISTS bans (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  ban_time DATETIME NOT NULL,
	ban_reason VARCHAR(255) NOT NULL
);`

const createMutesTable string = `
  CREATE TABLE IF NOT EXISTS mutes (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  expiration_time DATETIME NOT NULL,
	mute_reason VARCHAR(255) NOT NULL,
	roles TEXT NOT NULL
);`

func InitializeDatabase() (*DatabaseUtilities, error) {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		log.Printf("HERE 1")
		return nil, err
	}

	if _, err := db.Exec(createWarnsTable); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createBansTable); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createMutesTable); err != nil {
		return nil, err
	}

	insert_warn, err := db.Prepare("INSERT INTO warns VALUES(NULL,?,?,?);")
	if err != nil {
		return nil, err
	}

	list_warns, err := db.Prepare("SELECT id, user_id, warn_time, warn_reason FROM warns WHERE user_id=?;")
	if err != nil {
		return nil, err
	}

	remove_warn, err := db.Prepare("DELETE FROM warns WHERE id=?;")
	if err != nil {
		return nil, err
	}

	insert_mute, err := db.Prepare("INSERT INTO mutes VALUES(NULL,?,?,?,?);")
	if err != nil {
		return nil, err
	}

	log.Printf("[+] Database initialized.")

	return &DatabaseUtilities{
		db: db,
		Methods: &DatabaseMethods{
			InsertWarn: insert_warn,
			ListWarns:  list_warns,
			RemoveWarn: remove_warn,
			InsertMute: insert_mute,
		},
	}, nil
}
