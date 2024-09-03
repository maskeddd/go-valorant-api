package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type PlayerTitlesService service

type PlayerTitle struct {
	UUID               string  `json:"uuid"`
	DisplayName        *string `json:"displayName"`
	TitleText          *string `json:"titleText"`
	IsHiddenIfNotOwned bool    `json:"isHiddenIfNotOwned"`
	AssetPath          string  `json:"assetPath"`
}

func (s *PlayerTitlesService) List(ctx context.Context, opts *RequestOptions) ([]*PlayerTitle, *http.Response, error) {
	u := "playertitles"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var playerTitles struct {
		Data []*PlayerTitle `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &playerTitles)
	if err != nil {
		return nil, resp, err
	}

	return playerTitles.Data, resp, nil
}

func (s *PlayerTitlesService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*PlayerTitle, *http.Response, error) {
	u := fmt.Sprintf("playertitles/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var playerTitle struct {
		Data *PlayerTitle `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &playerTitle)
	if err != nil {
		return nil, resp, err
	}

	return playerTitle.Data, resp, nil
}
