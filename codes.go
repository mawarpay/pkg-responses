package response

import "strconv"

// Service codes identify which service owns the response (2 digits, "00"–"21").
// Pass these as the serviceCode argument to BuildResponseCode and the Gin writers.
const (
	ServiceCodeCommon                      = "00" // Common/General services
	ServiceCodeAuth                        = "01" // Authentication service
	ServiceCodeTransaction                 = "02" // Transaction service
	ServiceCodeWithdrawal                  = "03" // Withdrawal service
	ServiceCodeUser                        = "04" // User service
	ServiceCodeAdmin                       = "05" // Admin service
	ServiceCodeMerchant                    = "06" // Merchant service
	ServiceCodeSetting                     = "07" // Setting service
	ServiceCodeRole                        = "08" // Role service
	ServiceCodePermission                  = "09" // Permission service
	ServiceCodeNotificationTemplate        = "10" // Notification template service
	ServiceCodeNotificationTemplateChannel = "11" // Notification template channel service
	ServiceCodeNotification                = "12" // Notification service
	ServiceCodeIPWhitelist                 = "13" // IP Whitelist service
	ServiceCodeApiKey                      = "14" // API Key service
	ServiceCodeDeposit                     = "15" // Deposit service
	ServiceCodeWallet                      = "16" // Wallet service
	ServiceCodePhone                       = "17" // Phone / phone verification service
	ServiceCodeEmail                       = "18" // Email / email verification service
	ServiceCodeTwoFactor                   = "19" // Two-factor authentication service
	ServiceCodeBank                        = "20" // Bank service
	ServiceCodeBankAccount                 = "21" // Bank account service
)

// Case codes identify the outcome within a service (2 digits, "01"–"99").
// Ranges: success 01–10, validation 11–20, auth 21–30, not-found 31–44,
// business 45–54, server 55–64, conflict 65–69, email/phone/2FA/bank 70–99.
const (
	// Success cases (01-10)
	CaseCodeSuccess            = "01" // General success
	CaseCodeCreated            = "02" // Resource created
	CaseCodeUpdated            = "03" // Resource updated
	CaseCodeDeleted            = "04" // Resource deleted
	CaseCodeRetrieved          = "05" // Resource retrieved
	CaseCodeListRetrieved      = "06" // List retrieved
	CaseCodeLoginSuccess       = "07" // Login successful
	CaseCodeLogoutSuccess      = "08" // Logout successful
	CaseCodePasswordChanged    = "09" // Password changed
	CaseCodeOperationCompleted = "10" // Operation completed

	// Validation errors (11-20)
	CaseCodeValidationError  = "11" // General validation error
	CaseCodeRequiredField    = "12" // Required field missing
	CaseCodeInvalidFormat    = "13" // Invalid format
	CaseCodeInvalidValue     = "14" // Invalid value
	CaseCodeDuplicateEntry   = "15" // Duplicate entry
	CaseCodeInvalidEmail     = "16" // Invalid email format
	CaseCodeInvalidPassword  = "17" // Invalid password
	CaseCodePasswordTooShort = "18" // Password too short
	CaseCodeInvalidDate      = "19" // Invalid date format
	CaseCodeInvalidRange     = "20" // Invalid range

	// Authentication errors (21-30)
	CaseCodeUnauthorized       = "21" // Unauthorized access
	CaseCodeInvalidToken       = "22" // Invalid token
	CaseCodeTokenExpired       = "23" // Token expired
	CaseCodeInvalidCredentials = "24" // Invalid credentials
	CaseCodeAccountLocked      = "25" // Account locked
	CaseCodeAccountDisabled    = "26" // Account disabled
	CaseCodePermissionDenied   = "27" // Permission denied
	CaseCodeSessionExpired     = "28" // Session expired
	CaseCodeTwoFactorRequired  = "29" // Two-factor authentication required
	CaseCodeInvalidOTP         = "30" // Invalid OTP

	// Not found errors (31-40)
	CaseCodeNotFound                            = "31" // Resource not found
	CaseCodeUserNotFound                        = "32" // User not found
	CaseCodeAdminNotFound                       = "33" // Admin not found
	CaseCodeMerchantNotFound                    = "34" // Merchant not found
	CaseCodeTransactionNotFound                 = "35" // Transaction not found
	CaseCodeSettingNotFound                     = "36" // Setting not found
	CaseCodeRoleNotFound                        = "37" // Role not found
	CaseCodeNotificationTemplateNotFound        = "38" // Notification template not found
	CaseCodeNotificationTemplateChannelNotFound = "39" // Notification template channel not found
	CaseCodeNotificationNotFound                = "40" // Notification not found
	CaseCodeMethodNotFound                      = "41" // Method not found
	CaseCodeRouteNotFound                       = "42" // Route not found
	CaseCodeResourceNotFound                    = "43" // General resource not found
	CaseCodeApiKeyNotFound                      = "44" // API key not found

	// Business logic errors (45-53)
	CaseCodeInsufficientBalance = "45" // Insufficient balance
	CaseCodeInvalidAmount       = "46" // Invalid amount
	CaseCodeTransactionFailed   = "47" // Transaction failed
	CaseCodeLimitExceeded       = "48" // Limit exceeded
	CaseCodeInvalidStatus       = "49" // Invalid status
	CaseCodeOperationNotAllowed = "50" // Operation not allowed
	CaseCodeAlreadyProcessed    = "51" // Already processed
	CaseCodePendingTransaction  = "52" // Pending transaction
	CaseCodeExpiredTransaction  = "53" // Expired transaction
	CaseCodeInvalidCurrency     = "54" // Invalid currency

	// Server errors (55-63)
	CaseCodeInternalError        = "55" // Internal server error
	CaseCodeDatabaseError        = "56" // Database error
	CaseCodeExternalServiceError = "57" // External service error
	CaseCodeTimeout              = "58" // Request timeout
	CaseCodeServiceUnavailable   = "59" // Service unavailable
	CaseCodeMaintenance          = "60" // Under maintenance
	CaseCodeRateLimitExceeded    = "61" // Rate limit exceeded
	CaseCodeConfigurationError   = "62" // Configuration error
	CaseCodeEncryptionError      = "63" // Encryption error
	CaseCodeDecryptionError      = "64" // Decryption error

	// Conflict errors (65-69)
	CaseCodeConflict               = "65" // General conflict
	CaseCodeResourceExists         = "66" // Resource already exists
	CaseCodeConcurrentModification = "67" // Concurrent modification
	CaseCodeVersionMismatch        = "68" // Version mismatch
	CaseCodeStateConflict          = "69" // State conflict

	// Email change (70-74)
	CaseCodeEmailChangeRequested    = "70" // Email change OTP sent
	CaseCodeEmailChangeVerified     = "71" // Email change confirmed
	CaseCodeEmailChangeCancelled    = "72" // Email change cancelled
	CaseCodeEmailAlreadyUsed        = "73" // Email already registered to another account
	CaseCodeEmailVerificationFailed = "74" // Email OTP verification failed

	// Phone change (75-79)
	CaseCodePhoneChangeRequested    = "75" // Phone change OTP sent
	CaseCodePhoneChangeVerified     = "76" // Phone change confirmed
	CaseCodePhoneChangeCancelled    = "77" // Phone change cancelled
	CaseCodePhoneAlreadyUsed        = "78" // Phone already registered to another account
	CaseCodePhoneVerificationFailed = "79" // Phone OTP verification failed

	// Two-factor authentication (80-89)
	CaseCode2FAEnabled             = "80" // 2FA enabled
	CaseCode2FADisabled            = "81" // 2FA disabled
	CaseCode2FASetupInitiated      = "82" // 2FA setup started (secret + QR generated)
	CaseCode2FASetupVerified       = "83" // 2FA setup confirmed with valid OTP
	CaseCode2FAInvalidCode         = "84" // 2FA code invalid or expired
	CaseCode2FACodeRequired        = "85" // 2FA code missing from request
	CaseCode2FANotEnabled          = "86" // 2FA not enabled on account
	CaseCode2FAAlreadyEnabled      = "87" // 2FA already enabled
	CaseCode2FARecoveryCodeUsed    = "88" // 2FA bypass via recovery code
	CaseCode2FARecoveryCodeInvalid = "89" // Recovery code invalid or already used

	// Bank (90-94)
	CaseCodeBankNotFound      = "90" // Bank not found
	CaseCodeBankInactive      = "91" // Bank is inactive / not available
	CaseCodeBankAlreadyExists = "92" // Bank already registered
	CaseCodeBankCreated       = "93" // Bank created
	CaseCodeBankUpdated       = "94" // Bank updated

	// Bank account (95-99)
	CaseCodeBankAccountNotFound      = "95" // Bank account not found
	CaseCodeBankAccountInvalid       = "96" // Bank account number invalid or unverifiable
	CaseCodeBankAccountAlreadyExists = "97" // Bank account already registered
	CaseCodeBankAccountCreated       = "98" // Bank account created
	CaseCodeBankAccountDeleted       = "99" // Bank account deleted
)

