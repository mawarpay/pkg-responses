package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommonResponse is the standard JSON body for success and most error responses.
// Code is the 7-digit composite (HTTP status + service + case).
type CommonResponse struct {
	Code    int         `json:"code"` // Custom response code: HTTP_STATUS + SERVICE_CODE + CASE_CODE (e.g., 2000401)
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// WriteParams groups fields for [Write]. Prefer this over long Result argument lists.
type WriteParams struct {
	HTTPStatus  int
	ServiceCode string
	CaseCode    string
	Data        interface{}
	Message     string
}

// Write sends a [CommonResponse] using [BuildResponseCode] from WriteParams.
// Prefer Write for new call sites; [Result] is a thin wrapper kept for compatibility.
func Write(ctx *gin.Context, p WriteParams) {
	responseCode := BuildResponseCode(p.HTTPStatus, p.ServiceCode, p.CaseCode)
	ctx.JSON(p.HTTPStatus, CommonResponse{
		Code:    responseCode,
		Message: p.Message,
		Data:    p.Data,
	})
}

// Result writes a [CommonResponse] with the given HTTP status and composite code parts.
// Prefer [Write] with [WriteParams] for new call sites.
func Result(
	ctx *gin.Context,
	httpStatus int,
	serviceCode, caseCode string,
	data interface{},
	message string,
) {
	Write(ctx, WriteParams{
		HTTPStatus:  httpStatus,
		ServiceCode: serviceCode,
		CaseCode:    caseCode,
		Data:        data,
		Message:     message,
	})
}

// ResultWithCode writes a [CommonResponse] with an already-composed responseCode.
// Use when the caller already has a 7-digit code and must not rebuild it.
func ResultWithCode(
	ctx *gin.Context,
	httpStatus int,
	responseCode int,
	data interface{},
	message string,
) {
	ctx.JSON(httpStatus, CommonResponse{
		Code:    responseCode,
		Message: message,
		Data:    data,
	})
}

// Ok writes HTTP 200 with ServiceCodeCommon and CaseCodeSuccess, no data payload.
func Ok(ctx *gin.Context) {
	Result(
		ctx,
		http.StatusOK,
		ServiceCodeCommon,
		CaseCodeSuccess,
		nil,
		"success",
	)
}

// OkWithMessage writes HTTP 200 success with a custom message and no data payload.
func OkWithMessage(ctx *gin.Context, message string) {
	Result(
		ctx,
		http.StatusOK,
		ServiceCodeCommon,
		CaseCodeSuccess,
		nil,
		message,
	)
}

// OkWithData writes HTTP 200 with CaseCodeRetrieved and the given data payload.
func OkWithData(ctx *gin.Context, data interface{}) {
	Result(
		ctx,
		http.StatusOK,
		ServiceCodeCommon,
		CaseCodeRetrieved,
		data,
		"success",
	)
}

// CursorPaginatedResponse is the top-level JSON shape for cursor-based pagination.
type CursorPaginatedResponse struct {
	Code       int         `json:"code"`       // Custom response code
	Message    string      `json:"message"`    // Response message
	Data       interface{} `json:"data"`       // The actual data array
	NextCursor *string     `json:"nextCursor"` // Cursor for the next page (null if no more pages)
	HasNext    bool        `json:"hasNext"`    // Whether there are more items available
}

// CursorPaginatedParams groups fields for [WriteCursorPaginated].
type CursorPaginatedParams struct {
	HTTPStatus  int
	ServiceCode string
	CaseCode    string
	Pagination  CursorPaginationResponse
	Message     string
}

// WriteCursorPaginated writes a [CursorPaginatedResponse] from CursorPaginatedParams.
func WriteCursorPaginated(ctx *gin.Context, p CursorPaginatedParams) {
	responseCode := BuildResponseCode(p.HTTPStatus, p.ServiceCode, p.CaseCode)
	ctx.JSON(p.HTTPStatus, CursorPaginatedResponse{
		Code:       responseCode,
		Message:    p.Message,
		Data:       p.Pagination.Data,
		NextCursor: p.Pagination.NextCursor,
		HasNext:    p.Pagination.HasNext,
	})
}

// CursorPaginated writes a cursor-based paginated JSON response.
// Prefer [WriteCursorPaginated] for new call sites.
func CursorPaginated(
	ctx *gin.Context,
	httpStatus int,
	serviceCode, caseCode string,
	pagination CursorPaginationResponse,
	message string,
) {
	WriteCursorPaginated(ctx, CursorPaginatedParams{
		HTTPStatus:  httpStatus,
		ServiceCode: serviceCode,
		CaseCode:    caseCode,
		Pagination:  pagination,
		Message:     message,
	})
}

// SimplePaginatedResponse is the top-level JSON shape for offset pagination.
type SimplePaginatedResponse struct {
	Code       int         `json:"code"`       // Custom response code
	Message    string      `json:"message"`    // Response message
	Data       interface{} `json:"data"`       // The actual data array
	PageNumber int         `json:"pageNumber"` // Current page number
	PageSize   int         `json:"pageSize"`   // Number of items per page
	HasNext    bool        `json:"hasNext"`    // Whether there is a next page
	HasPrev    bool        `json:"hasPrev"`    // Whether there is a previous page
}

// SimplePaginatedParams groups fields for [WriteSimplePaginated].
type SimplePaginatedParams struct {
	HTTPStatus  int
	ServiceCode string
	CaseCode    string
	Pagination  SimplePaginationResponse
	Message     string
}

// WriteSimplePaginated writes a [SimplePaginatedResponse] from SimplePaginatedParams.
func WriteSimplePaginated(ctx *gin.Context, p SimplePaginatedParams) {
	responseCode := BuildResponseCode(p.HTTPStatus, p.ServiceCode, p.CaseCode)
	ctx.JSON(p.HTTPStatus, SimplePaginatedResponse{
		Code:       responseCode,
		Message:    p.Message,
		Data:       p.Pagination.Data,
		PageNumber: p.Pagination.PageNumber,
		PageSize:   p.Pagination.PageSize,
		HasNext:    p.Pagination.HasNext,
		HasPrev:    p.Pagination.HasPrev,
	})
}

// SimplePaginated writes an offset-paginated JSON response.
// Prefer [WriteSimplePaginated] for new call sites.
func SimplePaginated(
	ctx *gin.Context,
	httpStatus int,
	serviceCode, caseCode string,
	pagination SimplePaginationResponse,
	message string,
) {
	WriteSimplePaginated(ctx, SimplePaginatedParams{
		HTTPStatus:  httpStatus,
		ServiceCode: serviceCode,
		CaseCode:    caseCode,
		Pagination:  pagination,
		Message:     message,
	})
}

// OkWithDetailed writes a success-shaped [CommonResponse] with full control over
// HTTP status, service code, case code, data, and message.
func OkWithDetailed(
	ctx *gin.Context,
	httpStatus int,
	serviceCode, caseCode string,
	data interface{},
	message string,
) {
	Result(
		ctx,
		httpStatus,
		serviceCode,
		caseCode,
		data,
		message,
	)
}

// Created writes HTTP 201 with CaseCodeCreated.
// An empty message defaults to "Resource created successfully".
func Created(ctx *gin.Context, serviceCode string, data interface{}, message string) {
	if message == "" {
		message = "Resource created successfully"
	}
	Result(
		ctx,
		http.StatusCreated,
		serviceCode,
		CaseCodeCreated,
		data,
		message,
	)
}

// Updated writes HTTP 200 with CaseCodeUpdated.
// An empty message defaults to "Resource updated successfully".
func Updated(ctx *gin.Context, serviceCode string, data interface{}, message string) {
	if message == "" {
		message = "Resource updated successfully"
	}
	Result(
		ctx,
		http.StatusOK,
		serviceCode,
		CaseCodeUpdated,
		data,
		message,
	)
}

// Deleted writes HTTP 200 with CaseCodeDeleted and a nil data payload.
// An empty message defaults to "Resource deleted successfully".
func Deleted(ctx *gin.Context, serviceCode string, message string) {
	if message == "" {
		message = "Resource deleted successfully"
	}
	Result(
		ctx,
		http.StatusOK,
		serviceCode,
		CaseCodeDeleted,
		nil,
		message,
	)
}

// Fail writes HTTP 500 with ServiceCodeCommon and CaseCodeInternalError.
func Fail(ctx *gin.Context) {
	Result(
		ctx,
		http.StatusInternalServerError,
		ServiceCodeCommon,
		CaseCodeInternalError,
		nil,
		"failure",
	)
}

// FailWithMessage writes HTTP 500 with a custom message and CaseCodeInternalError.
func FailWithMessage(ctx *gin.Context, message string) {
	Result(
		ctx,
		http.StatusInternalServerError,
		ServiceCodeCommon,
		CaseCodeInternalError,
		nil,
		message,
	)
}

// FailWithDetailed writes an error-shaped [CommonResponse] with full control over
// HTTP status, service code, case code, data, and message.
func FailWithDetailed(
	ctx *gin.Context,
	httpStatus int,
	serviceCode, caseCode string,
	data interface{},
	message string,
) {
	Result(
		ctx,
		httpStatus,
		serviceCode,
		caseCode,
		data,
		message,
	)
}

// ValidationError writes HTTP 422 with CaseCodeValidationError and a Laravel-style
// errors map derived from err via [FormatValidationError].
func ValidationError(ctx *gin.Context, serviceCode string, err error) {
	errors := FormatValidationError(err)
	message := "The given data was invalid."

	responseCode := BuildResponseCode(
		http.StatusUnprocessableEntity,
		serviceCode,
		CaseCodeValidationError,
	)

	ctx.JSON(http.StatusUnprocessableEntity, ValidationErrorResponse{
		Code:    responseCode,
		Message: message,
		Errors:  errors,
	})
}

// ValidationErrorWithMessage writes HTTP 422 with a custom message and errors map.
// An empty message defaults to "The given data was invalid.".
// A nil errors map is replaced with an empty map.
func ValidationErrorWithMessage(
	ctx *gin.Context,
	serviceCode string,
	message string,
	errors map[string][]string,
) {
	if message == "" {
		message = "The given data was invalid."
	}
	if errors == nil {
		errors = make(map[string][]string)
	}

	responseCode := BuildResponseCode(
		http.StatusUnprocessableEntity,
		serviceCode,
		CaseCodeValidationError,
	)

	ctx.JSON(http.StatusUnprocessableEntity, ValidationErrorResponse{
		Code:    responseCode,
		Message: message,
		Errors:  errors,
	})
}

// ValidationErrorSimple writes HTTP 422 for a single field error.
func ValidationErrorSimple(
	ctx *gin.Context,
	serviceCode string,
	fieldName string,
	errorMessage string,
) {
	errors := map[string][]string{
		fieldName: {errorMessage},
	}
	ValidationErrorWithMessage(
		ctx,
		serviceCode,
		"The given data was invalid.",
		errors,
	)
}

// UnauthorizedError writes HTTP 401 with ServiceCodeAuth and CaseCodeUnauthorized.
// An empty message defaults to "Unauthorized".
func UnauthorizedError(ctx *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	Result(
		ctx,
		http.StatusUnauthorized,
		ServiceCodeAuth,
		CaseCodeUnauthorized,
		nil,
		message,
	)
}

// NotFoundError writes HTTP 404 for a missing resource.
// An empty message defaults to "Resource not found".
// An empty caseCode defaults to CaseCodeNotFound.
func NotFoundError(ctx *gin.Context, serviceCode, caseCode string, message string) {
	if message == "" {
		message = "Resource not found"
	}
	if caseCode == "" {
		caseCode = CaseCodeNotFound
	}
	Result(
		ctx,
		http.StatusNotFound,
		serviceCode,
		caseCode,
		nil,
		message,
	)
}

// ConflictError writes HTTP 409 with CaseCodeConflict.
// An empty message defaults to "Resource conflict".
func ConflictError(ctx *gin.Context, serviceCode string, message string) {
	if message == "" {
		message = "Resource conflict"
	}
	Result(
		ctx,
		http.StatusConflict,
		serviceCode,
		CaseCodeConflict,
		nil,
		message,
	)
}

// ForbiddenError writes HTTP 403 with ServiceCodeAuth and CaseCodePermissionDenied.
// An empty message defaults to "Forbidden".
func ForbiddenError(ctx *gin.Context, message string) {
	if message == "" {
		message = "Forbidden"
	}
	Result(
		ctx,
		http.StatusForbidden,
		ServiceCodeAuth,
		CaseCodePermissionDenied,
		nil,
		message,
	)
}
