package database_utils

import "time"

type Warn struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	WarnTime   time.Time `json:"time"`
	WarnReason string    `json:"reason"`
	Revoked    bool
}

type Ban struct {
	ID        int
	UserID    int
	BanTime   time.Time
	BanReason string
	Revoked   bool
}

type Mute struct {
	ID             int
	UserID         int
	MuteTime       time.Time
	MuteExpiration time.Time
	MuteReason     string
	Roles          string
	Revoked        bool
}

type Count struct {
	count int
}
