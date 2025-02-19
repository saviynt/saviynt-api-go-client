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

This is an auto-generated client using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) `7.10.0`.

## OpenAPI Spec References

The follow API Reference are generated using Redocly and the OpenAPI specs used to build this client SDK.

1. [Delegated Administration](https://saviynt.github.io/saviynt-api-go-client/api_delegatedadministration.html)
1. [File Directory](https://saviynt.github.io/saviynt-api-go-client/api_filedirectory.html)
1. [Job Control](https://saviynt.github.io/saviynt-api-go-client/api_jobcontrol.html)
1. [mTLS Authentication](https://saviynt.github.io/saviynt-api-go-client/api_mtlsauthentication.html)
1. [SAV Roles](https://saviynt.github.io/saviynt-api-go-client/api_savroles.html)
1. [Tasks](https://saviynt.github.io/saviynt-api-go-client/api_tasks.html)
1. [Transport](https://saviynt.github.io/saviynt-api-go-client/api_transport.html)
1. [Users](https://saviynt.github.io/saviynt-api-go-client/api_users.html)

## Automated Tests

Tests are run with real credentials with the following environment variable:

`SAVIYNT_TEST_CREDENTIALS={"serverURL":"https://myidentitycloud.com","username":"myusername","password":"mypassword"}`

## Supported Saviynt APIs

| No. | Tag | Name | Endpoint | In Spec | In SDK | SDK Test: Manual | SDK Test: Automated |
| - | - | - | - | - | - | - | - |
| 1 | Delegated Administration | Get Delegate User List | `GET /ECM/api/v5/getDelegateUserList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 2 | Delegated Administration | Create Delegate | `POST /ECM/api/v5/createDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 3 | Delegated Administration | Fetch Delegates List | `POST /ECM/api/v5/fetchDelegatesList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 4 | Delegated Administration | Edit Delegate | `POST /ECM/api/v5/editDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 5 | Delegated Administration | Delete Delegate | `POST /ECM/api/v5/deleteDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 6 | Email | Send Email | `POST /ECM/api/v5/sendEmail` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 7 | File Directory | Upload New File | `POST /ECM/api/v5/uploadSchemaFile` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 8 | Job Control | Create and Update Trigger | `POST /ECM/api/v5/createUpdateTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 9 | Job Control | Check Job Status | `POST /ECM/api/v5/checkJobStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 10 | Job Control | Delete Trigger | `POST /ECM/api/v5/deleteTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 11 | Job Control | Run Job Trigger | `POST /ECM/api/v5/runJobTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 12 | Job Control | Fetch Job Metadata | `POST /ECM/api/v5/fetchJobMetadata` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 13 | Job Control | Create Triggers | `POST /ECM/api/v5/createTriggers` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 14 | Job Control | Resume Pause Jobs | `POST /ECM/api/v5/resumePauseJobs` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 15 | mTLS Authentication | Upload KeyStore | `POST /ECM/api/v5/uploadKeyStore` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 16 | mTLS Authentication | Get KeyStore Details | `POST /ECM/api/v5/getKeyStoreCertificateDetails` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 17 | mTLS Authentication | Delete KeyStore | `POST /ECM/api/v5/deleteKeyStoreAlias/{alias}` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 18 | SAV Roles | Get All SAV Roles | `POST /ECMv6/api/userms/savroles` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 19 | SAV Roles | Get SAV Role Users | `POST /ECMv6/api/userms/savroles/{roleName}/users` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 20 | Tasks | Check Task Status | `POST /ECM/api/v5/checkTaskStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 21 | Tasks | Update Tasks | `POST /ECM/api/v5/updateTasks` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 22 | Transport | Export Package | `POST /ECM/api/v5/exportTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 23 | Transport | Import Package | `POST /ECM/api/v5/importTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 24 | Transport | Transport Status | `GET /ECM/api/v5/transportPackageStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 25 | Users | Get User Details | `POST /ECM/api/v5/getUser` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
