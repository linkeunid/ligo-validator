// Package ligovalidator provides a DI-ready wrapper for go-playground/validator,
// registering a singleton *validator.Validate in the Ligo DI container.
package ligovalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/linkeunid/ligo"
)

// Provider returns a [ligo.Provider] that registers a *[validator.Validate]
// as a singleton in the DI container.
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
		return validator.New()
	})
}

// Module returns a Ligo module that registers a *[validator.Validate] singleton
// via DI. It is a convenient zero-config drop-in for applications that want a
// single shared validator instance without declaring it inside their own modules.
//
// For custom validator configuration (registering tag names, custom rules, etc.),
// use [Provider] directly inside your own module and wrap the factory.
//
//	app.Register(ligovalidator.Module(), myModule())
func Module() ligo.Module {
	return ligo.NewModule("ligo-validator",
		ligo.Providers(
			Provider(),
		),
	)
}
