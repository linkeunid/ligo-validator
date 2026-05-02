// Package ligovalidator provides a DI-ready wrapper for go-playground/validator,
// registering a singleton *validator.Validate in the Ligo DI container.
// The singleton is pre-configured with built-in custom tags via [RegisterAll].
package ligovalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/linkeunid/ligo"
)

// RegisterAll registers all built-in ligo-validator custom struct tags on v.
//
// Custom tags:
//   - "not_blank"      — string must contain at least one non-whitespace character
//   - "arr_not_empty"  — slice or array must have at least one element
//   - "not_empty_obj"  — struct or map must have at least one non-zero field/entry
func RegisterAll(v *validator.Validate) {
	if err := v.RegisterValidation("not_blank", func(fl validator.FieldLevel) bool {
		return IsNotBlank(fl.Field().String())
	}); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("arr_not_empty", func(fl validator.FieldLevel) bool {
		return fl.Field().Len() > 0
	}); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("not_empty_obj", func(fl validator.FieldLevel) bool {
		return IsNotEmptyObject(fl.Field().Interface())
	}); err != nil {
		panic(err)
	}
}

// Provider returns a [ligo.Provider] that registers a *[validator.Validate]
// singleton in the DI container with all built-in custom tags pre-registered.
//
// Use inside your module's [ligo.Providers] list:
//
//	func User() ligo.Module {
//	    return ligo.NewModule("user",
//	        ligo.Providers(
//	            ligovalidator.Provider(),
//	            ligo.Factory[repository.UserRepository](memory.NewUserRepository),
//	            ligo.Factory[*usecase.UserUseCase](usecase.NewUserUseCase),
//	        ),
//	        ligo.Controllers(controller.NewUserController),
//	    )
//	}
//
// Accept the validator in your constructor:
//
//	func NewUserUseCase(repo repository.UserRepository, log ligo.Logger, verify *validator.Validate) *UserUseCase {
//	    return &UserUseCase{repo: repo, log: log, verify: verify}
//	}
func Provider() ligo.Provider {
	return ligo.Factory[*validator.Validate](func() *validator.Validate {
		v := validator.New()
		RegisterAll(v)
		return v
	})
}

// Module returns a Ligo module that registers a *[validator.Validate] singleton
// via DI with all built-in custom tags pre-registered. It is a convenient
// zero-config drop-in.
//
//	app.Register(ligovalidator.Module(), myModule())
func Module() ligo.Module {
	return ligo.NewModule("ligo-validator",
		ligo.Providers(
			Provider(),
		),
	)
}
