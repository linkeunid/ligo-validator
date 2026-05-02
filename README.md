# ligo-validator

A DI-ready wrapper for [go-playground/validator](https://github.com/go-playground/validator) for [Ligo](https://github.com/linkeunid/ligo), with class-validator style built-in helpers.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue)](https://go.dev/dl)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Tests](https://img.shields.io/badge/tests-38%20passing-brightgreen)](https://github.com/linkeunid/ligo-validator)

## Install

```bash
go get github.com/linkeunid/ligo-validator
```

## Quick start

Register the provider in your module and inject `*validator.Validate` into your constructors:

```go
import (
    "github.com/go-playground/validator/v10"
    "github.com/linkeunid/ligo"
    ligovalidator "github.com/linkeunid/ligo-validator"
)

func UserModule() ligo.Module {
    return ligo.NewModule("user",
        ligo.Providers(
            ligovalidator.Provider(),
            ligo.Factory[*usecase.UserUseCase](usecase.NewUserUseCase),
        ),
    )
}

func NewUserUseCase(repo repository.UserRepository, verify *validator.Validate) *UserUseCase {
    return &UserUseCase{repo: repo, verify: verify}
}
```

For a zero-config drop-in:

```go
app.Register(ligovalidator.Module(), myModule())
```

## See also

- [DI integration & custom tags](docs/features/module.md)
- [String helpers](docs/features/str.md)
- [Slice helpers](docs/features/slice.md)
- [Number helpers](docs/features/num.md)
- [Object helpers](docs/features/obj.md)
