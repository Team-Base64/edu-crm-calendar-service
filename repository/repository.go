package chat

import (
	m "main/domain/model"
)

type StoreInterface interface {
	GetCalendar(teacherID int) (*m.CalendarParams, error)
	AddCalendar(teacherID int, calendarID string) (int, error)
}
