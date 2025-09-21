package database

import "database/sql"

type Models struct {
	Users     UserModel
	Events    EventModel
	Attenddes AttenddesModel
}

func NewModels(db *sql.DB)Models{
	return Models{
		Users: UserModel{DB:db},
		Events: EventModel{DB:db},
		Attenddes: AttenddesModel{DB:db},
	}
}