package auth

import (
	"strings"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
)

const (
	AuthorizationKey = "authorization" // Authorization context metadata key.
	bearerPrefix     = "Bearer "
)

var (
	ErrMissingMetadata                = "missing metadata"
	ErrAuthorizationMetadataNotFound  = AuthorizationKey + " metadata not found"
	ErrMalformedAuthorizationMetadata = AuthorizationKey + " metadata is not a Bearer token"
)

// Authorizer provides convenience methods for dealing with JWT tokens inside grpc metadata.
// All errors returned by implemetations should be micro framework compatibile.
type Authorizer interface {
	// WithAuthorization creates new context with authorization metadata set with token.
	// When it returns error it is with InternalServerError status.
	WithAuthorization(ctx context.Context, username string) (context.Context, error)
	// Extracts username from authorization metadata. When it returns error it is with
	// Unauthorized status.
	FromContext(context.Context) (username string, err error)
}

// NewAuthorizer creates Authorizer which uses JWT tokens and serviceID as micro errors' ID.
func NewAuthorizer(serviceID string) Authorizer {
	return authorizer{
		serviceID: serviceID,
	}
}

type authorizer struct {
	serviceID string
}

func (a authorizer) WithAuthorization(ctx context.Context, username string) (context.Context, error) {
	token, err := CreateToken(username)
	if err != nil {
		return ctx, errors.InternalServerError(a.serviceID, err.Error())
	}
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata(map[string]string{
			AuthorizationKey: bearerPrefix + token,
		})
	}
	return metadata.NewContext(ctx, md), nil
}

func (a authorizer) FromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return "", errors.Unauthorized(a.serviceID, ErrMissingMetadata)
	}
	v, ok := md[AuthorizationKey]
	if !ok {
		return "", errors.Unauthorized(a.serviceID, ErrAuthorizationMetadataNotFound)
	}
	token := strings.TrimPrefix(v, bearerPrefix)
	if token == v {
		return "", errors.Unauthorized(a.serviceID, ErrMalformedAuthorizationMetadata)
	}
	user, err := ReadUser(token)
	if err != nil {
		return "", errors.Unauthorized(a.serviceID, err.Error())
	}
	return user, nil
}
