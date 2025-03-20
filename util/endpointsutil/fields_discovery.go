package endpointsutil

import (
	"encoding/json"
	"io"
	"sort"
)

type GetEndpointsResponseFieldDiscovery struct {
	Endpoints []map[string]any `json:"endpoints"`
	ErrorCode string           `json:"errorCode"`
	Message   string           `json:"message"`
}

func ParseGetEndpointsFieldsDiscovery(r io.Reader) ([]string, error) {
	resp := GetEndpointsResponseFieldDiscovery{}
	if b, err := io.ReadAll(r); err != nil {
		return []string{}, err
	} else if err := json.Unmarshal(b, &resp); err != nil {
		return []string{}, err
	} else {
		return resp.EndpointFields(), nil
	}
}

func (resp GetEndpointsResponseFieldDiscovery) EndpointFields() []string {
	var fields []string
	seen := map[string]int{}
	for _, e := range resp.Endpoints {
		for k := range e {
			if _, ok := seen[k]; ok {
				continue
			} else {
				fields = append(fields, k)
				seen[k]++
			}
		}
	}
	sort.Strings(fields)
	return fields
}
