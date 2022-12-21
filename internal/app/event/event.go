package event

import (
	"fmt"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	EventDao *model.Event
	VideoDao *model.Video
}

func NewService() *Service {
	return &Service{
		EventDao: &model.Event{},
	}
}

func (s *Service) CreateEvent(userOpenID, courtID int32, startTime, endTime time.Time) (*model.Event, error) {
	// create event
	event, err := s.EventDao.Create(&model.Event{
		OpenID:      userOpenID,
		CourtID:     courtID,
		StartTime:   startTime,
		EndTime:     endTime,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	// create video record
	currentDate := time.Now().Format("2006-01-02")
	videoUrl := fmt.Sprintf("%s%s/%s_%s", "url", currentDate, startTime, endTime)
	_, err = s.VideoDao.Create(&model.Video{Url: videoUrl})
	if err != nil {
		return nil, err
	}
	return event, err
}

func (s *Service) GetEventsByUser(userOpenID int32) ([]model.Event, error) {
	events, err := s.EventDao.GetsByDesc(&model.Event{OpenID: userOpenID})
	if err != nil {
		return nil, err
	}
	return events, nil
}
