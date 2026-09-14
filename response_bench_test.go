package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BenchmarkBuildResponseCode(b *testing.B) {
	b.ReportAllocs()
	var sink int
	for b.Loop() {
		sink = BuildResponseCode(http.StatusOK, ServiceCodeCommon, CaseCodeSuccess)
	}
	_ = sink
}

func BenchmarkOkWithData(b *testing.B) {
	gin.SetMode(gin.TestMode)
	payload := map[string]any{"id": "550e8400-e29b-41d4-a716-446655440000", "name": "bench"}

	b.ReportAllocs()
	for b.Loop() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
		OkWithData(c, payload)
	}
}

func BenchmarkResult(b *testing.B) {
	gin.SetMode(gin.TestMode)
	payload := map[string]any{"ok": true}

	b.ReportAllocs()
	for b.Loop() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
		Result(
			c,
			http.StatusOK,
			ServiceCodeCommon,
			CaseCodeSuccess,
			payload,
			"success",
		)
	}
}

func BenchmarkParseResponseCode(b *testing.B) {
	b.ReportAllocs()
	var status int
	var svc, cs string
	for b.Loop() {
		status, svc, cs = ParseResponseCode(2010301)
	}
	_, _, _ = status, svc, cs
}

func BenchmarkFormatValidationError(b *testing.B) {
	type req struct {
		Email string `json:"email" validate:"required,email"`
		Name  string `json:"name" validate:"min=2"`
		Age   int    `json:"age" validate:"gte=18"`
	}
	validate := validator.New()
	err := validate.Struct(req{Email: "bad", Name: "x", Age: 1})
	if err == nil {
		b.Fatal("expected validation errors")
	}

	b.ReportAllocs()
	var sink map[string][]string
	for b.Loop() {
		sink = FormatValidationError(err)
	}
	_ = sink
}

func BenchmarkFormatValidationError_Generic(b *testing.B) {
	err := errors.New("something failed")
	b.ReportAllocs()
	var sink map[string][]string
	for b.Loop() {
		sink = FormatValidationError(err)
	}
	_ = sink
}
