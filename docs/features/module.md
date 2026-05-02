# DI integration & custom tags

`module.go` provides the Ligo DI wiring and registers all built-in custom struct tags on the shared `*validator.Validate` instance.

---

## Provider()

Returns a `ligo.Provider` that registers a `*validator.Validate` singleton in the DI container with all built-in custom tags pre-registered via `RegisterAll`.

```go
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

## Module()

Zero-config drop-in that registers the singleton without declaring a provider inside your own module.

```go
app.Register(ligovalidator.Module(), myModule())
```

---

## RegisterAll(v)

Registers all built-in custom struct tags on an existing `*validator.Validate`. Panics if any tag registration fails (e.g. duplicate tag name). Useful when you need to layer additional custom rules on top:

```go
v := validator.New()
ligovalidator.RegisterAll(v)
v.RegisterValidation("my_rule", myFunc)
```

### Built-in tags

| Tag | Description |
|-----|-------------|
| `not_blank` | String must contain at least one non-whitespace character |
| `arr_not_empty` | Slice or array must have at least one element |
| `not_empty_obj` | Struct or map must have at least one non-zero field/entry |

```go
type CreateUserDTO struct {
    Name  string         `validate:"not_blank"`
    Roles []string       `validate:"arr_not_empty"`
    Meta  map[string]any `validate:"not_empty_obj"`
}

if err := verify.Struct(dto); err != nil {
    // handle validation errors
}
```
