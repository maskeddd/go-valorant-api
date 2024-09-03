package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type LevelBordersService service

type LevelBorder struct {
	UUID                      string `json:"uuid"`
	DisplayName               string `json:"displayName"`
	StartingLevel             int    `json:"startingLevel"`
	LevelNumberAppearance     string `json:"levelNumberAppearance"`
	SmallPlayerCardAppearance string `json:"smallPlayerCardAppearance"`
	AssetPath                 string `json:"assetPath"`
}

func (s *LevelBordersService) List(ctx context.Context, opts *RequestOptions) ([]*LevelBorder, *http.Response, error) {
	u := "levelborders"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var levelBorders struct {
		Data []*LevelBorder `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &levelBorders)
	if err != nil {
		return nil, resp, err
	}

	return levelBorders.Data, resp, nil
}

func (s *LevelBordersService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*LevelBorder, *http.Response, error) {
	u := fmt.Sprintf("levelborders/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var levelBorder struct {
		Data *LevelBorder `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &levelBorder)
	if err != nil {
		return nil, resp, err
	}

	return levelBorder.Data, resp, nil
}
