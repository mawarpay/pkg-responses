package response

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------------------
// BuildResponseCode
// ---------------------------------------------------------------------------

func TestBuildResponseCode(t *testing.T) {
	tests := []struct {
		name        string
		httpStatus  int
		serviceCode string
		caseCode    string
		want        int
	}{
		// Core service codes
		{
			name:        "200 common success",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeCommon,
			caseCode:    CaseCodeSuccess,
			want:        2000001,
		},
		{
			name:        "201 withdrawal created",
			httpStatus:  http.StatusCreated,
			serviceCode: ServiceCodeWithdrawal,
			caseCode:    CaseCodeCreated,
			want:        2010302,
		},
		{
			name:        "404 user not found",
			httpStatus:  http.StatusNotFound,
			serviceCode: ServiceCodeUser,
			caseCode:    CaseCodeNotFound,
			want:        4040431,
		},
		{
			name:        "422 validation error",
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodeCommon,
			caseCode:    CaseCodeValidationError,
			want:        4220011,
		},
		{
			name:        "500 internal error",
			httpStatus:  http.StatusInternalServerError,
			serviceCode: ServiceCodeCommon,
			caseCode:    CaseCodeInternalError,
			want:        5000055,
		},

		// HTTP status edge cases
		{
			name:        "single-digit http status",
			httpStatus:  5,
			serviceCode: "00",
			caseCode:    "01",
			want:        50001,
		},
		{
			name:        "two-digit http status",
			httpStatus:  10,
			serviceCode: "00",
			caseCode:    "01",
			want:        100001,
		},
		{
			name:        "three-digit http status",
			httpStatus:  100,
			serviceCode: "00",
			caseCode:    "01",
			want:        1000001,
		},

		// Auth service
		{
			name:        "401 invalid token",
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeAuth,
			caseCode:    CaseCodeInvalidToken,
			want:        4010122,
		},
		{
			name:        "401 token expired",
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeAuth,
			caseCode:    CaseCodeTokenExpired,
			want:        4010123,
		},
		{
			name:        "401 2FA required",
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeAuth,
			caseCode:    CaseCodeTwoFactorRequired,
			want:        4010129,
		},
		{
			name:        "401 invalid OTP",
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeAuth,
			caseCode:    CaseCodeInvalidOTP,
			want:        4010130,
		},

		// Email service (new)
		{
			name:        "200 email change requested",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailChangeRequested,
			want:        2001870,
		},
		{
			name:        "200 email change verified",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailChangeVerified,
			want:        2001871,
		},
		{
			name:        "200 email change cancelled",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailChangeCancelled,
			want:        2001872,
		},
		{
			name:        "409 email already used",
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailAlreadyUsed,
			want:        4091873,
		},
		{
			name:        "422 email verification failed",
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailVerificationFailed,
			want:        4221874,
		},

		// Phone service (new)
		{
			name:        "200 phone change requested",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodePhone,
			caseCode:    CaseCodePhoneChangeRequested,
			want:        2001775,
		},
		{
			name:        "200 phone change verified",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodePhone,
			caseCode:    CaseCodePhoneChangeVerified,
			want:        2001776,
		},
		{
			name:        "200 phone change cancelled",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodePhone,
			caseCode:    CaseCodePhoneChangeCancelled,
			want:        2001777,
		},
		{
			name:        "409 phone already used",
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodePhone,
			caseCode:    CaseCodePhoneAlreadyUsed,
			want:        4091778,
		},
		{
			name:        "422 phone verification failed",
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodePhone,
			caseCode:    CaseCodePhoneVerificationFailed,
			want:        4221779,
		},

		// Two-factor authentication service (new)
		{
			name:        "200 2FA enabled",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FAEnabled,
			want:        2001980,
		},
		{
			name:        "200 2FA disabled",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FADisabled,
			want:        2001981,
		},
		{
			name:        "200 2FA setup initiated",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FASetupInitiated,
			want:        2001982,
		},
		{
			name:        "200 2FA setup verified",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FASetupVerified,
			want:        2001983,
		},
		{
			name:        "401 2FA invalid code",
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FAInvalidCode,
			want:        4011984,
		},
		{
			name:        "422 2FA code required",
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FACodeRequired,
			want:        4221985,
		},
		{
			name:        "422 2FA not enabled",
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FANotEnabled,
			want:        4221986,
		},
		{
			name:        "409 2FA already enabled",
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FAAlreadyEnabled,
			want:        4091987,
		},
		{
			name:        "200 2FA recovery code used",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FARecoveryCodeUsed,
			want:        2001988,
		},
		{
			name:        "401 2FA recovery code invalid",
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FARecoveryCodeInvalid,
			want:        4011989,
		},

		// Bank service (new)
		{
			name:        "404 bank not found",
			httpStatus:  http.StatusNotFound,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankNotFound,
			want:        4042090,
		},
		{
			name:        "422 bank inactive",
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankInactive,
			want:        4222091,
		},
		{
			name:        "409 bank already exists",
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankAlreadyExists,
			want:        4092092,
		},
		{
			name:        "201 bank created",
			httpStatus:  http.StatusCreated,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankCreated,
			want:        2012093,
		},
		{
			name:        "200 bank updated",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankUpdated,
			want:        2002094,
		},

		// Bank account service (new)
		{
			name:        "404 bank account not found",
			httpStatus:  http.StatusNotFound,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountNotFound,
			want:        4042195,
		},
		{
			name:        "422 bank account invalid",
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountInvalid,
			want:        4222196,
		},
		{
			name:        "409 bank account already exists",
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountAlreadyExists,
			want:        4092197,
		},
		{
			name:        "201 bank account created",
			httpStatus:  http.StatusCreated,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountCreated,
			want:        2012198,
		},
		{
			name:        "200 bank account deleted",
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountDeleted,
			want:        2002199,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildResponseCode(tt.httpStatus, tt.serviceCode, tt.caseCode)
			assert.Equal(t, tt.want, got,
				"BuildResponseCode(%d, %q, %q)", tt.httpStatus, tt.serviceCode, tt.caseCode)
		})
	}
}

