package auth

import (
	"log"
	"reflect"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

// Typical usage.
func TestCreateToken(t *testing.T) {
	token, err := CreateToken("pawel")
	if err != nil {
		t.Fatalf("unexpected CreateToken error = %s", err)
	}
	user, err := ReadUser(token)
	if err != nil {
		t.Fatalf("unexpected ReadUser error = %s", err)
	}
	if user != "pawel" {
		log.Fatalf("get %v, want pawel", user)
	}
}

// All tests below lock behavious when something goes wrong.

func TestReadUser_BadSecret(t *testing.T) {
	oldSecret := secret
	defer func() {
		secret = oldSecret
	}()
	token, err := CreateToken("pawel")
	if err != nil {
		t.Fatalf("unexpected CreatToken error = %s", err)
	}

	// Change the secret.
	secret = []byte("bad secret")
	_, err = ReadUser(token)
	werr := AuthError{Err: ErrParseToken}
	if !reflect.DeepEqual(clearedDetails(err), werr) {
		t.Fatalf("got %v, want %v", err, werr)
	}
}

func TestReadUser_BadToken(t *testing.T) {
	tests := []struct {
		Token string
		Error error
	}{
		{
			Token: "bad token",
			Error: AuthError{Err: ErrParseToken},
		},
		{
			// Example RS256 signed token.
			Token: `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.EkN-DOsnsuRjRO6BxXemmJDm3HbxrbRzXglbN2S4sOkopdU4IsDxTI8jO19W_A4K8ZPJijNLis4EZsHeY559a4DFOd50_OqgHGuERTqYZyuhtF39yxJPAjUESwxk2J5k_4zM3O-vtd1Ghyo4IbqKKSy6J9mTniYJPenn5-HIirE`,
			Error: AuthError{Err: ErrSigningMethod},
		},
	}
	for i, tt := range tests {
		_, err := ReadUser(tt.Token)
		if !reflect.DeepEqual(clearedDetails(err), tt.Error) {
			t.Errorf("#%d: got %v, want %v", i, err, tt.Error)
		}
	}
}

func TestReadUser_BadClaims(t *testing.T) {
	tests := []struct {
		Claims jwt.MapClaims
		Error  error
	}{
		{
			Claims: jwt.MapClaims{},
			Error:  AuthError{Err: ErrMissingClaim},
		},
		{
			Claims: jwt.MapClaims{"username": 1},
			Error:  AuthError{Err: ErrClaimType},
		},
	}
	for i, tt := range tests {
		token, err := createToken(tt.Claims)
		if err != nil {
			t.Errorf("#%d: unexpected SignedString error = %v", i, err)
		}
		_, err = ReadUser(token)
		if !reflect.DeepEqual(clearedDetails(err), tt.Error) {
			t.Errorf("#%d: got %v, want %v", i, err, tt.Error)
		}
	}
}

func clearedDetails(err error) error {
	if auth, ok := err.(AuthError); ok {
		auth.Details = nil
		return auth
	}
	return err
}
