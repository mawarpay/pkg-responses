/*
Package response provides standardized Gin JSON envelopes and a 7-digit composite
response code for IlonaPay / MawarPay HTTP APIs.

# Composite code

Every response carries an integer code built as:

	HTTP_STATUS (3 digits) + SERVICE_CODE (2 digits) + CASE_CODE (2 digits)

Example: 2010301 is HTTP 201, service 03 (withdrawal), case 01 (success).
Use [BuildResponseCode] and [ParseResponseCode] to compose or split codes.
Service and case values live as package constants (see ServiceCode* and CaseCode*).

# Envelope shapes

  - [CommonResponse] — success and most error paths (code, message, data)
  - [ValidationErrorResponse] — HTTP 422 with Laravel-style field errors
  - [CursorPaginatedResponse] / [SimplePaginatedResponse] — list endpoints

Prefer the Write* helpers ([Write], [WriteCursorPaginated], [WriteSimplePaginated])
for new call sites; the older multi-arg Result / CursorPaginated / SimplePaginated
wrappers remain for compatibility.

# Constraints

This package only shapes JSON and builds codes. It must not contain business
logic, and it must not depend on domain or use-case packages. The only I/O is
writing to *gin.Context.
*/
package response
