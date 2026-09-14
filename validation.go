package response

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationErrorResponse is the JSON shape for HTTP 422 validation failures.
// Errors maps field names to one or more human-readable messages (Laravel style).
type ValidationErrorResponse struct {
	Code    int                 `json:"code"`    // Custom response code
	Message string              `json:"message"` // General error message
	Errors  map[string][]string `json:"errors"`  // Field-specific errors
}

// FormatValidationError converts err into a map of field names to message slices.
// If err is validator.ValidationErrors, each field is keyed separately with
// Laravel-style messages. Any other error type is stored under "general".
// A nil err returns an empty (non-nil) map.
func FormatValidationError(err error) map[string][]string {
	if err == nil {
		return make(map[string][]string)
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return map[string][]string{
			"general": {err.Error()},
		}
	}

	// Pre-size to avoid map growth when several fields fail.
	out := make(map[string][]string, len(validationErrors))
	for _, fieldError := range validationErrors {
		fieldName := getFieldName(fieldError)
		errorMessage := getErrorMessage(fieldError, fieldName)
		out[fieldName] = append(out[fieldName], errorMessage)
	}

	return out
}

// getFieldName extracts the field name from a validation error.
// It prefers JSON-oriented names when struct context is available.
func getFieldName(fieldError validator.FieldError) string {
	structField := fieldError.StructField()
	namespace := fieldError.Namespace()
	fieldName := fieldError.Field()

	hasStructContext := structField != "" && namespace != "" && strings.Contains(namespace, ".")
	if hasStructContext {
		if fieldName != "" && fieldName != structField {
			return fieldName
		}

		return toCamelCase(structField)
	}

	if fieldName != "" {
		return fieldName
	}

	if structField != "" {
		return toCamelCase(structField)
	}

	return strings.ToLower(fieldName)
}

// toCamelCase converts "FirstName" to "firstName".
func toCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	// Check if already in camelCase (starts with lowercase)
	if s[0] >= 'a' && s[0] <= 'z' {
		return s
	}

	// Convert first letter to lowercase
	return strings.ToLower(s[:1]) + s[1:]
}

// validationTagMessages maps validator tags to Laravel-style sprintf templates.
// Templates use one %s (field) or two %s (field, param).
var validationTagMessages = map[string]string{
	"required":         "The %s field is required.",
	"email":            "The %s must be a valid email address.",
	"min":              "The %s must be at least %s characters.",
	"max":              "The %s may not be greater than %s characters.",
	"len":              "The %s must be exactly %s characters.",
	"numeric":          "The %s must be a number.",
	"alpha":            "The %s may only contain letters.",
	"alphanum":         "The %s may only contain letters and numbers.",
	"url":              "The %s must be a valid URL.",
	"uuid":             "The %s must be a valid UUID.",
	"oneof":            "The %s must be one of: %s.",
	"gte":              "The %s must be greater than or equal to %s.",
	"lte":              "The %s must be less than or equal to %s.",
	"gt":               "The %s must be greater than %s.",
	"lt":               "The %s must be less than %s.",
	"eq":               "The %s must be equal to %s.",
	"ne":               "The %s must not be equal to %s.",
	"unique":           "The %s has already been taken.",
	"exists":           "The selected %s is invalid.",
	"date":             "The %s must be a valid date.",
	"datetime":         "The %s must be a valid date and time.",
	"timezone":         "The %s must be a valid timezone.",
	"json":             "The %s must be a valid JSON string.",
	"ip":               "The %s must be a valid IP address.",
	"ipv4":             "The %s must be a valid IPv4 address.",
	"ipv6":             "The %s must be a valid IPv6 address.",
	"base64":           "The %s must be a valid base64 string.",
	"required_if":      "The %s field is required when %s is present.",
	"required_unless":  "The %s field is required unless %s is present.",
	"required_with":    "The %s field is required when %s is present.",
	"required_without": "The %s field is required when %s is not present.",
}

// getErrorMessage generates a human-readable error message from a validation error.
func getErrorMessage(fieldError validator.FieldError, fieldName string) string {
	tag := fieldError.Tag()
	param := fieldError.Param()

	if tmpl, ok := validationTagMessages[tag]; ok {
		if strings.Count(tmpl, "%s") == 2 {
			return fmt.Sprintf(tmpl, fieldName, param)
		}

		return fmt.Sprintf(tmpl, fieldName)
	}

	if param != "" {
		return fmt.Sprintf("The %s field is invalid. (%s: %s)", fieldName, tag, param)
	}

	return fmt.Sprintf("The %s field is invalid. (%s)", fieldName, tag)
}
