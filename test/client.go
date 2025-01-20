package test

import (
	"context"
	"errors"
	"os"
	"strings"

	saviyntapigoclient "github.com/grokify/saviynt-api-go-client"
)

const EnvSaviyntTestCredentials = "SAVIYNT_TEST_CREDENTIALS" // env var

func client() (*saviyntapigoclient.Client, bool, error) {
	v := strings.TrimSpace(os.Getenv(EnvSaviyntTestCredentials))
	if v == "" {
		return nil, false, nil
	}
	clt, err := saviyntapigoclient.NewClientPasswordEnv(context.Background(), EnvSaviyntTestCredentials)
	if err != nil {
		return nil, false, err
	} else if clt == nil {
		return nil, false, errors.New("client not generated")
	} else {
		return clt, true, nil
	}
}
