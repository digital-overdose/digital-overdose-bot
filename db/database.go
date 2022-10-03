package database_utils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Stores the database and it's associated pre-configured statements.
type DatabaseUtilities struct {
	db      *sql.DB
	Methods *DatabaseMethods
}

// Stores the various prepared statements that may be made to the database.
type DatabaseMethods struct {
	InsertWarn, ListWarns, RemoveWarn, CountWarns                 *sql.Stmt
	InsertBan, ReasonBan, RemoveBan, CountBans                    *sql.Stmt
	InsertMute, ListMutes, GetCurrentMute, RemoveMute, CountMutes *sql.Stmt
	ActiveMutes, CountActiveMutes                                 *sql.Stmt
}

// The general container for all database related actions.
var Database *DatabaseUtilities

// The name of the database file on disk.
const file string = "digital-overdose.db"

// The SQL strings used for database table creation.
const createWarnsTable string = `
  CREATE TABLE IF NOT EXISTS warns (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  warn_time DATETIME NOT NULL,
	warn_reason VARCHAR(255) NOT NULL,
	revoked BOOL NOT NULL DEFAULT FALSE
);`

const createBansTable string = `
  CREATE TABLE IF NOT EXISTS bans (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  ban_time DATETIME NOT NULL,
	ban_reason VARCHAR(255) NOT NULL,
	revoked BOOL NOT NULL DEFAULT FALSE
);`

const createMutesTable string = `
  CREATE TABLE IF NOT EXISTS mutes (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  mute_time DATETIME NOT NULL,
  expiration_time DATETIME NOT NULL,
	mute_reason VARCHAR(255) NOT NULL,
	roles TEXT NOT NULL,
	revoked BOOL NOT NULL DEFAULT FALSE
);`

// Initializes the database on program start, creating it if need be, and initializing all of the prepared statements as well.
func InitializeDatabase() (*DatabaseUtilities, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	// Creates the various tables.
	if _, err := db.Exec(createWarnsTable); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createBansTable); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createMutesTable); err != nil {
		return nil, err
	}

	// Insertion, Listing, Counting and Removal methods relating to warns.
	insert_warn, err := db.Prepare("INSERT INTO warns VALUES(NULL,?,?,?,0);")
	if err != nil {
		return nil, err
	}

	list_warns, err := db.Prepare("SELECT id, user_id, warn_time, warn_reason, revoked FROM warns WHERE user_id=?;")
	if err != nil {
		return nil, err
	}

	remove_warn, err := db.Prepare("UPDATE warns SET revoked=1 WHERE id=?;")
	if err != nil {
		return nil, err
	}

	count_warn, err := db.Prepare("SELECT id AS count FROM warns ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return nil, err
	}

	// Insertion, Identifying, Counting and Removal methods relating to bans.
	insert_ban, err := db.Prepare("INSERT INTO bans VALUES(NULL,?,?,?,0);")
	if err != nil {
		return nil, err
	}

	reason_ban, err := db.Prepare("SELECT id, user_id, ban_time, ban_reason, revoked FROM bans WHERE user_id=? ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return nil, err
	}

	remove_ban, err := db.Prepare("UPDATE bans SET revoked=1 WHERE id=?;")
	if err != nil {
		return nil, err
	}

	count_ban, err := db.Prepare("SELECT id AS count FROM bans ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return nil, err
	}

	// Insertion, Listing, Counting and Removal methods relating to mutes.
	insert_mute, err := db.Prepare("INSERT INTO mutes VALUES(NULL,?,?,?,?,?,0);")
	if err != nil {
		return nil, err
	}

	get_current_mute, err := db.Prepare("SELECT id, user_id, mute_time, expiration_time, mute_reason, roles, revoked FROM mutes WHERE user_id=? AND revoked=0;")
	if err != nil {
		return nil, err
	}

	list_mutes, err := db.Prepare("SELECT id, user_id, mute_time, expiration_time, mute_reason, roles, revoked FROM mutes WHERE user_id=?;")
	if err != nil {
		return nil, err
	}

	remove_mute, err := db.Prepare("UPDATE mutes SET revoked=1, roles=\"\" WHERE id=?;")
	if err != nil {
		return nil, err
	}

	count_mute, err := db.Prepare("SELECT id AS count FROM mutes ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return nil, err
	}

	// Insertion, Listing, Counting and Removal methods relating to currently active mutes.
	list_active_mutes, err := db.Prepare("SELECT id, user_id, mute_time, expiration_time, mute_reason, roles, revoked FROM mutes WHERE revoked=0;")
	if err != nil {
		return nil, err
	}

	count_active_mutes, err := db.Prepare("SELECT COUNT(*) AS count FROM mutes WHERE revoked=0;")
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
			CountWarns: count_warn,

			InsertBan: insert_ban,
			ReasonBan: reason_ban,
			RemoveBan: remove_ban,
			CountBans: count_ban,

			InsertMute:     insert_mute,
			GetCurrentMute: get_current_mute,
			ListMutes:      list_mutes,
			RemoveMute:     remove_mute,
			CountMutes:     count_mute,

			ActiveMutes:      list_active_mutes,
			CountActiveMutes: count_active_mutes,
		},
	}, nil
}

// Wrapper methods to retrieve the active count of each data structure.
func GetTotalWarnsCount() (int, error) {
	return getTotalCount((*Database.Methods).CountWarns)
}

func GetTotalBansCount() (int, error) {
	return getTotalCount((*Database.Methods).CountBans)
}

func GetTotalMutesCount() (int, error) {
	return getTotalCount((*Database.Methods).CountMutes)
}

func getTotalCount(stmt *sql.Stmt) (int, error) {
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("SQL: Could not get warn count because: '%v'", err)
	}

	defer rows.Close()

	listCount := []Count{}
	for rows.Next() {
		i := Count{}
		err := rows.Scan(&i.Count)
		if err != nil {
			return -1, err
		}
		listCount = append(listCount, i)
	}

	if len(listCount) == 0 {
		return -1, nil
	}

	return listCount[0].Count, nil
}
