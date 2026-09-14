package response

import (
	"errors"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"FirstName", "firstName"},
		{"id", "id"},
		{"", ""},
		{"U", "u"},
		{"HTTP", "hTTP"},
	}
	for _, tt := range tests {
		got := toCamelCase(tt.in)
		if got != tt.want {
			t.Errorf("toCamelCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestFormatValidationError_Nil(t *testing.T) {
	got := FormatValidationError(nil)
	if got == nil {
		t.Fatal("FormatValidationError(nil) returned nil map")
	}
	if len(got) != 0 {
		t.Errorf("FormatValidationError(nil) = %v, want empty map", got)
	}
}

func TestFormatValidationError_GenericError(t *testing.T) {
	err := errors.New("something failed")
	got := FormatValidationError(err)
	if len(got) != 1 || len(got["general"]) != 1 || got["general"][0] != "something failed" {
		t.Errorf("FormatValidationError(plain error) = %v, want {\"general\":[\"something failed\"]}", got)
	}
}

func TestFormatValidationError_ValidatorErrors(t *testing.T) {
	type req struct {
		Email string `json:"email" validate:"required,email"`
		Name  string `json:"name" validate:"min=2"`
	}
	validate := validator.New()
	v := req{Email: "", Name: "x"}
	err := validate.Struct(v)
	if err == nil {
		t.Fatal("expected validation to fail")
	}
	got := FormatValidationError(err)
	if len(got) == 0 {
		t.Error("FormatValidationError(ValidationErrors) expected non-empty map")
	}
	if _, has := got["email"]; !has {
		if _, hasG := got["general"]; !hasG {
			t.Errorf("expected 'email' or 'general' in %v", got)
		}
	}
}

func TestGetErrorMessage_RequiredTag(t *testing.T) {
	type r struct {
		F string `validate:"required"`
	}
	validate := validator.New()
	err := validate.Struct(r{})
	if err == nil {
		t.Fatal("expected validation to fail")
	}

	m := FormatValidationError(err)
	for _, msgs := range m {
		for _, msg := range msgs {
			if strings.HasPrefix(msg, "The ") && strings.Contains(msg, "required") {
				return
			}
		}
	}
	if len(m["general"]) > 0 {
		return
	}
	t.Errorf("expected a required-field message, got %v", m)
}
