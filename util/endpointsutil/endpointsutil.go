package endpointsutil

import (
	"context"
	"errors"
	"fmt"

	saviyntapigoclient "github.com/saviynt/saviynt-api-go-client"
	"github.com/saviynt/saviynt-api-go-client/endpoints"
)

const (
	ParamMaxMaximum = int32(500)
)

type Endpoints []endpoints.Endpoint

func getEndpointsMulti(ctx context.Context, api *endpoints.EndpointsAPIService, endpointname string) (Endpoints, error) {
	eps := Endpoints{}
	if api == nil {
		return eps, errors.New("api client cannot be nil")
	}

	req := endpoints.GetEndpointsRequest{
		Max:    saviyntapigoclient.Pointer(ParamMaxMaximum),
		Offset: saviyntapigoclient.Pointer(int32(0))}

	for {
		resp, httpRes, err := api.
			GetEndpoints(ctx).
			GetEndpointsRequest(req).
			Execute()
		if err != nil {
			return eps, err
		} else if httpRes == nil {
			return eps, errors.New("received nil `*http.Response` value")
		} else if httpRes.StatusCode >= 300 {
			return eps, fmt.Errorf("received bad status code from API (%d)", httpRes.StatusCode)
		} else if resp == nil {
			continue
		}
		if endpointname != "" {
			for _, ep := range resp.Endpoints {
				if ep.Endpointname == endpointname {
					return Endpoints{ep}, nil
				}
			}
		} else {
			eps = append(eps, resp.Endpoints...)
		}
		if resp.TotalCount == nil {
			break
		}
		req.Offset = saviyntapigoclient.Pointer(*req.Offset + ParamMaxMaximum)
		if *req.Offset >= *resp.TotalCount {
			break
		}
	}
	if endpointname != "" {
		return Endpoints{}, fmt.Errorf("endpoint for endpoint name not found (%s)", endpointname)
	} else {
		return eps, nil
	}
}

func GetAllEndpoints(ctx context.Context, api *endpoints.EndpointsAPIService) (Endpoints, error) {
	return getEndpointsMulti(ctx, api, "")
}

func GetEndpoint(ctx context.Context, api *endpoints.EndpointsAPIService, endpointname string) (*endpoints.Endpoint, error) {
	if eps, err := getEndpointsMulti(ctx, api, endpointname); err != nil {
		return nil, err
	} else if len(eps) == 0 {
		return nil, fmt.Errorf("endpoint for endpoint name not found (%s)", endpointname)
	} else if len(eps) == 1 {
		return saviyntapigoclient.Pointer(eps[0]), nil
	} else {
		return nil, fmt.Errorf("endpoint for endpoint name resulted in multiple matches (%s)", endpointname)
	}
}
