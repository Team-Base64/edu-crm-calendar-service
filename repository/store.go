package chat

import (
	e "main/domain/errors"
	"main/domain/model"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	GetCalendarDB(teacherID int) (*model.CalendarParams, error)
	CreateCalendarDB(teacherID int, googleID string) (int, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) GetCalendarDB(teacherID int) (*model.CalendarParams, error) {
	ans := model.CalendarParams{}
	row := s.db.QueryRow(`SELECT id, idInGoogle FROM calendars WHERE teacherID = $1;`, teacherID)
	if err := row.Scan(&ans.ID, &ans.IDInGoogle); err != nil {
		return nil, e.StacktraceError(err)
	}
	return &ans, nil
}

func (s *Store) CreateCalendarDB(teacherID int, googleID string) (int, error) {
	id := 1
	err := s.db.QueryRow(`INSERT INTO calendars (teacherID, idInGoogle) VALUES ($1, $2) RETURNING id;`, teacherID, googleID).Scan(&id)
	if err != nil {
		return 0, e.StacktraceError(err)
	}
	return id, nil
}
