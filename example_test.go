package response_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	response "github.com/mawarpay/pkg-responses"
)

func ExampleBuildResponseCode() {
	code := response.BuildResponseCode(
		http.StatusCreated,
		response.ServiceCodeWithdrawal,
		response.CaseCodeSuccess,
	)
	fmt.Println(code)
	// Output: 2010301
}

func ExampleParseResponseCode() {
	status, service, caseCode := response.ParseResponseCode(2010301)
	fmt.Println(status, service, caseCode)
	// Output: 201 03 01
}

func ExampleOkWithDetailed() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.OkWithDetailed(
		c,
		http.StatusOK,
		response.ServiceCodeAuth,
		response.CaseCodeLoginSuccess,
		map[string]string{"token": "abc"},
		"login successful",
	)

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())
	// Output:
	// 200
	// {"code":2000107,"message":"login successful","data":{"token":"abc"}}
}

func ExampleValidationErrorSimple() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.ValidationErrorSimple(
		c,
		response.ServiceCodeAuth,
		"email",
		"The email field is required.",
	)

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())
	// Output:
	// 422
	// {"code":4220111,"message":"The given data was invalid.","errors":{"email":["The email field is required."]}}
}

func ExampleWriteCursorPaginated() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	next := "cursor-2"
	response.WriteCursorPaginated(c, response.CursorPaginatedParams{
		HTTPStatus:  http.StatusOK,
		ServiceCode: response.ServiceCodeUser,
		CaseCode:    response.CaseCodeListRetrieved,
		Message:     "ok",
		Pagination: response.CursorPaginationResponse{
			Data:       []string{"a", "b"},
			NextCursor: &next,
			HasNext:    true,
		},
	})

	fmt.Println(w.Body.String())
	// Output:
	// {"code":2000406,"message":"ok","data":["a","b"],"nextCursor":"cursor-2","hasNext":true}
}
