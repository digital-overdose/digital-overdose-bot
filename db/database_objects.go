package database_utils

import "time"

// Stores the fields relevant to a member warn as stored in the database.
type Warn struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	WarnTime   time.Time `json:"time"`
	WarnReason string    `json:"reason"`
	Revoked    bool
}

// Stores the fields relevant to a member ban as stored in the database.
type Ban struct {
	ID        int
	UserID    int
	BanTime   time.Time
	BanReason string
	Revoked   bool
}

// Stores the fields relevant to a member mute as stored in the database.
// As mutes are able to expire, we store more data than for warns and bans.
type Mute struct {
	ID             int
	UserID         int
	MuteTime       time.Time
	MuteExpiration time.Time
	MuteReason     string
	Roles          string
	Revoked        bool
}

// Helper structure to help unmarshal a result of a count request.
type Count struct {
	Count int
}
