package ctrl

//пока не смог подключить

import (
	context "context"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	usc "main/usecase"
	"strconv"
	"strings"
	"time"
)

type CtrlServiceInterface interface {
	GetEvents(classID int) (model.CalendarEvents, error)
	// CreateEvent(EventData) error
	// UpdateEvent(EventData) error
	// DeleteEvent(DeleteEventRequest) error
}

type CtrlService struct {
	UnimplementedCalendarControllerServer
	uc        usc.UsecaseInterface
	urlDomain string
}

func NewCtrlService(uc usc.UsecaseInterface, ud string) *CtrlService {
	return &CtrlService{
		uc:        uc,
		urlDomain: ud,
	}
}

func (cs *CtrlService) GetEvents(ctx context.Context, req *GetEventsRequest) (*GetEventsResponse, error) {
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

	//ans := []model.CalendarEvent{}
	ans := GetEventsResponse{}
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

		time1, err := time.Parse(time.RFC3339, item.Start.DateTime)
		if err != nil {
			log.Println("Error while parsing date :", err)
			return nil, e.StacktraceError(err)
		}
		time2, err := time.Parse(time.RFC3339, item.End.DateTime)
		if err != nil {
			log.Println("Error while parsing date :", err)
			return nil, e.StacktraceError(err)
		}
		// tmp := model.CalendarEvent{Title: item.Summary, Description: item.Description,
		// 	StartDate: time1, EndDate: time2, ClassID: classID, ID: item.Id}
		tmp := EventData{Title: item.Summary, Description: item.Description,
			StartDate: time1.String(), EndDate: time2.String(), Id: item.Id, ClassID: int32(clID)}

		ans.Events = append(ans.Events, &tmp)
	}

	return &ans, nil
}
