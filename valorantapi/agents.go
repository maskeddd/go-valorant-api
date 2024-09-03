package valorantapi

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type AgentsService service

type Agent struct {
	UUID                      string           `json:"uuid"`
	DisplayName               string           `json:"displayName"`
	Description               string           `json:"description"`
	DeveloperName             string           `json:"developerName"`
	CharacterTags             *[]string        `json:"characterTags"`
	DisplayIcon               string           `json:"displayIcon"`
	DisplayIconSmall          string           `json:"displayIconSmall"`
	BustPortrait              *string          `json:"bustPortrait"`
	FullPortrait              *string          `json:"fullPortrait"`
	FullPortraitV2            *string          `json:"fullPortraitV2"`
	KillfeedPortrait          string           `json:"killfeedPortrait"`
	Background                *string          `json:"background"`
	BackgroundGradientColors  []string         `json:"backgroundGradientColors"`
	AssetPath                 string           `json:"assetPath"`
	IsFullPortraitRightFacing bool             `json:"isFullPortraitRightFacing"`
	IsPlayableCharacter       bool             `json:"isPlayableCharacter"`
	IsAvailableForTest        bool             `json:"isAvailableForTest"`
	IsBaseContent             bool             `json:"isBaseContent"`
	Role                      *Role            `json:"role"`
	RecruitmentData           *RecruitmentData `json:"recruitmentData"`
	Abilities                 []Ability        `json:"abilities"`
	VoiceLine                 []VoiceLine      `json:"voiceLine"`
}

type Role struct {
	UUID        string  `json:"uuid"`
	DisplayName string  `json:"displayName"`
	Description string  `json:"description"`
	DisplayIcon *string `json:"displayIcon"`
	AssetPath   string  `json:"assetPath"`
}

type RecruitmentData struct {
	CounterID              string    `json:"counterId"`
	MilestoneID            string    `json:"milestoneId"`
	MilestoneThreshold     int32     `json:"milestoneThreshold"`
	UseLevelVpCostOverride bool      `json:"useLevelVpCostOverride"`
	LevelVpCostOverride    int32     `json:"levelVpCostOverride"`
	StartDate              time.Time `json:"startDate"`
	EndDate                time.Time `json:"endDate"`
}

type Ability struct {
	Slot        string `json:"slot"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	DisplayIcon string `json:"displayIcon"`
}

type VoiceLine struct {
	MinDuration float32 `json:"minDuration"`
	MaxDuration float32 `json:"maxDuration"`
	MediaList   []Media `json:"mediaList"`
}

type Media struct {
	ID    int32  `json:"id"`
	Wwise string `json:"wwise"`
	Wave  string `json:"wave"`
}

type AgentsListOptions struct {
	RequestOptions
	IsPlayableCharacter bool `url:"isPlayableCharacter,omitempty"`
}

func (s *AgentsService) List(ctx context.Context, opts *AgentsListOptions) ([]*Agent, *http.Response, error) {
	u := "agents"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var agents struct {
		Data []*Agent `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &agents)
	if err != nil {
		return nil, resp, err
	}

	return agents.Data, resp, nil
}

func (s *AgentsService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Agent, *http.Response, error) {
	u := fmt.Sprintf("agents/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var agent struct {
		Data *Agent `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &agent)
	if err != nil {
		return nil, resp, err
	}

	return agent.Data, resp, nil
}
