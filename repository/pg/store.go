package pg

import (
	"database/sql"

	e "main/domain/errors"
	m "main/domain/model"
	rep "main/repository"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) rep.StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) GetCalendar(teacherID int) (*m.CalendarParams, error) {
	ans := m.CalendarParams{}
	row := s.db.QueryRow(`SELECT id, internalApiID FROM calendars WHERE teacherID = $1;`, teacherID)
	if err := row.Scan(&ans.ID, &ans.InternalApiID); err != nil {
		return nil, e.StacktraceError(err)
	}
	return &ans, nil
}

func (s *Store) AddCalendar(teacherID int, calendarID string) (int, error) {
	id := 1
	err := s.db.QueryRow(`INSERT INTO calendars (teacherID, internalApiID) VALUES ($1, $2) RETURNING id;`, teacherID, calendarID).Scan(&id)
	if err != nil {
		return 0, e.StacktraceError(err)
	}
	return id, nil
}
