package bpx_api_sdk_go

import (
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"github.com/imroc/req/v3"
)

type system struct {
	url    string
	client *req.Client
}

func NewSystem(url string, client *req.Client) *system {
	return &system{url: url, client: client}
}

func (s *system) GetStatus() (*types.SystemStatus, error) {
	response, err := s.client.R().Get(fmt.Sprintf("%s/api/v1/status", s.url))
	if err != nil {
		return nil, err
	}

	result, err := dealResponse(new(types.SystemStatus), response, err)
	if err != nil {
		return nil, err
	}
	fmtRes := result.(types.SystemStatus)
	return &fmtRes, nil
}

func (s *system) Ping() error {
	response, err := s.client.R().Get(fmt.Sprintf("%s/api/v1/ping", s.url))
	if err != nil {
		return err
	}

	content := response.String()
	fmt.Printf("content: %s \n", content)
	return nil
}

func (s *system) GetTime() error {
	response, err := s.client.R().Get(fmt.Sprintf("%s/api/v1/time", s.url))
	if err != nil {
		return err
	}
	content := response.String()
	fmt.Printf("content: %s \n", content)
	return nil
}
