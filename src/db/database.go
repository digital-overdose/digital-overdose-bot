package database_utils

import (
	"database/sql"
	"fmt"
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
	InsertUserEvent, LookupUserEvents, LookupUserEventsImportant, LookupUserEventsServer, LookupUserEventsStats *sql.Stmt
	InsertOpsEvent, LookupOpsEvents, LookupOpsEventsImportant                                                   *sql.Stmt
	InsertRoleTracking, LookupRoleTracking                                                                      *sql.Stmt
}

// The general container for all database related actions.
var Database *DatabaseUtilities

// The name of the database file on disk.
const file string = "digital-overdose.db"

const createUserEventsTable string = `
	CREATE TABLE IF NOT EXISTS user_events (
		id INTEGER NOT NULL PRIMARY KEY,
		type INTEGER NOT NULL,
		time DATETIME NOT NULL,
		subject_id INTEGER NOT NULL,
		subject_name VARCHAR(255) NOT NULL,
		source_id INTEGER,
		source_name VARCHAR(255),
		data TEXT NOT NULL
	)
`

const createOpsEventsTable string = `
	CREATE TABLE IF NOT EXISTS ops_events (
		id INTEGER NOT NULL PRIMARY KEY,
		type INTEGER NOT NULL,
		time DATETIME NOT NULL,
		data TEXT NOT NULL
	)
`

const createRolesTrackingTable string = `
	CREATE TABLE IF NOT EXISTS roles_tracker (
		id INTEGER NOT NULL PRIMARY KEY,
		event_id INTEGER NOT NULL,
		subject_id INTEGER NOT NULL,
		roles TEXT NOT NULL,

		FOREIGN KEY(event_id) REFERENCES events(id)
	)
`

// Initializes the database on program start, creating it if need be, and initializing all of the prepared statements as well.
func InitializeDatabase() (*DatabaseUtilities, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	// Creates the various tables.
	if _, err := db.Exec(createUserEventsTable); err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [createUserEventsTable] %w", err)
	}

	if _, err := db.Exec(createOpsEventsTable); err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [createOpsEventsTable] %w", err)
	}

	if _, err := db.Exec(createRolesTrackingTable); err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [createRolesTrackingTable] %w", err)
	}

	insertUserEvent, err := db.Prepare("INSERT INTO user_events VALUES(NULL,?,?,?,?,?,?,?) RETURNING id;")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [insertUserEvent] %w", err)
	}

	lookupUserEventsAll, err := db.Prepare("SELECT * FROM user_events WHERE subject_id=? ORDER BY ID DESC LIMIT 20 OFFSET ?;")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [lookupUserEventsAll] %w", err)
	}

	lookupUserEventsImportant, err := db.Prepare("SELECT * FROM user_events WHERE subject_id=? AND type IN (0,1,4,5,6,7,8,9) ORDER BY ID DESC;")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [lookupUserEventsImportant] %w", err)
	}

	lookupUserEventsServer, err := db.Prepare("SELECT * FROM user_events WHERE subject_id=? AND type IN (0,1,10,11);")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [lookupUserEventsServer] %w", err)
	}

	lookupUserEventsStats, err := db.Prepare("SELECT type AS event_type, COUNT(*) AS result FROM user_events WHERE subject_id=? AND type IN (16,17,18,32,33,34,35);")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [lookupUserEventsStats] %w", err)
	}

	insertOpsEvent, err := db.Prepare("INSERT into ops_events VALUES(NULL,?,?,?);")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [insertOpsEvent] %w", err)
	}

	lookupOpsEventsAll, err := db.Prepare("SELECT * FROM ops_events ORDER BY ID DESC LIMIT 20 OFFSET ?;")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [lookupOpsEventsAll] %w", err)
	}

	lookupOpsEventsImportant, err := db.Prepare("SELECT * FROM ops_events WHERE type IN (160, 161, 162) ORDER BY ID DESC LIMIT 20 OFFSET ?;")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [lookupOpsEventsImportant] %w", err)
	}

	insertRolesTrackingLog, err := db.Prepare("INSERT INTO roles_tracker VALUES(NULL,?,?,?);")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [insertRolesTrackingLog] %w", err)
	}

	LookupRolesTrackingLog, err := db.Prepare("SELECT * FROM roles_tracker WHERE subject_id=? ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return nil, fmt.Errorf("[db.InitializeDatabase] [LookupRolesTrackingLog] %w", err)
	}

	log.Printf("[+] Database initialized.")

	return &DatabaseUtilities{
		db: db,
		Methods: &DatabaseMethods{
			InsertUserEvent:           insertUserEvent,
			LookupUserEvents:          lookupUserEventsAll,
			LookupUserEventsImportant: lookupUserEventsImportant,
			LookupUserEventsServer:    lookupUserEventsServer,
			LookupUserEventsStats:     lookupUserEventsStats,

			InsertOpsEvent:           insertOpsEvent,
			LookupOpsEvents:          lookupOpsEventsAll,
			LookupOpsEventsImportant: lookupOpsEventsImportant,

			InsertRoleTracking: insertRolesTrackingLog,
			LookupRoleTracking: LookupRolesTrackingLog,
		},
	}, nil
}

// // Wrapper methods to retrieve the active count of each data structure.
// func GetTotalWarnsCount() (int, error) {
// 	return getTotalCount((*Database.Methods).CountWarns)
// }

// func GetTotalBansCount() (int, error) {
// 	return getTotalCount((*Database.Methods).CountBans)
// }

// func GetTotalMutesCount() (int, error) {
// 	return getTotalCount((*Database.Methods).CountMutes)
// }

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
