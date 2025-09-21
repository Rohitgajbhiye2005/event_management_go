package database

import "database/sql"

type AttenddesModel struct {
	DB *sql.DB
}

type Attenddes struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}
