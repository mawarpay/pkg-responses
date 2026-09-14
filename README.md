# pkg-responses

[![Go Version](https://img.shields.io/github/go-mod/go-version/mawarpay/pkg-responses)](https://go.dev/)
[![License](https://img.shields.io/github/license/mawarpay/pkg-responses)](./LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/mawarpay/pkg-responses/test.yml?branch=main)](https://github.com/mawarpay/pkg-responses/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/mawarpay/pkg-responses.svg)](https://pkg.go.dev/github.com/mawarpay/pkg-responses)
[![Go Report Card](https://goreportcard.com/badge/github.com/mawarpay/pkg-responses)](https://goreportcard.com/report/github.com/mawarpay/pkg-responses)

Standardized Gin JSON responses for IlonaPay / MawarPay APIs using a 7-digit composite code:

`HTTP_STATUS` (3) + `SERVICE_CODE` (2) + `CASE_CODE` (2)

```go
code := response.BuildResponseCode(http.StatusCreated, response.ServiceCodeWithdrawal, response.CaseCodeSuccess)
// code == 2010301
```

## Getting Started

```bash
go get github.com/mawarpay/pkg-responses@latest
```

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mawarpay/pkg-responses"
)

func login(c *gin.Context) {
	response.OkWithDetailed(
		c,
		http.StatusOK,
		response.ServiceCodeAuth,
		response.CaseCodeLoginSuccess,
		gin.H{"token": "..."},
		"login successful",
	)
}

func createUser(c *gin.Context) {
	response.Created(c, response.ServiceCodeUser, gin.H{"id": "1"}, "")
}

func badRequest(c *gin.Context, err error) {
	response.ValidationError(c, response.ServiceCodeAuth, err)
}
```

Envelope:

```json
{
  "code": 2000107,
  "message": "login successful",
  "data": { "token": "..." }
}
```

## Features

### Composite response codes

[BuildResponseCode](https://pkg.go.dev/github.com/mawarpay/pkg-responses#BuildResponseCode) and [ParseResponseCode](https://pkg.go.dev/github.com/mawarpay/pkg-responses#ParseResponseCode) compose and split the 7-digit code. Service and case values are package constants (`ServiceCode*`, `CaseCode*`).

| Segment | Width | Example |
| --- | --- | --- |
| HTTP status | 3 | `201` |
| Service | 2 | `03` (withdrawal) |
| Case | 2 | `01` (success) |

### Gin writers

Prefer the `Write*` helpers for new call sites:

| Helper | When to use |
| --- | --- |
| `Write` / `Ok*` / `Created` / `Updated` / `Deleted` | Success paths |
| `Fail*` / `UnauthorizedError` / `ForbiddenError` / `NotFoundError` / `ConflictError` | Error paths |
| `ValidationError*` | HTTP 422 with Laravel-style field maps |
| `WriteCursorPaginated` / `WriteSimplePaginated` | List endpoints |

### Validation formatting

[FormatValidationError](https://pkg.go.dev/github.com/mawarpay/pkg-responses#FormatValidationError) turns `validator.ValidationErrors` into `map[string][]string`. Non-validator errors land under `"general"`.

### Pagination defaults

`DefaultPageNumber` (1), `DefaultPageSize` (10), `MaxPageSize` (100) are shared clamps for handlers and repositories.

### Package layout

| File | Responsibility |
| --- | --- |
| `codes.go` | Service/case constants, `BuildResponseCode`, `ParseResponseCode` |
| `response.go` | Gin writers and envelope types |
| `validation.go` | Laravel-style field maps from `validator` errors |
| `pagination.go` | Cursor and offset pagination payloads |

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a PR.

## License

This project is licensed under the MIT License — see [LICENSE](LICENSE).
