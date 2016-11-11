package auth

import (
	"fmt"
	"testing"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"

	"golang.org/x/net/context"
)

func TestAuhtorizer(t *testing.T) {
	a := NewAuthorizer("test.service")
	ctx, err := a.WithAuthorization(context.TODO(), "pawel")
	if err != nil {
		t.Fatalf("unexpected WithAuthorization error = %v", err)
	}
	user, err := a.FromContext(ctx)
	if err != nil {
		t.Fatalf("unexpected FromContext error = %v", err)
	}
	if user != "pawel" {
		t.Errorf("got %s, want pawel", user)
	}
}

func TestAuhtorizerErrors(t *testing.T) {
	a := NewAuthorizer("test.service")
	tests := []struct {
		MD metadata.Metadata
	}{
		{MD: nil},
		{MD: md()},
		{MD: md(AuthorizationKey, "")},
		{MD: md(AuthorizationKey, "not_a_valid_bearer_token")},
		{MD: md(AuthorizationKey, bearerPrefix+"a")},
	}
	for i, tt := range tests {
		ctx := context.TODO()
		if tt.MD != nil {
			ctx = metadata.NewContext(ctx, tt.MD)
		}
		_, err := a.FromContext(ctx)
		if err == nil {
			t.Errorf("#%d: got %v, want nil", i, err)
		}
		code := errors.Parse(err.Error()).Code
		if code != 401 {
			t.Errorf("#%d: got %v, want 404", i, code)
		}
	}
}

func md(pairs ...string) metadata.Metadata {
	if len(pairs)%2 == 1 {
		panic(fmt.Sprintf("md got the odd number of input pairs for metadata: %d", len(pairs)))
	}
	md := metadata.Metadata{}
	var k string
	for i, s := range pairs {
		if i%2 == 0 {
			k = s
			continue
		}
		md[k] = s
	}
	return md
}
