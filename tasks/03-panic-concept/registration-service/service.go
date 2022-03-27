package registrationservice

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("email is invalid")
	ErrWeakPassword = errors.New("password is too weak")
)

type Service struct {
	emailExpr     string
	passwordExprs []string

	emailRegexp     *regexp.Regexp
	passwordRegexps []*regexp.Regexp
}

func NewService(emailExpr string, passwordExprs []string) (*Service, error) {
	if passwordExprs == nil {
		return nil, errors.New("no password expressions")
	}
	return &Service{
		emailExpr:     emailExpr,
		passwordExprs: passwordExprs,
	}, nil
}

func (s *Service) SignUp(email, password string) error {
	if err := s.validateEmail(email); err != nil {
		return fmt.Errorf("validate email: %w", err)
	}

	if err := s.validatePassword(password); err != nil {
		return fmt.Errorf("validate password: %w", err)
	}

	// Other sign up code...
	return nil
}

func (s *Service) validateEmail(email string) error {
	if s.emailRegexp == nil {
		s.emailRegexp = regexp.MustCompile(s.emailExpr)
	}

	if !s.emailRegexp.MatchString(email) {
		return fmt.Errorf("%w: %q", ErrInvalidEmail, email)
	}
	return nil
}

func (s *Service) validatePassword(password string) error {
	if len(s.passwordRegexps) == 0 {
		s.passwordRegexps = make([]*regexp.Regexp, len(s.passwordExprs))
		for i, v := range s.passwordExprs {
			s.passwordRegexps[i] = regexp.MustCompile(v)
		}
	}

	for _, r := range s.passwordRegexps {
		if !r.MatchString(password) {
			return fmt.Errorf("%w: %q", ErrWeakPassword, password)
		}
	}
	return nil
}
