package valorantapi

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type EventsService service

type Event struct {
	UUID             string    `json:"uuid"`
	DisplayName      string    `json:"displayName"`
	ShortDisplayName string    `json:"shortDisplayName"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	AssetPath        string    `json:"assetPath"`
}

func (s *EventsService) List(ctx context.Context, opts *RequestOptions) ([]*Event, *http.Response, error) {
	u := "events"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var events struct {
		Data []*Event `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &events)
	if err != nil {
		return nil, resp, err
	}

	return events.Data, resp, nil
}

func (s *EventsService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Event, *http.Response, error) {
	u := fmt.Sprintf("events/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var event struct {
		Data *Event `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &event)
	if err != nil {
		return nil, resp, err
	}

	return event.Data, resp, nil
}
