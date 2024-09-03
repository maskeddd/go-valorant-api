package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type SpraysService service

type Spray struct {
	UUID                string       `json:"uuid"`
	DisplayName         string       `json:"displayName"`
	Category            *string      `json:"category"`
	ThemeUUID           *string      `json:"themeUuid"`
	IsNullSpray         bool         `json:"isNullSpray"`
	HideIfNotOwned      bool         `json:"hideIfNotOwned"`
	DisplayIcon         string       `json:"displayIcon"`
	FullIcon            *string      `json:"fullIcon"`
	FullTransparentIcon *string      `json:"fullTransparentIcon"`
	AnimationPng        *string      `json:"animationPng"`
	AnimationGif        *string      `json:"animationGif"`
	AssetPath           string       `json:"assetPath"`
	Levels              []SprayLevel `json:"levels"`
}

type SprayLevel struct {
	UUID        string  `json:"uuid"`
	SprayLevel  int     `json:"sprayLevel"`
	DisplayName string  `json:"displayName"`
	DisplayIcon *string `json:"displayIcon"`
	AssetPath   string  `json:"assetPath"`
}

func (s *SpraysService) List(ctx context.Context, opts *RequestOptions) ([]*Spray, *http.Response, error) {
	u := "sprays"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var sprays struct {
		Data []*Spray `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &sprays)
	if err != nil {
		return nil, resp, err
	}

	return sprays.Data, resp, nil
}

func (s *SpraysService) ListLevels(ctx context.Context, opts *RequestOptions) ([]*SprayLevel, *http.Response, error) {
	u := "sprays/levels"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var sprayLevels struct {
		Data []*SprayLevel `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &sprayLevels)
	if err != nil {
		return nil, resp, err
	}

	return sprayLevels.Data, resp, nil
}

func (s *SpraysService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Spray, *http.Response, error) {
	u := fmt.Sprintf("sprays/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var spray struct {
		Data *Spray `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &spray)
	if err != nil {
		return nil, resp, err
	}

	return spray.Data, resp, nil
}

func (s *SpraysService) GetLevel(ctx context.Context, uuid string, opts *RequestOptions) (*SprayLevel, *http.Response, error) {
	u := fmt.Sprintf("sprays/levels/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var sprayLevel struct {
		Data *SprayLevel `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &sprayLevel)
	if err != nil {
		return nil, resp, err
	}

	return sprayLevel.Data, resp, nil
}
