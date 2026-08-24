package directives

import (
	"context"
	"testing"
)

const iso2 = `^[A-Z]{2}$`

func TestValidateRegex_TypedNilPointerOrNil(t *testing.T) {
	next := func(ctx context.Context) (any, error) {
		var s *string
		return s, nil
	}
	res, err := ValidateRegex(context.Background(), nil, next, iso2, true, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatalf("expected nil, got %#v", res)
	}
}

func TestValidateRegex_UntypedNilOrNil(t *testing.T) {
	next := func(ctx context.Context) (any, error) {
		return nil, nil
	}
	res, err := ValidateRegex(context.Background(), nil, next, iso2, true, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatalf("expected nil, got %#v", res)
	}
}

func TestValidateRegex_TypedNilPointerRequired(t *testing.T) {
	next := func(ctx context.Context) (any, error) {
		var s *string
		return s, nil
	}
	_, err := ValidateRegex(context.Background(), nil, next, iso2, true, false)
	if err != ErrValueIsNil {
		t.Fatalf("expected ErrValueIsNil, got %v", err)
	}
}

func TestValidateRegex_ValidValue(t *testing.T) {
	code := "US"
	next := func(ctx context.Context) (any, error) {
		return &code, nil
	}
	res, err := ValidateRegex(context.Background(), nil, next, iso2, true, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, ok := res.(*string)
	if !ok || got == nil || *got != "US" {
		t.Fatalf("expected *string US, got %#v", res)
	}
}
