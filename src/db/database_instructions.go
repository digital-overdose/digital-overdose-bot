package database_utils

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

var DatabaseSink = make(chan *DatabaseInstruction)

// Is this secure? No.
// Could this memory leak? Yes.
// Am I going to write something to avoid this? Probably.
// TODO
var DatabaseResponses map[uuid.UUID]*sql.Rows

type InstrType int

const (
	INSERT_USER_EVENT InstrType = 0x00

	LOOKUP_USER_EVENT_ALL       InstrType = 0x10
	LOOKUP_USER_EVENT_IMPORTANT InstrType = 0x11
	LOOKUP_USER_EVENT_SERVER    InstrType = 0x12
	LOOKUP_USER_EVENT_STATS     InstrType = 0x13

	INSERT_OPS_EVENT InstrType = 0x30

	LOOKUP_OPS_EVENT_ALL       InstrType = 0x40
	LOOKUP_OPS_EVENT_IMPORTANT InstrType = 0x41

	INSERT_ROLE_TRACKING InstrType = 0x60

	LOOKUP_ROLE_TRACKING InstrType = 0x70
)

// Tries to formalise the approach to be able to pipe instructions to a channel.
type DatabaseInstruction struct {
	ID        uuid.UUID
	InstrType *InstrType
	Params    []any
}

func InitializeDatabaseSink() error {
	go databaseInstructionAccumulator()
	return fmt.Errorf("not implemented")
}

func databaseInstructionAccumulator() {
	handlerRunning := true
	retryCounter := 0
	for handlerRunning {
		instr := <-DatabaseSink
		if instr.InstrType == nil {
			handlerRunning = false
		} else {
			if res, err := databaseInstructionHandler(instr); err != nil {
				if retryCounter < 3 {
					DatabaseSink <- instr
					retryCounter += 1
				} else {
					retryCounter = 0
					// TODO LOG ERROR
				}
			} else {
				DatabaseResponses[instr.ID] = res
			}
		}
	}
}

func databaseInstructionHandler(instr *DatabaseInstruction) (*sql.Rows, error) {
	var stmt *sql.Stmt

	switch *instr.InstrType {
	case INSERT_USER_EVENT:
		stmt = (*Database).Methods.InsertUserEvent
	case LOOKUP_USER_EVENT_ALL:
		stmt = (*Database).Methods.LookupUserEventsAll
	case LOOKUP_USER_EVENT_IMPORTANT:
		stmt = (*Database).Methods.LookupUserEventsImportant
	case LOOKUP_USER_EVENT_SERVER:
		stmt = (*Database).Methods.LookupUserEventsServer
	case LOOKUP_USER_EVENT_STATS:
		stmt = (*Database).Methods.LookupUserEventsStats
	case INSERT_OPS_EVENT:
		stmt = (*Database).Methods.InsertOpsEvent
	case LOOKUP_OPS_EVENT_ALL:
		stmt = (*Database).Methods.LookupOpsEventsAll
	case LOOKUP_OPS_EVENT_IMPORTANT:
		stmt = (*Database).Methods.LookupOpsEventsImportant
	case INSERT_ROLE_TRACKING:
		stmt = (*Database).Methods.InsertRoleTracking
	case LOOKUP_ROLE_TRACKING:
		stmt = (*Database).Methods.LookupRoleTracking
	}

	return stmt.Query(instr.Params...)
}