// BuildResponseCode composes the 7-digit API code from HTTP status, service, and case.
// Format: HTTP_STATUS (3) + SERVICE_CODE (2) + CASE_CODE (2).
//
// Parameters:
//   - httpStatus: HTTP status used for both the wire status and the first three digits
//   - serviceCode: two-digit service constant (for example ServiceCodeWithdrawal)
//   - caseCode: two-digit case constant (for example CaseCodeSuccess)
//
// Callers must pass numeric digit strings; non-digit characters produce undefined codes.
//
// Example:
//
//	code := BuildResponseCode(http.StatusCreated, ServiceCodeWithdrawal, CaseCodeSuccess)
//	// code == 2010301
func BuildResponseCode(httpStatus int, serviceCode, caseCode string) int {
	svc, _ := strconv.Atoi(serviceCode)
	cs, _ := strconv.Atoi(caseCode)
	return httpStatus*10000 + svc*100 + cs
}

// twoDigitStrings caches "00"…"99" so ParseResponseCode can return service/case
// without per-call string allocation (bench: 0 allocs/op vs Itoa+slice).
var twoDigitStrings = func() (a [100]string) {
	for i := range a {
		a[i] = string([]byte{byte('0' + i/10), byte('0' + i%10)})
	}
	return a
}()

// ParseResponseCode splits a composite code into HTTP status, service code, and case code.
// Codes with fewer than 7 decimal digits return zero values for all three results.
// Longer codes use the leading 7 digits (same as taking codeStr[:7] on the decimal form).
// Leading zeros in the numeric form are not preserved beyond the digit extraction
// (for example status 200 is returned as 200).
func ParseResponseCode(code int) (httpStatus int, serviceCode, caseCode string) {
	if code < 1_000_000 {
		return 0, "", ""
	}
	// Keep leading 7 digits when the value has more than 7 decimal digits.
	for code >= 10_000_000 {
		code /= 10
	}

	httpStatus = code / 10_000
	serviceCode = twoDigitStrings[(code/100)%100]
	caseCode = twoDigitStrings[code%100]
	return httpStatus, serviceCode, caseCode
}
