package registrationservice

import (
	"regexp/syntax"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	brokenEmailRegexp = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"
	validEmailRegexp  = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"
)

var (
	invalidPasswordRegexps = []string{
		"^(?=(.*[a-z]){3,})(?=(.*[A-Z]){2,})(?=(.*[0-9]){2,})(?=(.*[!@#$%^&*()\\-__+.]){1,}).{8,}$",
	}
	validPasswordRegexps = []string{
		"(.*[a-z]){3,}",               // Минимум 3 строчных буквы.
		"(.*[A-Z]){2,}",               // Минимум 2 заглавных буквы.
		"(.*[0-9]){2,}",               // Минимум 2 цифры.
		"(.*[!@#$%^&*()\\-__+.]){1,}", // Минимум один специальный символ !@#$%^&*()\-__+..
		"^.{8,}$",                     // Минимум 8 знаков длиной.
	}
)

func TestService_New(t *testing.T) {
	cases := []struct {
		name          string
		emailExpr     string
		passwordExprs []string
		errExpected   error
	}{
		{
			name:          "valid email regexp  valid password regexps",
			emailExpr:     validEmailRegexp,
			passwordExprs: validPasswordRegexps,
			errExpected:   nil,
		},
		{
			name:          "valid email regexp  invalid password regexps",
			emailExpr:     validEmailRegexp,
			passwordExprs: invalidPasswordRegexps,
			errExpected:   &syntax.Error{Code: syntax.ErrInvalidPerlOp, Expr: "(?="},
		},
		{
			name:          "invalid email regexp  valid password regexps",
			emailExpr:     brokenEmailRegexp,
			passwordExprs: validPasswordRegexps,
			errExpected:   &syntax.Error{Code: syntax.ErrUnexpectedParen, Expr: brokenEmailRegexp},
		},
		{
			name:          "invalid email regexp  invalid password regexps",
			emailExpr:     brokenEmailRegexp,
			passwordExprs: invalidPasswordRegexps,
			errExpected:   &syntax.Error{Code: syntax.ErrUnexpectedParen, Expr: brokenEmailRegexp},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewService(tt.emailExpr, tt.passwordExprs)
			if tt.errExpected == nil {
				require.NoError(t, err)
				assert.NotNil(t, s)
				return
			}

			errSyntax := new(syntax.Error)
			require.ErrorAs(t, err, &errSyntax)
			assert.EqualValues(t, errSyntax, tt.errExpected)
			assert.Nil(t, s)
		})
	}
}

func TestService_New_NoPasswordRegexps(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		s, err := NewService(validEmailRegexp, nil)
		require.Error(t, err)
		assert.Nil(t, s)
	})

	t.Run("empty slice", func(t *testing.T) {
		s, err := NewService(validEmailRegexp, []string{})
		require.Error(t, err)
		assert.Nil(t, s)
	})
}

func TestService_SignUp(t *testing.T) {
	cases := []struct {
		name        string
		email       string
		password    string
		errExpected error
	}{
		{
			name:        "valid email  valid password",
			email:       "sensei@golang-ninja.ru",
			password:    "AB!12cdef",
			errExpected: nil,
		},
		// Минимум 3 строчных буквы.
		{
			name:        "valid email  invalid password 1",
			email:       "sensei@golang-ninja.ru",
			password:    "AB!12c..f",
			errExpected: ErrWeakPassword,
		},
		// Минимум 2 заглавных буквы.
		{
			name:        "valid email  invalid password 2",
			email:       "sensei@golang-ninja.ru",
			password:    "Ab!12cdef",
			errExpected: ErrWeakPassword,
		},
		// Минимум 2 цифры.
		{
			name:        "valid email  invalid password 3",
			email:       "sensei@golang-ninja.ru",
			password:    "AB!1Ecdef",
			errExpected: ErrWeakPassword,
		},
		// Минимум один специальный символ !@#$%^&*()\-__+..
		{
			name:        "valid email  invalid password 4",
			email:       "sensei@golang-ninja.ru",
			password:    "ABno12cdefABno12cdef",
			errExpected: ErrWeakPassword,
		},
		// Минимум 8 знаков длиной.
		{
			name:        "valid email  invalid password 5",
			email:       "sensei@golang-ninja.ru",
			password:    "AB!12cd",
			errExpected: ErrWeakPassword,
		},
		{
			name:        "invalid email  valid password",
			email:       "senseigolang-ninja.ru",
			password:    "AB!12cdef",
			errExpected: ErrInvalidEmail,
		},
		{
			name:        "invalid email  invalid password",
			email:       "senseigolang-ninja.ru",
			password:    "12345",
			errExpected: ErrInvalidEmail,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewService(validEmailRegexp, validPasswordRegexps)
			require.NoError(t, err)

			err = s.SignUp(tt.email, tt.password)
			require.ErrorIs(t, err, tt.errExpected)
		})
	}
}
