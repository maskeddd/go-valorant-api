package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type MapsService service

type Map struct {
	UUID                    string      `json:"uuid"`
	DisplayName             string      `json:"displayName"`
	NarrativeDescription    interface{} `json:"narrativeDescription"`
	TacticalDescription     *string     `json:"tacticalDescription"`
	Coordinates             *string     `json:"coordinates"`
	DisplayIcon             *string     `json:"displayIcon"`
	ListViewIcon            string      `json:"listViewIcon"`
	ListViewIconTall        *string     `json:"listViewIconTall"`
	Splash                  string      `json:"splash"`
	StylizedBackgroundImage *string     `json:"stylizedBackgroundImage"`
	PremierBackgroundImage  *string     `json:"premierBackgroundImage"`
	AssetPath               string      `json:"assetPath"`
	MapURL                  string      `json:"mapUrl"`
	XMultiplier             float64     `json:"xMultiplier"`
	YMultiplier             float64     `json:"yMultiplier"`
	XScalarToAdd            float64     `json:"xScalarToAdd"`
	YScalarToAdd            float64     `json:"yScalarToAdd"`
	Callouts                []Callout   `json:"callouts"`
}

type Callout struct {
	RegionName      string   `json:"regionName"`
	SuperRegionName string   `json:"superRegionName"`
	Location        Location `json:"location"`
}

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (s *MapsService) List(ctx context.Context, opts *RequestOptions) ([]*Map, *http.Response, error) {
	u := "maps"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var maps struct {
		Data []*Map `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &maps)
	if err != nil {
		return nil, resp, err
	}

	return maps.Data, resp, nil
}

func (s *MapsService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Map, *http.Response, error) {
	u := fmt.Sprintf("maps/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var mapData struct {
		Data *Map `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &mapData)
	if err != nil {
		return nil, resp, err
	}

	return mapData.Data, resp, nil
}
