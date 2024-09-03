package valorantapi

import (
	"context"
	"net/http"
)

type CurrenciesService service

type Currency struct {
	UUID                string `json:"uuid"`
	DisplayName         string `json:"displayName"`
	DisplayNameSingular string `json:"displayNameSingular"`
	DisplayIcon         string `json:"displayIcon"`
	LargeIcon           string `json:"largeIcon"`
	AssetPath           string `json:"assetPath"`
}

func (s *CurrenciesService) List(ctx context.Context, opts *RequestOptions) ([]*Currency, *http.Response, error) {
	u := "currencies"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var currencies struct {
		Data []*Currency `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &currencies)
	if err != nil {
		return nil, resp, err
	}

	return currencies.Data, resp, nil
}

func (s *CurrenciesService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Currency, *http.Response, error) {
	u := "currencies/" + uuid
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var currency struct {
		Data *Currency `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &currency)
	if err != nil {
		return nil, resp, err
	}

	return currency.Data, resp, nil
}