// ---------------------------------------------------------------------------
// ParseResponseCode
// ---------------------------------------------------------------------------

func TestParseResponseCode(t *testing.T) {
	tests := []struct {
		name        string
		code        int
		wantHTTP    int
		wantService string
		wantCase    string
	}{
		{
			name:        "200 common success",
			code:        2000001,
			wantHTTP:    200,
			wantService: "00",
			wantCase:    "01",
		},
		{
			name:        "201 withdrawal created",
			code:        2010302,
			wantHTTP:    201,
			wantService: "03",
			wantCase:    "02",
		},
		{
			name:        "404 user not found",
			code:        4040431,
			wantHTTP:    404,
			wantService: "04",
			wantCase:    "31",
		},
		{
			name:        "422 validation error",
			code:        4220011,
			wantHTTP:    422,
			wantService: "00",
			wantCase:    "11",
		},
		{
			name:        "500 internal error",
			code:        5000055,
			wantHTTP:    500,
			wantService: "00",
			wantCase:    "55",
		},

		// New service codes
		{
			name:        "email change requested",
			code:        2001870,
			wantHTTP:    200,
			wantService: "18",
			wantCase:    "70",
		},
		{
			name:        "email already used",
			code:        4091873,
			wantHTTP:    409,
			wantService: "18",
			wantCase:    "73",
		},
		{
			name:        "phone change verified",
			code:        2001776,
			wantHTTP:    200,
			wantService: "17",
			wantCase:    "76",
		},
		{
			name:        "phone already used",
			code:        4091778,
			wantHTTP:    409,
			wantService: "17",
			wantCase:    "78",
		},
		{
			name:        "2FA enabled",
			code:        2001980,
			wantHTTP:    200,
			wantService: "19",
			wantCase:    "80",
		},
		{
			name:        "2FA invalid code",
			code:        4011984,
			wantHTTP:    401,
			wantService: "19",
			wantCase:    "84",
		},
		{
			name:        "2FA recovery invalid",
			code:        4011989,
			wantHTTP:    401,
			wantService: "19",
			wantCase:    "89",
		},
		// Bank
		{
			name:        "bank not found",
			code:        4042090,
			wantHTTP:    404,
			wantService: "20",
			wantCase:    "90",
		},
		{
			name:        "bank created",
			code:        2012093,
			wantHTTP:    201,
			wantService: "20",
			wantCase:    "93",
		},
		// Bank account
		{
			name:        "bank account not found",
			code:        4042195,
			wantHTTP:    404,
			wantService: "21",
			wantCase:    "95",
		},
		{
			name:        "bank account created",
			code:        2012198,
			wantHTTP:    201,
			wantService: "21",
			wantCase:    "98",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpStatus, serviceCode, caseCode := ParseResponseCode(tt.code)
			assert.Equal(t, tt.wantHTTP, httpStatus, "httpStatus for code %d", tt.code)
			assert.Equal(t, tt.wantService, serviceCode, "serviceCode for code %d", tt.code)
			assert.Equal(t, tt.wantCase, caseCode, "caseCode for code %d", tt.code)
		})
	}
}

func TestParseResponseCode_ShortCode(t *testing.T) {
	// Codes shorter than 7 digits must return zero values.
	httpStatus, serviceCode, caseCode := ParseResponseCode(200001)
	assert.Equal(t, 0, httpStatus)
	assert.Empty(t, serviceCode)
	assert.Empty(t, caseCode)
}

func TestParseResponseCode_Zero(t *testing.T) {
	httpStatus, serviceCode, caseCode := ParseResponseCode(0)
	assert.Equal(t, 0, httpStatus)
	assert.Empty(t, serviceCode)
	assert.Empty(t, caseCode)
}

func TestParseResponseCode_LongCode(t *testing.T) {
	// Longer than 7 digits: use leading 7 (12010301 → 120 / 10 / 30).
	httpStatus, serviceCode, caseCode := ParseResponseCode(12010301)
	assert.Equal(t, 120, httpStatus)
	assert.Equal(t, "10", serviceCode)
	assert.Equal(t, "30", caseCode)
}

