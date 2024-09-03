package valorantapi

import (
	"context"
	"net/http"
	"time"
)

type VersionService service

type Version struct {
	ManifestID        string    `json:"manifestId"`
	Branch            string    `json:"branch"`
	Version           string    `json:"version"`
	BuildVersion      string    `json:"buildVersion"`
	EngineVersion     string    `json:"engineVersion"`
	RiotClientVersion string    `json:"riotClientVersion"`
	RiotClientBuild   string    `json:"riotClientBuild"`
	BuildDate         time.Time `json:"buildDate"`
}

func (s *VersionService) Get(ctx context.Context) (*Version, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "version", nil)
	if err != nil {
		return nil, nil, err
	}

	var version struct {
		Data *Version `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &version)
	if err != nil {
		return nil, resp, err
	}

	return version.Data, resp, nil
}
