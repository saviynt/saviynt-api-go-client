# Saviynt API Go Client

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/saviynt/saviynt-api-go-client/workflows/test/badge.svg
 [build-status-url]: https://github.com/saviynt/saviynt-api-go-client/actions/workflows/test.yaml
 [lint-status-svg]: https://github.com/saviynt/saviynt-api-go-client/workflows/lint/badge.svg
 [lint-status-url]: https://github.com/saviynt/saviynt-api-go-client/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/saviynt/saviynt-api-go-client
 [goreport-url]: https://goreportcard.com/report/github.com/saviynt/saviynt-api-go-client
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/saviynt/saviynt-api-go-client
 [docs-godoc-url]: https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/saviynt/saviynt-api-go-client/blob/master/LICENSE

## Overview

This is a Client SDK in Go for the [Saviynt API](https://saviynt.com/api-reference).

## OpenAPI Spec References

The follow API Reference are generated using Redocly and the OpenAPI specs used to build this client SDK.

1. [Connections](https://saviynt.github.io/saviynt-api-go-client/api_connections.html)
1. [Delegated Administration](https://saviynt.github.io/saviynt-api-go-client/api_delegatedadministration.html)
1. [File Directory](https://saviynt.github.io/saviynt-api-go-client/api_filedirectory.html)
1. [Job Control](https://saviynt.github.io/saviynt-api-go-client/api_jobcontrol.html)
1. [mTLS Authentication](https://saviynt.github.io/saviynt-api-go-client/api_mtlsauthentication.html)
1. [SAV Roles](https://saviynt.github.io/saviynt-api-go-client/api_savroles.html)
1. [Tasks](https://saviynt.github.io/saviynt-api-go-client/api_tasks.html)
1. [Transport](https://saviynt.github.io/saviynt-api-go-client/api_transport.html)
1. [Users](https://saviynt.github.io/saviynt-api-go-client/api_users.html)

## Supported Saviynt APIs

| No. | Tag | Name | Endpoint | In Spec | In SDK | SDK Test: Manual | SDK Test: Automated |
| - | - | - | - | - | - | - | - |
| 1 | Connections | Get List of Connections | `POST /ECM/api/v5/getConnections` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 2 | Connections | Get Connection Details | `POST /ECM/api/v5/getConnectionDetails` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 3 | Connections | Create or Update Connection | `POST /ECM/api/v5/testConnection` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 4 | Connections | Delete Connection | `DELETE /ECM/api/v5/deleteConnection` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 5 | Delegated Administration | Get Delegate User List | `GET /ECM/api/v5/getDelegateUserList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 6 | Delegated Administration | Create Delegate | `POST /ECM/api/v5/createDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 7 | Delegated Administration | Fetch Delegates List | `POST /ECM/api/v5/fetchDelegatesList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 8 | Delegated Administration | Edit Delegate | `POST /ECM/api/v5/editDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 9 | Delegated Administration | Delete Delegate | `POST /ECM/api/v5/deleteDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 10 | Email | Send Email | `POST /ECM/api/v5/sendEmail` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 11 | File Directory | Upload New File | `POST /ECM/api/v5/uploadSchemaFile` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 12 | Job Control | Create and Update Trigger | `POST /ECM/api/v5/createUpdateTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 13 | Job Control | Check Job Status | `POST /ECM/api/v5/checkJobStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 14 | Job Control | Delete Trigger | `POST /ECM/api/v5/deleteTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 15 | Job Control | Run Job Trigger | `POST /ECM/api/v5/runJobTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 16 | Job Control | Fetch Job Metadata | `POST /ECM/api/v5/fetchJobMetadata` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 17 | Job Control | Create Triggers | `POST /ECM/api/v5/createTriggers` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 18 | Job Control | Resume Pause Jobs | `POST /ECM/api/v5/resumePauseJobs` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 19 | mTLS Authentication | Upload KeyStore | `POST /ECM/api/v5/uploadKeyStore` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 20 | mTLS Authentication | Get KeyStore Details | `POST /ECM/api/v5/getKeyStoreCertificateDetails` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 21 | mTLS Authentication | Delete KeyStore | `POST /ECM/api/v5/deleteKeyStoreAlias/{alias}` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 22 | SAV Roles | Get All SAV Roles | `POST /ECMv6/api/userms/savroles` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 23 | SAV Roles | Get SAV Role Users | `POST /ECMv6/api/userms/savroles/{roleName}/users` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 24 | Tasks | Check Task Status | `POST /ECM/api/v5/checkTaskStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 25 | Tasks | Update Tasks | `POST /ECM/api/v5/updateTasks` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 26 | Transport | Export Package | `POST /ECM/api/v5/exportTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 27 | Transport | Import Package | `POST /ECM/api/v5/importTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 28 | Transport | Transport Status | `GET /ECM/api/v5/transportPackageStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 29 | Users | Get User Details | `POST /ECM/api/v5/getUser` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |

## Automated Tests

Tests are run with real credentials with the following environment variable:

`SAVIYNT_TEST_CREDENTIALS={"serverURL":"https://myidentitycloud.com","username":"myusername","password":"mypassword"}`

To run the tests, you can execute the following:

```
% git clone https://github.com/saviynt/saviynt-api-go-client
% cd saviynt-api-go-client
% go mod tidy
% export SAVIYNT_TEST_CREDENTIALS='{...}'
% cd test
% go test -v ./...
% go test -run Test_delegatedadministration_DelegatedAdministrationAPIService
```

Use the following to run the tests:

1. Run all tests: `go test -v ./...`
2. Run a specific suite of tests, e.g. for `delegatedadministration`: `go test -run Test_delegatedadministration_DelegatedAdministrationAPIService`

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. The best ways to contribute to this SDK are:

1. API coverage: to add/update API coverage, update the OpenAPI spec files in the [`docs`](docs) folder.
2. SDK functionality: if SDK functionality is desired that is not covered by [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator), create an isue or PR with that project.
3. Test coverage: tests are always welcome, in the [`tests`](tests) folder.

Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release.

## Credits

This is an auto-generated client using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) `7.10.0`.