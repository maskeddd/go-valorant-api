package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type GamemodesService service

type Gamemode struct {
	UUID                  string                 `json:"uuid"`
	DisplayName           string                 `json:"displayName"`
	Description           *string                `json:"description"`
	Duration              *string                `json:"duration"`
	EconomyType           *string                `json:"economyType"`
	AllowsMatchTimeouts   bool                   `json:"allowsMatchTimeouts"`
	IsTeamVoiceAllowed    bool                   `json:"isTeamVoiceAllowed"`
	IsMinimapHidden       bool                   `json:"isMinimapHidden"`
	OrbCount              int                    `json:"orbCount"`
	RoundsPerHalf         int                    `json:"roundsPerHalf"`
	TeamRoles             []string               `json:"teamRoles"`
	GameFeatureOverrides  []GameFeatureOverride  `json:"gameFeatureOverrides"`
	GameRuleBoolOverrides []GameRuleBoolOverride `json:"gameRuleBoolOverrides"`
	DisplayIcon           *string                `json:"displayIcon"`
	ListViewIconTall      *string                `json:"listViewIconTall"`
	AssetPath             string                 `json:"assetPath"`
}

type GameFeatureOverride struct {
	FeatureName string `json:"featureName"`
	State       bool   `json:"state"`
}

type GameRuleBoolOverride struct {
	RuleName string `json:"ruleName"`
	State    bool   `json:"state"`
}

type GamemodeEquippable struct {
	UUID           string `json:"uuid"`
	DisplayName    string `json:"displayName"`
	Category       string `json:"category"`
	DisplayIcon    string `json:"displayIcon"`
	KillStreamIcon string `json:"killStreamIcon"`
	AssetPath      string `json:"assetPath"`
}

func (s *GamemodesService) List(ctx context.Context, opts *RequestOptions) ([]*Gamemode, *http.Response, error) {
	u := "gamemodes"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var gamemodes struct {
		Data []*Gamemode `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &gamemodes)
	if err != nil {
		return nil, resp, err
	}

	return gamemodes.Data, resp, nil
}

func (s *GamemodesService) ListEquipabbles(ctx context.Context, opts *RequestOptions) ([]*GamemodeEquippable, *http.Response, error) {
	u := "gamemodes/equippables"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var gamemodes struct {
		Data []*GamemodeEquippable `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &gamemodes)
	if err != nil {
		return nil, resp, err
	}

	return gamemodes.Data, resp, nil
}

func (s *GamemodesService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Gamemode, *http.Response, error) {
	u := fmt.Sprintf("gamemodes/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var gamemode struct {
		Data *Gamemode `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &gamemode)
	if err != nil {
		return nil, resp, err
	}

	return gamemode.Data, resp, nil
}

func (s *GamemodesService) GetEquipabble(ctx context.Context, uuid string, opts *RequestOptions) (*GamemodeEquippable, *http.Response, error) {
	u := fmt.Sprintf("gamemodes/equippables/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var gamemode struct {
		Data *GamemodeEquippable `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &gamemode)
	if err != nil {
		return nil, resp, err
	}

	return gamemode.Data, resp, nil
}
