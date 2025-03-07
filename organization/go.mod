module test_organization

go 1.18

require (
	example.com/generated-client v0.0.0 // You can set a dummy version here
	github.com/stretchr/testify v1.10.0
)

replace example.com/generated-client => ../organization

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
