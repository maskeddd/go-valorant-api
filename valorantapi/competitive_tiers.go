package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type CompetitiveTiersService service

type CompetitiveTier struct {
	UUID            string `json:"uuid"`
	AssetObjectName string `json:"assetObjectName"`
	Tiers           []Tier `json:"tiers"`
	AssetPath       string `json:"assetPath"`
}

type Tier struct {
	Tier                 int     `json:"tier"`
	TierName             string  `json:"tierName"`
	Division             string  `json:"division"`
	DivisionName         string  `json:"divisionName"`
	Color                string  `json:"color"`
	BackgroundColor      string  `json:"backgroundColor"`
	SmallIcon            *string `json:"smallIcon"`
	LargeIcon            *string `json:"largeIcon"`
	RankTriangleDownIcon *string `json:"rankTriangleDownIcon"`
	RankTriangleUpIcon   *string `json:"rankTriangleUpIcon"`
}

func (s *CompetitiveTiersService) List(ctx context.Context, opts *RequestOptions) ([]*CompetitiveTier, *http.Response, error) {
	u := "competitivetiers"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var competitiveTiers struct {
		Data []*CompetitiveTier `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &competitiveTiers)
	if err != nil {
		return nil, resp, err
	}

	return competitiveTiers.Data, resp, nil
}

func (s *CompetitiveTiersService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*CompetitiveTier, *http.Response, error) {
	u := fmt.Sprintf("competitivetiers/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var competitiveTier struct {
		Data *CompetitiveTier `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &competitiveTier)
	if err != nil {
		return nil, resp, err
	}

	return competitiveTier.Data, resp, nil
}
