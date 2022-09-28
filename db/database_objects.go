package database_utils

import "time"

type Warn struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	WarnTime   time.Time `json:"time"`
	WarnReason string    `json:"reason"`
}

type Ban struct {
	ID        int
	UserID    int
	BanTime   time.Time
	BanReason string
}

type Mute struct {
	ID             int
	UserID         int
	MuteExpiration time.Time
	MuteReason     string
	Roles          []string
}
