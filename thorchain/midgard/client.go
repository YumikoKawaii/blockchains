package midgard

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client interface {
	GetPoolsList(ctx context.Context, request *GetPoolsListRequest) (*GetPoolsListResponse, error)
}

type clientImpl struct {
	httpClient *http.Client
}

func NewClient() Client {
	return &clientImpl{
		httpClient: &http.Client{},
	}
}

func (c *clientImpl) GetPoolsList(ctx context.Context, request *GetPoolsListRequest) (*GetPoolsListResponse, error) {

	// build query
	query := url.Values{}
	if request != nil {
		if len(request.Status) != 0 {
			query.Add("status", request.Status)
		}

		if len(request.Period) != 0 {
			query.Add("period", request.Period)
		}
	}

	reqUrl := fmt.Sprintf("%s%s?%s", BaseURL, GetPoolsListEndpoint, fmt.Sprintf("&%s", query.Encode()))
	resp, err := c.httpClient.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var pools []Pool
	if err := json.NewDecoder(resp.Body).Decode(&pools); err != nil {
		return nil, err
	}

	return &GetPoolsListResponse{
		Pools: pools,
	}, nil
}
