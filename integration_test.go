package ligovalidator_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/linkeunid/ligo"
	"github.com/linkeunid/ligo/adapters/echo"
	ligovalidator "github.com/linkeunid/ligo-validator"
)

func startApp(t *testing.T, modules ...ligo.Module) *ligo.App {
	t.Helper()
	app := ligo.New(
		ligo.WithRouter(echo.NewAdapter()),
		ligo.WithAddr(":0"),
	)
	app.Register(modules...)
	go func() { _ = app.Run() }()
	time.Sleep(200 * time.Millisecond)
	return app
}

func containsType(types []reflect.Type, want reflect.Type) bool {
	for _, typ := range types {
		if typ == want {
			return true
		}
	}
	return false
}

// TestModuleIntegratesWithApp verifies that ligovalidator.Module() registers
// successfully in a Ligo app and exposes *validator.Validate in the DI container.
func TestModuleIntegratesWithApp(t *testing.T) {
	app := startApp(t, ligovalidator.Module())

	want := reflect.TypeFor[*validator.Validate]()
	if !containsType(app.Container().Types(), want) {
		t.Fatalf("*validator.Validate not found in DI container; registered types: %v", app.Container().Types())
	}
}

// TestProviderInjectsValidate verifies that Provider() registers a
// *validator.Validate that the DI container can inject into a dependent factory.
func TestProviderInjectsValidate(t *testing.T) {
	type payload struct {
		Name string `validate:"required"`
	}

	var injected *validator.Validate

	mod := ligo.NewModule("test",
		ligo.Providers(
			ligovalidator.Provider(),
			ligo.Factory[*payload](func(v *validator.Validate) *payload {
				injected = v
				return &payload{}
			}),
		),
	)
	app := startApp(t, mod)

	want := reflect.TypeFor[*validator.Validate]()
	if !containsType(app.Container().Types(), want) {
		t.Fatalf("*validator.Validate not found in DI container; registered types: %v", app.Container().Types())
	}

	if injected != nil {
		p := &payload{Name: "Alice"}
		if err := injected.Struct(p); err != nil {
			t.Fatalf("expected valid struct, got error: %v", err)
		}

		empty := &payload{}
		if err := injected.Struct(empty); err == nil {
			t.Fatal("expected validation error for empty Name, got nil")
		}
	}
}
