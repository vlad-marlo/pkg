# pkg [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

`fe` allows combining error message, logger fields and response code in Go `error` together.

## Features

- **Interoperable**:
  fe interoperates with the Go standard library's error APIs seamlessly:
  - The `errors.Is` and `errors.As` functions *just work*.
- **Lightweight**:
  fe comes with virtually no dependencies.

## Installation

```bash
go get -u github.com/vlad-marlo/pkg@latest
```

## Status

Stable: No breaking changes will be made before 2.0.

-------------------------------------------------------------------------------

Released under the [MIT License].

[MIT License]: LICENSE

[doc-img]: https://pkg.go.dev/badge/github.com/vlad-marlo/pkg

[doc]: https://pkg.go.dev/github.com/vlad-marlo/pkg

[ci-img]: https://github.com/vlad-marlo/pkg/actions/workflows/go.yml/badge.svg

[cov-img]: https://codecov.io/gh/vlad-marlo/pkg/branch/main/graph/badge.svg?token=9ECKEYBGHR

[ci]: https://github.com/vlad-marlo/pkg/actions/workflows/go.yml

[cov]: https://codecov.io/gh/vlad-marlo/pkg