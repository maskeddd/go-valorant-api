package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type ContentTiersService service

type ContentTier struct {
	UUID           string `json:"uuid"`
	DisplayName    string `json:"displayName"`
	DevName        string `json:"devName"`
	Rank           int    `json:"rank"`
	JuiceValue     int    `json:"juiceValue"`
	JuiceCost      int    `json:"juiceCost"`
	HighlightColor string `json:"highlightColor"`
	DisplayIcon    string `json:"displayIcon"`
	AssetPath      string `json:"assetPath"`
}

func (s *ContentTiersService) List(ctx context.Context, opts *RequestOptions) ([]*ContentTier, *http.Response, error) {
	u := "contenttiers"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var contentTiers struct {
		Data []*ContentTier `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &contentTiers)
	if err != nil {
		return nil, resp, err
	}

	return contentTiers.Data, resp, nil
}

func (s *ContentTiersService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*ContentTier, *http.Response, error) {
	u := fmt.Sprintf("contenttiers/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var contentTier struct {
		Data *ContentTier `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &contentTier)
	if err != nil {
		return nil, resp, err
	}

	return contentTier.Data, resp, nil
}
