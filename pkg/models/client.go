package models

import "time"

type ClientType int

const (
	VIP ClientType = iota
	Regular
)

type Client struct {
	ID         int64      `json:"id"`
	DPI        string     `json:"dpi"`
	Name       string     `json:"name"`
	ClientType ClientType `json:"clientType"`
}

type HistoryEntry struct {
	ID                 int64     `json:"id"`
	Client             Client    `json:"client"`
	AttendedBy         User      `json:"attendedBy"`
	RequiredOperations int       `json:"requiredOperations"`
	AttendedAt         time.Time `json:"attendedAt"`
}
