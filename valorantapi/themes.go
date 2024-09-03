package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type ThemesService service

type Theme struct {
	UUID               string  `json:"uuid"`
	DisplayName        string  `json:"displayName"`
	DisplayIcon        *string `json:"displayIcon"`
	StoreFeaturedImage *string `json:"storeFeaturedImage"`
	AssetPath          string  `json:"assetPath"`
}

func (s *ThemesService) List(ctx context.Context, opts *RequestOptions) ([]*Theme, *http.Response, error) {
	u := "themes"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var themes struct {
		Data []*Theme `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &themes)
	if err != nil {
		return nil, resp, err
	}

	return themes.Data, resp, nil
}

func (s *ThemesService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Theme, *http.Response, error) {
	u := fmt.Sprintf("themes/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var theme struct {
		Data *Theme `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &theme)
	if err != nil {
		return nil, resp, err
	}

	return theme.Data, resp, nil
}
