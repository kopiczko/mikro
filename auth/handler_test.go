package auth

import (
	"testing"

	"github.com/kopiczko/mikro/auth/authpb"
)

func TestDBAccessorHandler(t *testing.T) {
	var _ authpb.AuthHandler = new(Auth) // Auth should be created with New
}
