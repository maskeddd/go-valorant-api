package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type PlayerCardsService service

type PlayerCard struct {
	UUID               string  `json:"uuid"`
	DisplayName        string  `json:"displayName"`
	IsHiddenIfNotOwned bool    `json:"isHiddenIfNotOwned"`
	ThemeUUID          *string `json:"themeUuid"`
	DisplayIcon        string  `json:"displayIcon"`
	SmallArt           string  `json:"smallArt"`
	WideArt            string  `json:"wideArt"`
	LargeArt           string  `json:"largeArt"`
	AssetPath          string  `json:"assetPath"`
}

func (s *PlayerCardsService) List(ctx context.Context, opts *RequestOptions) ([]*PlayerCard, *http.Response, error) {
	u := "playercards"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var playerCards struct {
		Data []*PlayerCard `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &playerCards)
	if err != nil {
		return nil, resp, err
	}

	return playerCards.Data, resp, nil
}

func (s *PlayerCardsService) Get(ctx context.Context, id string, opts *RequestOptions) (*PlayerCard, *http.Response, error) {
	u := fmt.Sprintf("playercards/%s", id)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var playerCard struct {
		Data *PlayerCard `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, playerCard)
	if err != nil {
		return nil, resp, err
	}

	return playerCard.Data, resp, nil
}
