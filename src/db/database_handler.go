package db

import (
	"database/sql"
	"fmt"

	"atomicmaya.me/digital-overdose-bot/src/db/db_instr"
	"github.com/google/uuid"
)

var DatabaseSink = make(chan *db_instr.DatabaseInstruction)
var DatabaseShutdownHook = make(chan int)

// Is this secure? No.
// Could this memory leak? Yes.
// Am I going to write something to avoid this? Probably.
// TODO
var databaseResponses map[uuid.UUID]*sql.Rows

var RegisterResponseAwaiter(channel chan, id, uuid.uuid) {

}

func InitializeDatabaseSink() error {
	go databaseInstructionAccumulator()
	return fmt.Errorf("not implemented")
}

func DatabaseAddToQueue(instr *db_instr.DatabaseInstruction) uuid.UUID {
	instr.ID = uuid.New()
	DatabaseSink <- instr
	return instr.ID
}

func databaseInstructionAccumulator() {
	handlerRunning := true
	retryCounter := 0
	for handlerRunning {
		instr := <-DatabaseSink
		if instr.InstrType == -1 {
			handlerRunning = false
			DatabaseShutdownHook <- 1
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

func databaseInstructionHandler(instr *db_instr.DatabaseInstruction) (*sql.Rows, error) {
	var stmt *sql.Stmt

	switch instr.InstrType {
	case db_instr.INSERT_USER_EVENT:
		stmt = (*Database).Methods.InsertUserEvent
	case db_instr.LOOKUP_USER_EVENT_ALL:
		stmt = (*Database).Methods.LookupUserEventsAll
	case db_instr.LOOKUP_USER_EVENT_IMPORTANT:
		stmt = (*Database).Methods.LookupUserEventsImportant
	case db_instr.LOOKUP_USER_EVENT_SERVER:
		stmt = (*Database).Methods.LookupUserEventsServer
	case db_instr.LOOKUP_USER_EVENT_STATS:
		stmt = (*Database).Methods.LookupUserEventsStats
	case db_instr.INSERT_OPS_EVENT:
		stmt = (*Database).Methods.InsertOpsEvent
	case db_instr.LOOKUP_OPS_EVENT_ALL:
		stmt = (*Database).Methods.LookupOpsEventsAll
	case db_instr.LOOKUP_OPS_EVENT_IMPORTANT:
		stmt = (*Database).Methods.LookupOpsEventsImportant
	case db_instr.INSERT_ROLE_TRACKING:
		stmt = (*Database).Methods.InsertRoleTracking
	case db_instr.LOOKUP_ROLE_TRACKING:
		stmt = (*Database).Methods.LookupRoleTracking
	}

	return stmt.Query(instr.Params...)
}
