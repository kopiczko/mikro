package auth

import (
	"errors"
	"fmt"
	"reflect"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte("do not tell anybody")

var (
	ErrClaimType     = errors.New("unexpected claim type")
	ErrInvalidToken  = errors.New("invalid token")
	ErrMissingClaim  = errors.New("missing claim")
	ErrParseToken    = errors.New("cannot parse token")
	ErrSigningMethod = errors.New("unexpected signing method")
)

type AuthError struct {
	Err     error
	Details interface{}
}

func (e AuthError) Error() string {
	msg := e.Err.Error()
	if e.Details != nil && e.Details != "" {
		msg = fmt.Sprintf("%s (%v)", msg, e.Details)
	}
	return msg
}

// CreateToken creates a signed token with username claim set.
func CreateToken(user string) (string, error) {
	return createToken(jwt.MapClaims{"username": user})
}

func createToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// readUser reads username claim from the token.
func readUser(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check singing method.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			alg := interface{}("")
			if h, ok := token.Header["alg"]; ok {
				alg = h
			}
			return nil, AuthError{Err: ErrSigningMethod, Details: alg}
		}
		return secret, nil
	})
	// Check parsing.
	if err != nil {
		if jwtErr, ok := err.(*jwt.ValidationError); ok {
			if _, ok := jwtErr.Inner.(AuthError); ok && jwtErr.Inner != nil {
				return "", jwtErr.Inner
			}
		}
		return "", AuthError{Err: ErrParseToken, Details: err}
	}
	// Extract claims.
	claims := t.Claims.(jwt.MapClaims)
	userClaim, ok := claims["username"]
	if !ok {
		return "", AuthError{Err: ErrMissingClaim, Details: "username"}
	}
	user, ok := userClaim.(string)
	if !ok {
		details := fmt.Sprintf("username: want string, got %s", reflect.TypeOf(userClaim))
		return "", AuthError{Err: ErrClaimType, Details: details}
	}
	return user, nil
}
