# Saviynt API Go Client

[![Build Status][build-status-svg]][build-status-link]
[![Lint Status][lint-status-svg]][lint-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://github.com/grokify/saviynt-api-go-client/workflows/test/badge.svg
 [build-status-link]: https://github.com/grokify/saviynt-api-go-client/actions/workflows/test.yaml
 [lint-status-svg]: https://github.com/grokify/saviynt-api-go-client/workflows/lint/badge.svg
 [lint-status-link]: https://github.com/grokify/saviynt-api-go-client/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/saviynt-api-go-client
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/saviynt-api-go-client
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/saviynt-api-go-client
 [docs-godoc-link]: https://pkg.go.dev/github.com/grokify/saviynt-api-go-client
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/saviynt-api-go-client/blob/master/LICENSE

## Overview

This is an auto-generated client using OpenAPI Generator `7.10.0`.

## Supported Saviynt APIs

| Name | Endpoint | Defined | Manual Tests | Automated Tests |
| [Delegated Administration](https://pkg.go.dev/github.com/grokify/saviynt-api-go-client/delegatedadministration) ([API docs](https://grokify.github.io/saviynt-api-go-client/api_delegatedadministration.html)) | | | | | 
| Get Delegate User List | `GET /ECM/api/v5/getDelegateUserList` | [x] | [x] | [x] |

## List

1. [x] [Delegated Administration](https://pkg.go.dev/github.com/grokify/saviynt-api-go-client/delegatedadministration) ([API docs](https://grokify.github.io/saviynt-api-go-client/api_delegatedadministration.html))
  1. [x] [x] `GET /ECM/api/v5/getDelegateUserList` - Get Delegate User List
  1. [x] [x] `POST /ECM/api/v5/createDelegate` - Create Delegate
  1. [x] [x] `POST /ECM/api/v5/fetchDelegatesList` - etch Existing Delegates list
  1. [x] [x] `POST /ECM/api/v5/editDelegate` - Edit Delegate
  1. [x] [x] `POST /ECM/api/v5/deleteDelegate` - Delete Delegate
1. [ ] Job Control
  1. [ ] `POST /ECM/api/v5/createUpdateTrigger` - Create and Update Trigger
  1. [x] `POST /ECM/api/v5/checkJobStatus` - Check Job Status
  1. [ ] `POST /ECM/api/v5/deleteTrigger` - Delete Trigger
  1. [ ] `POST /ECM/api/v5/runJobTrigger` - Run Job Trigger
  1. [ ] `POST /ECM/api/v5/fetchJobMetadata` - Fetch Job Metadata
  1. [ ] `POST /ECM/api/v5/createTriggers` - Create Triggers
  1. [x] `POST /ECM/api/v5/resumePauseJobs` - Resume Pause Jobs
1. [x] [mTLS Authentication](https://pkg.go.dev/github.com/grokify/saviynt-api-go-client/mtlsauthentication) ([API docs](https://grokify.github.io/saviynt-api-go-client/api_mtlsauthentication.html))
1. [x] [SAV Roles](https://pkg.go.dev/github.com/grokify/saviynt-api-go-client/savroles) ([API docs](https://grokify.github.io/saviynt-api-go-client/api_savroles.html))
1. [x] [Transport](https://pkg.go.dev/github.com/grokify/saviynt-api-go-client/transport) ([API docs](https://grokify.github.io/saviynt-api-go-client/api_transport.html))
