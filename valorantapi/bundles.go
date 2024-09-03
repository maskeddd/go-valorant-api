package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type BundlesService service

type Bundle struct {
	UUID                 string  `json:"uuid"`
	DisplayName          string  `json:"displayName"`
	DisplayNameSubText   *string `json:"displayNameSubText"`
	Description          string  `json:"description"`
	ExtraDescription     *string `json:"extraDescription"`
	PromoDescription     *string `json:"promoDescription"`
	UseAdditionalContext bool    `json:"useAdditionalContext"`
	DisplayIcon          string  `json:"displayIcon"`
	DisplayIcon2         string  `json:"displayIcon2"`
	LogoIcon             *string `json:"logoIcon"`
	VerticalPromoImage   *string `json:"verticalPromoImage"`
	AssetPath            string  `json:"assetPath"`
}

func (s *BundlesService) List(ctx context.Context, opts *RequestOptions) ([]*Bundle, *http.Response, error) {
	u := "bundles"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var bundles struct {
		Data []*Bundle `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &bundles)
	if err != nil {
		return nil, resp, err
	}

	return bundles.Data, resp, nil
}

func (s *BundlesService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Bundle, *http.Response, error) {
	u := fmt.Sprintf("bundles/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var bundle struct {
		Data *Bundle `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &bundle)
	if err != nil {
		return nil, resp, err
	}

	return bundle.Data, resp, nil
}
