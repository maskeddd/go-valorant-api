package valorantapi

import (
	"context"
	"net/http"
)

type GearService service

type Gear struct {
	UUID        string   `json:"uuid"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
	DisplayIcon string   `json:"displayIcon"`
	AssetPath   string   `json:"assetPath"`
	ShopData    ShopData `json:"shopData"`
}

type GridPosition struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

func (s *GearService) List(ctx context.Context, opts *RequestOptions) ([]*Gear, *http.Response, error) {
	u := "gear"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var gear struct {
		Data []*Gear `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &gear)
	if err != nil {
		return nil, resp, err
	}

	return gear.Data, resp, nil
}
