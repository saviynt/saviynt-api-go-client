# Go API client for connections

Connections

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import connections "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `connections.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), connections.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `connections.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), connections.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `connections.ContextOperationServerIndices` and `connections.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), connections.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), connections.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConnectionsAPI* | [**CreateOrUpdateConnection**](docs/ConnectionsAPI.md#createorupdateconnection) | **Post** /ECM/api/v5/testConnection | Create or Update Connection
*ConnectionsAPI* | [**DeleteConnection**](docs/ConnectionsAPI.md#deleteconnection) | **Delete** /ECM/api/v5/deleteConnection | Delete Connection
*ConnectionsAPI* | [**GetConnectionDetails**](docs/ConnectionsAPI.md#getconnectiondetails) | **Post** /ECM/api/v5/getConnectionDetails | Get Connection Details
*ConnectionsAPI* | [**GetConnections**](docs/ConnectionsAPI.md#getconnections) | **Post** /ECM/api/v5/getConnections | Get List of Connections


## Documentation For Models

 - [Connection](docs/Connection.md)
 - [ConnectionAttributes](docs/ConnectionAttributes.md)
 - [ConnectionTimeoutConfig](docs/ConnectionTimeoutConfig.md)
 - [CreateOrUpdateConnectionRequest](docs/CreateOrUpdateConnectionRequest.md)
 - [CreateOrUpdateConnectionResponse](docs/CreateOrUpdateConnectionResponse.md)
 - [DeleteConnectionRequest](docs/DeleteConnectionRequest.md)
 - [DeleteConnectionResponse](docs/DeleteConnectionResponse.md)
 - [GetConnectionDetailsRequest](docs/GetConnectionDetailsRequest.md)
 - [GetConnectionDetailsResponse](docs/GetConnectionDetailsResponse.md)
 - [GetConnectionsRequest](docs/GetConnectionsRequest.md)
 - [GetConnectionsResponse](docs/GetConnectionsResponse.md)
 - [VaultConfiguration](docs/VaultConfiguration.md)
 - [VaultConfigurationKeyMapping](docs/VaultConfigurationKeyMapping.md)
 - [VaultConfigurationKeyMappingPassword](docs/VaultConfigurationKeyMappingPassword.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

https://github.com/saviynt

