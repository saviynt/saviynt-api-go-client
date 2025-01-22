package test

import (
	"context"
	"errors"
	"os"
	"strings"

	saviyntapigoclient "github.com/grokify/saviynt-api-go-client"
)

const EnvSaviyntTestCredentials = "SAVIYNT_TEST_CREDENTIALS" // #nosec G101

func client() (*saviyntapigoclient.Client, saviyntapigoclient.Credentials, bool, error) {
	creds := saviyntapigoclient.Credentials{}
	if v := strings.TrimSpace(os.Getenv(EnvSaviyntTestCredentials)); v == "" {
		return nil, creds, false, nil
	} else if clt, creds, err := saviyntapigoclient.NewClientPasswordEnv(context.Background(),
		EnvSaviyntTestCredentials); err != nil {
		return nil, creds, false, err
	} else if clt == nil {
		return nil, creds, false, errors.New("client not generated")
	} else {
		return clt, creds, true, nil
	}
}
