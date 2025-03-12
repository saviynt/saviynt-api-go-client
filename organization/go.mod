module test_organization

go 1.18

require (
	github.com/saviynt/saviynt-api-go-client/organization v0.0.0
	github.com/stretchr/testify v1.10.0
)

replace github.com/saviynt/saviynt-api-go-client/organization => ../organization

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
