module github.com/saviynt/saviynt-api-go-client

go 1.24.0

require (
	github.com/stretchr/testify v1.10.0
	golang.org/x/oauth2 v0.28.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/saviynt/saviynt-api-go-client/entitlements v0.0.0
	github.com/saviynt/saviynt-api-go-client/organization v0.0.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/saviynt/saviynt-api-go-client/entitlements => ./entitlements

replace github.com/saviynt/saviynt-api-go-client/organization => ./organization
