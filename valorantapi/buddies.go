package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type BuddiesService service

type Buddy struct {
	UUID               string       `json:"uuid"`
	DisplayName        string       `json:"displayName"`
	IsHiddenIfNotOwned bool         `json:"isHiddenIfNotOwned"`
	ThemeUUID          *string      `json:"themeUuid"`
	DisplayIcon        string       `json:"displayIcon"`
	AssetPath          string       `json:"assetPath"`
	Levels             []BuddyLevel `json:"levels"`
}

type BuddyLevel struct {
	UUID           string `json:"uuid"`
	CharmLevel     int    `json:"charmLevel"`
	HideIfNotOwned bool   `json:"hideIfNotOwned"`
	DisplayName    string `json:"displayName"`
	DisplayIcon    string `json:"displayIcon"`
	AssetPath      string `json:"assetPath"`
}

func (s *BuddiesService) List(context context.Context, opts *RequestOptions) ([]*Buddy, *http.Response, error) {
	u := "buddies"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var buddies struct {
		Data []*Buddy `json:"data"`
	}
	resp, err := s.client.Do(context, req, &buddies)
	if err != nil {
		return nil, resp, err
	}

	return buddies.Data, resp, nil
}

func (s *BuddiesService) ListLevels(context context.Context, opts *RequestOptions) ([]*BuddyLevel, *http.Response, error) {
	u := "buddies/levels"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var levels struct {
		Data []*BuddyLevel `json:"data"`
	}
	resp, err := s.client.Do(context, req, &levels)
	if err != nil {
		return nil, resp, err
	}

	return levels.Data, resp, nil
}

func (s *BuddiesService) Get(context context.Context, uuid string, opts *RequestOptions) (*Buddy, *http.Response, error) {
	u := fmt.Sprintf("buddies/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var buddy struct {
		Data *Buddy `json:"data"`
	}
	resp, err := s.client.Do(context, req, &buddy)
	if err != nil {
		return nil, resp, err
	}

	return buddy.Data, resp, nil
}

func (s *BuddiesService) GetLevel(context context.Context, uuid string, opts *RequestOptions) (*BuddyLevel, *http.Response, error) {
	u := fmt.Sprintf("buddies/levels/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var level struct {
		Data *BuddyLevel `json:"data"`
	}
	resp, err := s.client.Do(context, req, &level)
	if err != nil {
		return nil, resp, err
	}

	return level.Data, resp, nil
}
