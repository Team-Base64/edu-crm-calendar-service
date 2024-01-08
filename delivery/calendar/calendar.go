package calendar

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	proto "main/delivery/calendar/proto"
	e "main/domain/errors"
	u "main/usecase"
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