// ---------------------------------------------------------------------------
// Round-trip
// ---------------------------------------------------------------------------

func TestBuildParseRoundTrip(t *testing.T) {
	cases := []struct {
		httpStatus  int
		serviceCode string
		caseCode    string
	}{
		// Existing codes
		{
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeCommon,
			caseCode:    CaseCodeSuccess,
		},
		{
			httpStatus:  http.StatusCreated,
			serviceCode: ServiceCodeWithdrawal,
			caseCode:    CaseCodeCreated,
		},
		{
			httpStatus:  http.StatusNotFound,
			serviceCode: ServiceCodeUser,
			caseCode:    CaseCodeNotFound,
		},
		{
			httpStatus:  http.StatusInternalServerError,
			serviceCode: ServiceCodeCommon,
			caseCode:    CaseCodeInternalError,
		},
		// Email
		{
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailChangeRequested,
		},
		{
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailChangeVerified,
		},
		{
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailAlreadyUsed,
		},
		{
			httpStatus:  http.StatusUnprocessableEntity,
			serviceCode: ServiceCodeEmail,
			caseCode:    CaseCodeEmailVerificationFailed,
		},
		// Phone
		{
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodePhone,
			caseCode:    CaseCodePhoneChangeRequested,
		},
		{
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodePhone,
			caseCode:    CaseCodePhoneAlreadyUsed,
		},
		// 2FA
		{
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FAEnabled,
		},
		{
			httpStatus:  http.StatusOK,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FASetupInitiated,
		},
		{
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FAInvalidCode,
		},
		{
			httpStatus:  http.StatusUnauthorized,
			serviceCode: ServiceCodeTwoFactor,
			caseCode:    CaseCode2FARecoveryCodeInvalid,
		},
		// Bank
		{
			httpStatus:  http.StatusNotFound,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankNotFound,
		},
		{
			httpStatus:  http.StatusCreated,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankCreated,
		},
		{
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodeBank,
			caseCode:    CaseCodeBankAlreadyExists,
		},
		// Bank account
		{
			httpStatus:  http.StatusNotFound,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountNotFound,
		},
		{
			httpStatus:  http.StatusCreated,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountCreated,
		},
		{
			httpStatus:  http.StatusConflict,
			serviceCode: ServiceCodeBankAccount,
			caseCode:    CaseCodeBankAccountAlreadyExists,
		},
	}

	for _, c := range cases {
		code := BuildResponseCode(c.httpStatus, c.serviceCode, c.caseCode)
		gotHTTP, gotService, gotCase := ParseResponseCode(code)
		assert.Equal(
			t, c.httpStatus, gotHTTP,
			"httpStatus round-trip for (%d,%s,%s)", c.httpStatus, c.serviceCode, c.caseCode,
		)
		assert.Equal(
			t, c.serviceCode, gotService,
			"serviceCode round-trip for (%d,%s,%s)", c.httpStatus, c.serviceCode, c.caseCode,
		)
		assert.Equal(
			t, c.caseCode, gotCase,
			"caseCode round-trip for (%d,%s,%s)", c.httpStatus, c.serviceCode, c.caseCode,
		)
	}
}

// ---------------------------------------------------------------------------
// Service code constant completeness
// ---------------------------------------------------------------------------

func TestServiceCodeConstants(t *testing.T) {
	// Verify every service code has a unique value so no two constants collide.
	seen := make(map[string]string)
	codes := map[string]string{
		"Common":                      ServiceCodeCommon,
		"Auth":                        ServiceCodeAuth,
		"Transaction":                 ServiceCodeTransaction,
		"Withdrawal":                  ServiceCodeWithdrawal,
		"User":                        ServiceCodeUser,
		"Admin":                       ServiceCodeAdmin,
		"Merchant":                    ServiceCodeMerchant,
		"Setting":                     ServiceCodeSetting,
		"Role":                        ServiceCodeRole,
		"Permission":                  ServiceCodePermission,
		"NotificationTemplate":        ServiceCodeNotificationTemplate,
		"NotificationTemplateChannel": ServiceCodeNotificationTemplateChannel,
		"Notification":                ServiceCodeNotification,
		"IPWhitelist":                 ServiceCodeIPWhitelist,
		"ApiKey":                      ServiceCodeApiKey,
		"Deposit":                     ServiceCodeDeposit,
		"Wallet":                      ServiceCodeWallet,
		"Phone":                       ServiceCodePhone,
		"Email":                       ServiceCodeEmail,
		"TwoFactor":                   ServiceCodeTwoFactor,
		"Bank":                        ServiceCodeBank,
		"BankAccount":                 ServiceCodeBankAccount,
	}
	for name, code := range codes {
		if prev, exists := seen[code]; exists {
			t.Errorf("ServiceCode collision: %q and %q both have value %q", name, prev, code)
		}
		seen[code] = name
	}
}
