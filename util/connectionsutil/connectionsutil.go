package connectionsutil

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/saviynt/saviynt-api-go-client/connections"
	"github.com/saviynt/saviynt-api-go-client/util/jsonutil"
)

func ReadFileCreateOrUpdateConnectionRequest(filename string) (*connections.CreateOrUpdateConnectionRequest, error) {
	req := &connections.CreateOrUpdateConnectionRequest{}
	if b, err := os.ReadFile(filename); err != nil {
		return nil, err
	} else if err = json.Unmarshal(b, req); err != nil {
		return nil, err
	} else {
		return req, nil
	}
}

func ReadFileConnectionDetails(filename string) (*connections.GetConnectionDetailsResponse, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &connections.GetConnectionDetailsResponse{}
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ConnectionDetailsResponseToRequest(r connections.GetConnectionDetailsResponse) (connections.CreateOrUpdateConnectionRequest, error) {
	w := connections.CreateOrUpdateConnectionRequest{}

	reqFields := jsonutil.ParseFields(w)
	reqFieldsString := reqFields.JSONTagNames(true)
	msa := map[string]any{}
	if r.ConnectionType != nil {
		msa["connectiontype"] = *r.ConnectionType
	}
	if r.Connectionkey != nil {
		msa["connectionkey"] = strconv.Itoa(int(*r.Connectionkey))
	}

	if b, err := json.Marshal(r.Connectionattributes); err != nil {
		return w, err
	} else if b, err = jsonutil.JSONFilterKeys(b, reqFieldsString, msa); err != nil {
		return w, err
	} else if err = json.Unmarshal(b, &w); err != nil {
		return w, err
	} else {
		return w, nil
	}
}
