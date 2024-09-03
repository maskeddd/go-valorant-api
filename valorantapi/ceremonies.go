package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type CeremoniesService service

type Ceremony struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayName"`
	AssetPath   string `json:"assetPath"`
}

func (s *CeremoniesService) List(ctx context.Context, opts *RequestOptions) ([]*Ceremony, *http.Response, error) {
	u := "ceremonies"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var ceremonies struct {
		Data []*Ceremony `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &ceremonies)
	if err != nil {
		return nil, resp, err
	}

	return ceremonies.Data, resp, nil
}

func (s *CeremoniesService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Ceremony, *http.Response, error) {
	u := fmt.Sprintf("ceremonies/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var ceremony struct {
		Data *Ceremony `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &ceremony)
	if err != nil {
		return nil, resp, err
	}

	return ceremony.Data, resp, nil
}
