package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type ContractsService service

type Contract struct {
	UUID                   string  `json:"uuid"`
	DisplayName            string  `json:"displayName"`
	DisplayIcon            *string `json:"displayIcon"`
	ShipIt                 bool    `json:"shipIt"`
	UseLevelVPCostOverride bool    `json:"useLevelVPCostOverride"`
	LevelVPCostOverride    int     `json:"levelVPCostOverride"`
	FreeRewardScheduleUUID string  `json:"freeRewardScheduleUuid"`
	Content                Content `json:"content"`
	AssetPath              string  `json:"assetPath"`
}

type Content struct {
	RelationType              *string   `json:"relationType"`
	RelationUUID              *string   `json:"relationUuid"`
	Chapters                  []Chapter `json:"chapters"`
	PremiumRewardScheduleUUID *string   `json:"premiumRewardScheduleUuid"`
	PremiumVPCost             int       `json:"premiumVPCost"`
}

type Chapter struct {
	IsEpilogue  bool            `json:"isEpilogue"`
	Levels      []ContractLevel `json:"levels"`
	FreeRewards []ContractLevel `json:"freeRewards"`
}

type ContractLevel struct {
	Reward                 Reward `json:"reward"`
	XP                     int    `json:"xp"`
	VPCost                 int    `json:"vpCost"`
	IsPurchasableWithVP    bool   `json:"isPurchasableWithVP"`
	DoughCost              int    `json:"doughCost"`
	IsPurchasableWithDough bool   `json:"isPurchasableWithDough"`
}
type Reward struct {
	Type          string `json:"type"`
	UUID          string `json:"uuid"`
	Amount        int    `json:"amount"`
	IsHighlighted bool   `json:"isHighlighted"`
}

func (s *ContractsService) List(ctx context.Context, opts *RequestOptions) ([]*Contract, *http.Response, error) {
	u := "contracts"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var contracts struct {
		Data []*Contract `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &contracts)
	if err != nil {
		return nil, resp, err
	}

	return contracts.Data, resp, nil
}

func (s *ContractsService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Contract, *http.Response, error) {
	u := fmt.Sprintf("contracts/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var contract struct {
		Data *Contract `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &contract)
	if err != nil {
		return nil, resp, err
	}

	return contract.Data, resp, nil
}
