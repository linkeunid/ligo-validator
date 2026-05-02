# ligo-validator

A DI-ready wrapper for [go-playground/validator](https://github.com/go-playground/validator) for [Ligo](https://github.com/linkeunid/ligo) — registers a singleton `*validator.Validate` in the DI container.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue)](https://go.dev/dl)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Tests](https://img.shields.io/badge/tests-7%20passing-brightgreen)](https://github.com/linkeunid/ligo-validator)

## Install

```bash
go get github.com/linkeunid/ligo-validator
```

## Quick start

Register the validator provider in your module and inject it into your use cases:

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
            ligo.Factory[repository.UserRepository](memory.NewUserRepository),
            ligo.Factory[*usecase.UserUseCase](usecase.NewUserUseCase),
        ),
        ligo.Controllers(controller.NewUserController),
    )
}

func NewUserUseCase(repo repository.UserRepository, log ligo.Logger, verify *validator.Validate) *UserUseCase {
    return &UserUseCase{repo: repo, log: log, verify: verify}
}
```

For a zero-config drop-in, use `ligovalidator.Module()` which registers a shared `*validator.Validate`:

```go
app.Register(ligovalidator.Module(), myModule())
```

## Why

Without this package, every use case calls `validator.New()` itself — creating multiple instances with no shared configuration. `ligo-validator` gives you a single singleton, registered once, injectable anywhere.

| Before | After |
|--------|-------|
| `validator.New()` in every use case | Single singleton, registered once |
| Custom validators scattered | Registered in one place in `Provider()` |
| Hard to replace in tests | Inject a pre-configured instance |
