package calendar

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	proto "main/delivery/calendar/proto"
	e "main/domain/errors"
	u "main/usecase"

	"google.golang.org/api/calendar/v3"
)

type CtrlService struct {
	proto.UnimplementedCalendarControllerServer
	uc        u.UsecaseInterface
	urlDomain string
}

func NewCtrlService(uc u.UsecaseInterface, urlDomain string) proto.CalendarControllerServer {
	return &CtrlService{
		uc:        uc,
		urlDomain: urlDomain,
	}
}

func (cs *CtrlService) GetEventsCalendar(ctx context.Context, req *proto.GetEventsRequestCalendar) (*proto.GetEventsResponse, error) {
	log.Println("Called GetEventsCalendar")
	srv, err := cs.uc.GetCalendarServiceClient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return nil, e.StacktraceError(err)
	}
	calendarDB, err := cs.uc.GetCalendarDB(int(req.GetTeacherID()))
	if err != nil {
		log.Println("DB err: ", err)
		return nil, e.StacktraceError(err)
	}
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(calendarDB.IDInGoogle).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(100).OrderBy("startTime").Do()
	if err != nil {
		log.Println("Unable to retrieve next ten of the user's events: ", err)
		return nil, e.StacktraceError(err)
	}

	ans := proto.GetEventsResponse{}
	for _, item := range events.Items {
		s := strings.Split(item.Summary, " ")
		clID := 0
		if len(s) > 2 && s[len(s)-2] == "Class" {
			clIDs := s[len(s)-1]
			clID, err = strconv.Atoi(clIDs)
			if err != nil {
				log.Println("error: ", err)
				return nil, e.StacktraceError(err)
			}
		}

		tmp := proto.EventData{
			Title:       item.Summary,
			Description: item.Description,
			StartDate:   item.Start.DateTime,
			EndDate:     item.End.DateTime,
			Id:          item.Id,
			ClassID:     int32(clID),
		}

		ans.Events = append(ans.Events, &tmp)
	}

	return &ans, nil
}

func (cs *CtrlService) CreateEvent(ctx context.Context, req *proto.CreateEventRequest) (*proto.CreateEventResponse, error) {
	log.Println("Called CreateEvent")
	srv, err := cs.uc.GetCalendarServiceClient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return &proto.CreateEventResponse{EventID: ""}, e.StacktraceError(err)
	}

	event := &calendar.Event{
		Summary:     req.Event.Title + " Class " + fmt.Sprintf("%d", req.Event.ClassID),
		Description: req.Event.Description,
		Start: &calendar.EventDateTime{
			DateTime: req.Event.StartDate,
			//TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: req.Event.EndDate,
			//TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}

	event, err = srv.Events.Insert(req.CalendarID, event).Do()
	if err != nil {
		log.Println("Unable to create event: ", err)
		return &proto.CreateEventResponse{EventID: ""}, e.StacktraceError(err)
	}

	return &proto.CreateEventResponse{EventID: event.Id}, nil
}

func (cs *CtrlService) DeleteEvent(ctx context.Context, req *proto.DeleteEventRequest) (*proto.Nothing, error) {
	log.Println("Called DeleteEvent")
	srv, err := cs.uc.GetCalendarServiceClient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}

	err = srv.Events.Delete(req.CalendarID, req.Id).Do()
	if err != nil {
		log.Println("Unable to delete event: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}
	return &proto.Nothing{}, nil
}

func (cs *CtrlService) UpdateEvent(ctx context.Context, req *proto.UpdateEventRequest) (*proto.Nothing, error) {
	log.Println("Called UpdateEvent")
	srv, err := cs.uc.GetCalendarServiceClient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}

	event := &calendar.Event{
		Summary:     req.Event.Title,
		Description: req.Event.Description,
		Start: &calendar.EventDateTime{
			DateTime: req.Event.StartDate,
			//TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: req.Event.EndDate,
			//TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}

	event, err = srv.Events.Update(req.CalendarID, req.Event.Id, event).Do()
	if err != nil {
		log.Println("Unable to update event: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}

	return &proto.Nothing{}, nil
}

func (cs *CtrlService) CreateCalendar(ctx context.Context, req *proto.CreateCalendarRequest) (*proto.Nothing, error) {
	log.Println("Called CreateCalendar")
	srv, err := cs.uc.GetCalendarServiceClient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}

	newCal := &calendar.Calendar{TimeZone: "Europe/Moscow", Summary: "EDUCRM Calendar"}
	cal, err := srv.Calendars.Insert(newCal).Do()
	if err != nil {
		log.Println("Unable to create calendar: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}

	Acl := &calendar.AclRule{Scope: &calendar.AclRuleScope{Type: "default"}, Role: "reader"}
	_, err = srv.Acl.Insert(cal.Id, Acl).Do()
	if err != nil {
		log.Println("Unable to create ACL: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}

	_, err = cs.uc.CreateCalendarDB(int(req.TeacherID), cal.Id)
	if err != nil {
		log.Println("DB err: ", err)
		return &proto.Nothing{}, e.StacktraceError(err)
	}
	return &proto.Nothing{}, nil
}
