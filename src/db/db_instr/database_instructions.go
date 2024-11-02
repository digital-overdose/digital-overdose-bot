package db_instr

import (
	"github.com/google/uuid"
)

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
	InstrType InstrType
	Params    []any
}
