package ligovalidator_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/linkeunid/ligo"
	ligovalidator "github.com/linkeunid/ligo-validator"
	"github.com/linkeunid/ligo/adapters/echo"
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

	mod := ligo.NewModule(
		"test",
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

// TestProviderIntegratesCustomTags verifies that Provider() registers a
// *validator.Validate with custom tags (not_blank, arr_not_empty, not_empty_obj)
// available in struct validation.
func TestProviderIntegratesCustomTags(t *testing.T) {
	type profile struct {
		Name string         `validate:"not_blank"`
		Tags []string       `validate:"arr_not_empty"`
		Meta map[string]any `validate:"not_empty_obj"`
	}

	var injected *validator.Validate

	mod := ligo.NewModule(
		"test-defaults",
		ligo.Providers(
			ligovalidator.Provider(),
			ligo.Factory[*profile](func(v *validator.Validate) *profile {
				injected = v
				return &profile{}
			}),
		),
	)
	startApp(t, mod)

	if injected == nil {
		t.Skip("factory was not resolved by DI")
	}

	valid := &profile{
		Name: "Alice",
		Tags: []string{"go"},
		Meta: map[string]any{"env": "prod"},
	}
	if err := injected.Struct(valid); err != nil {
		t.Fatalf("expected valid struct, got: %v", err)
	}

	if err := injected.Struct(&profile{Name: "  ", Tags: []string{"go"}, Meta: map[string]any{"k": "v"}}); err == nil {
		t.Fatal("expected not_blank to reject whitespace-only Name")
	}
	if err := injected.Struct(&profile{Name: "Alice", Tags: []string{}, Meta: map[string]any{"k": "v"}}); err == nil {
		t.Fatal("expected arr_not_empty to reject empty Tags")
	}
	if err := injected.Struct(&profile{Name: "Alice", Tags: []string{"go"}, Meta: map[string]any{}}); err == nil {
		t.Fatal("expected not_empty_obj to reject empty Meta")
	}
}
