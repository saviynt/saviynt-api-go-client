package saviyntapigoclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/grokify/saviynt-api-go-client/delegatedadministration"
	"github.com/grokify/saviynt-api-go-client/filedirectory"
	"github.com/grokify/saviynt-api-go-client/jobcontrol"
	"github.com/grokify/saviynt-api-go-client/mtlsauthentication"
	"github.com/grokify/saviynt-api-go-client/savroles"
	"github.com/grokify/saviynt-api-go-client/transport"
	"golang.org/x/oauth2"
)

const (
	apiLoginPath         = "/ECM/api/login"
	headerAuthorization  = "Authorization"
	tokenTypeBasic       = "Basic"
	paramOAuth2ExpiresIn = "expires_in"
)

type Client struct {
	serverURL                     string
	token                         *oauth2.Token
	httpClient                    *http.Client
	DelegatedAdministrationAPI    *delegatedadministration.DelegatedAdministrationAPIService
	delegatedAdministrationClient *delegatedadministration.APIClient
	FileDirectoryAPI              *filedirectory.FileDirectoryAPIService
	fileDirectoryClient           *filedirectory.APIClient
	JobControlAPI                 *jobcontrol.JobControlAPIService
	jobControlClient              *jobcontrol.APIClient
	MTLSAuthenticationAPI         *mtlsauthentication.MTLSAuthenticationAPIService
	mtlsAuthenticationClient      *mtlsauthentication.APIClient
	SAVRolesAPI                   *savroles.SAVRolesAPIService
	savRolesClient                *savroles.APIClient
	TransportAPI                  *transport.TransportAPIService
	transportClient               *transport.APIClient
}

func NewClient(ctx context.Context, serverURL string, httpClient *http.Client) *Client {
	c := &Client{
		serverURL:  serverURL,
		httpClient: httpClient}
	c.delegatedAdministrationClient = newClientDelegatedAdministration(c.APIBaseURL(), c.httpClient)
	c.DelegatedAdministrationAPI = c.delegatedAdministrationClient.DelegatedAdministrationAPI
	c.fileDirectoryClient = newClientFileDirectory(c.APIBaseURL(), c.httpClient)
	c.FileDirectoryAPI = c.fileDirectoryClient.FileDirectoryAPI
	c.jobControlClient = newClientJobControl(c.APIBaseURL(), c.httpClient)
	c.JobControlAPI = c.jobControlClient.JobControlAPI
	c.mtlsAuthenticationClient = newClientMTLSAuthentication(c.APIBaseURL(), c.httpClient)
	c.MTLSAuthenticationAPI = c.mtlsAuthenticationClient.MTLSAuthenticationAPI
	c.savRolesClient = newClientSAVRoles(c.APIBaseURL(), c.httpClient)
	c.SAVRolesAPI = c.savRolesClient.SAVRolesAPI
	c.transportClient = newClientTransport(c.APIBaseURL(), c.httpClient)
	c.TransportAPI = c.transportClient.TransportAPI
	return c
}

type Credentials struct {
	ServerURL string `json:"serverURL"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func NewClientPassword(ctx context.Context, creds Credentials) (*Client, error) {
	if tok, err := newOAuth2TokenBasicAuth(loginURL(creds.ServerURL), creds.Username, creds.Password); err != nil {
		return nil, err
	} else {
		c := NewClientToken(ctx, creds.ServerURL, tok)
		c.token = tok
		return c, nil
	}
}

func NewClientPasswordEnv(ctx context.Context, envvar string) (*Client, error) {
	v := os.Getenv(envvar)
	if v == "" {
		return nil, errors.New("env var cannot be empty")
	} else if strings.Index(strings.TrimSpace(v), "{") != 0 {
		return nil, errors.New("env var must be a json object")
	}
	creds := Credentials{}
	if err := json.Unmarshal([]byte(v), &creds); err != nil {
		return nil, err
	}
	return NewClientPassword(ctx, creds)
}

func NewClientToken(ctx context.Context, serverURL string, token *oauth2.Token) *Client {
	c := NewClient(ctx, serverURL, newClientToken(ctx, token))
	c.token = token
	return c
}

func (c *Client) APIBaseURL() string {
	return strings.TrimRight(strings.TrimSpace(c.serverURL), "/")
}

func (c *Client) Token() *oauth2.Token {
	return c.token
}

func newClientDelegatedAdministration(apiBaseURL string, httpClient *http.Client) *delegatedadministration.APIClient {
	cfg := delegatedadministration.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = delegatedadministration.ServerConfigurations{{URL: apiBaseURL}}
	return delegatedadministration.NewAPIClient(cfg)
}

func newClientFileDirectory(apiBaseURL string, httpClient *http.Client) *filedirectory.APIClient {
	cfg := filedirectory.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = filedirectory.ServerConfigurations{{URL: apiBaseURL}}
	return filedirectory.NewAPIClient(cfg)
}

func newClientJobControl(apiBaseURL string, httpClient *http.Client) *jobcontrol.APIClient {
	cfg := jobcontrol.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = jobcontrol.ServerConfigurations{{URL: apiBaseURL}}
	return jobcontrol.NewAPIClient(cfg)
}

func newClientMTLSAuthentication(apiBaseURL string, httpClient *http.Client) *mtlsauthentication.APIClient {
	cfg := mtlsauthentication.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = mtlsauthentication.ServerConfigurations{{URL: apiBaseURL}}
	return mtlsauthentication.NewAPIClient(cfg)
}

func newClientSAVRoles(apiBaseURL string, httpClient *http.Client) *savroles.APIClient {
	cfg := savroles.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = savroles.ServerConfigurations{{URL: apiBaseURL}}
	return savroles.NewAPIClient(cfg)
}

func newClientTransport(apiBaseURL string, httpClient *http.Client) *transport.APIClient {
	cfg := transport.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = transport.ServerConfigurations{{URL: apiBaseURL}}
	return transport.NewAPIClient(cfg)
}

func newOAuth2TokenBasicAuth(tokenURL, username, password string) (*oauth2.Token, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, tokenURL, nil)
	if err != nil {
		return nil, err
	}
	header := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Add(headerAuthorization, tokenTypeBasic+" "+header)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return parseToken(b)
}

func loginURL(serverURL string) string {
	return strings.TrimSuffix(strings.TrimSpace(serverURL), "/") + apiLoginPath
}

func parseToken(rawToken []byte) (*oauth2.Token, error) {
	tok := &oauth2.Token{}
	err := json.Unmarshal(rawToken, tok)
	if err != nil {
		return tok, err
	}
	msi := map[string]any{}
	err = json.Unmarshal(rawToken, &msi)
	if err != nil {
		return tok, err
	}
	tok = tok.WithExtra(msi)
	// convert `expires_in` to `Expiry` with 1 minute leeway.
	if tok.Expiry.IsZero() {
		expiresIn := tok.Extra(paramOAuth2ExpiresIn)
		if expiresIn != nil {
			if expiresInFloat, ok := expiresIn.(float64); ok {
				if expiresInFloat > 60 { // subtract 1 minute for defensive handling
					expiresInFloat -= 60
				}
				if expiresInFloat > 0 {
					tok.Expiry = time.Now().UTC().Add(time.Duration(expiresInFloat) * time.Second)
				}
			}
		}
	}
	return tok, nil
}

func newClientToken(ctx context.Context, tok *oauth2.Token) *http.Client {
	oAuthConfig := &oauth2.Config{}
	return oAuthConfig.Client(ctx, tok)
}

func Pointer[E any](e E) *E { return &e }
