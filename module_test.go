package ligovalidator

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/linkeunid/ligo"
)

func TestProviderType(t *testing.T) {
	p := Provider()
	want := reflect.TypeFor[*validator.Validate]()
	if p.Type() != want {
		t.Fatalf("expected type %v, got %v", want, p.Type())
	}
}

func TestProviderReturnsNonNil(t *testing.T) {
	p := Provider()
	if p.Fn() == nil {
		t.Fatal("expected non-nil factory function")
	}
}

func TestModuleName(t *testing.T) {
	m := Module()
	if m.Name != "ligo-validator" {
		t.Fatalf("expected module name %q, got %q", "ligo-validator", m.Name)
	}
}

func TestModuleRegistersValidateProvider(t *testing.T) {
	m := Module()
	want := reflect.TypeFor[*validator.Validate]()
	for _, raw := range m.Providers {
		if p, ok := raw.(ligo.Provider); ok && p.Type() == want {
			return
		}
	}
	t.Fatalf("Module must register *validator.Validate; providers: %v", m.Providers)
}

func TestModuleProviderType(t *testing.T) {
	m := Module()
	p := m.Providers[0].(ligo.Provider)
	want := reflect.TypeFor[*validator.Validate]()
	if p.Type() != want {
		t.Fatalf("Module provider type: expected %v, got %v", want, p.Type())
	}
}

type notBlankFixture struct {
	Name string `validate:"not_blank"`
}

type arrNotEmptyFixture struct {
	Tags []string `validate:"arr_not_empty"`
}

type notEmptyObjFixture struct {
	Meta map[string]any `validate:"not_empty_obj"`
}

func setupValidator(t *testing.T) *validator.Validate {
	t.Helper()
	v := validator.New()
	RegisterAll(v)
	return v
}

func TestRegisterAllNotBlank(t *testing.T) {
	v := setupValidator(t)

	if err := v.Struct(notBlankFixture{Name: "Alice"}); err != nil {
		t.Fatalf("expected valid, got: %v", err)
	}
	if err := v.Struct(notBlankFixture{Name: "   "}); err == nil {
		t.Fatal("expected invalid for whitespace-only name")
	}
	if err := v.Struct(notBlankFixture{Name: ""}); err == nil {
		t.Fatal("expected invalid for empty name")
	}
}

func TestRegisterAllArrNotEmpty(t *testing.T) {
	v := setupValidator(t)

	if err := v.Struct(arrNotEmptyFixture{Tags: []string{"go"}}); err != nil {
		t.Fatalf("expected valid, got: %v", err)
	}
	if err := v.Struct(arrNotEmptyFixture{Tags: []string{}}); err == nil {
		t.Fatal("expected invalid for empty slice")
	}
}

func TestRegisterAllNotEmptyObj(t *testing.T) {
	v := setupValidator(t)

	if err := v.Struct(notEmptyObjFixture{Meta: map[string]any{"k": "v"}}); err != nil {
		t.Fatalf("expected valid, got: %v", err)
	}
	if err := v.Struct(notEmptyObjFixture{Meta: map[string]any{}}); err == nil {
		t.Fatal("expected invalid for empty map")
	}
}
