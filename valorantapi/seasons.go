package valorantapi

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type SeasonsService service

type Season struct {
	UUID        string    `json:"uuid"`
	DisplayName string    `json:"displayName"`
	Title       *string   `json:"title"`
	Type        *string   `json:"type"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	ParentUUID  *string   `json:"parentUuid"`
	AssetPath   string    `json:"assetPath"`
}

type CompetitiveSeason struct {
	UUID                 string    `json:"uuid"`
	StartTime            time.Time `json:"startTime"`
	EndTime              time.Time `json:"endTime"`
	SeasonUUID           string    `json:"seasonUuid"`
	CompetitiveTiersUUID string    `json:"competitiveTiersUuid"`
	Borders              []Border  `json:"borders"`
	AssetPath            string    `json:"assetPath"`
}

type Border struct {
	UUID         string  `json:"uuid"`
	Level        int     `json:"level"`
	WinsRequired int     `json:"winsRequired"`
	DisplayIcon  string  `json:"displayIcon"`
	SmallIcon    *string `json:"smallIcon"`
	AssetPath    string  `json:"assetPath"`
}

func (s *SeasonsService) List(ctx context.Context, opts *RequestOptions) ([]*Season, *http.Response, error) {
	u := "seasons"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var seasons struct {
		Data []*Season `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &seasons)
	if err != nil {
		return nil, resp, err
	}

	return seasons.Data, resp, nil
}

func (s *SeasonsService) ListCompetitive(ctx context.Context, opts *RequestOptions) ([]*CompetitiveSeason, *http.Response, error) {
	u := "seasons/competitive"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var competitiveSeasons struct {
		Data []*CompetitiveSeason `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &competitiveSeasons)
	if err != nil {
		return nil, resp, err
	}

	return competitiveSeasons.Data, resp, nil
}

func (s *SeasonsService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Season, *http.Response, error) {
	u := fmt.Sprintf("seasons/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var season struct {
		Data *Season `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &season)
	if err != nil {
		return nil, resp, err
	}

	return season.Data, resp, nil
}

func (s *SeasonsService) GetCompetitive(ctx context.Context, uuid string, opts *RequestOptions) (*CompetitiveSeason, *http.Response, error) {
	u := fmt.Sprintf("seasons/competitive/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var competitiveSeason struct {
		Data *CompetitiveSeason `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &competitiveSeason)
	if err != nil {
		return nil, resp, err
	}

	return competitiveSeason.Data, resp, nil
}
