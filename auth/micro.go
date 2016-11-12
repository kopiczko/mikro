package auth

import (
	"strings"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
)

const (
	AuthorizationKey = "Authorization" // Authorization context metadata key.
	bearerPrefix     = "Bearer "
)

var (
	ErrMissingMetadata                = "missing metadata"
	ErrAuthorizationMetadataNotFound  = AuthorizationKey + " metadata not found"
	ErrMalformedAuthorizationMetadata = AuthorizationKey + " metadata is not a Bearer token"
)

// WithAuthorization creates new context with authorization metadata filled with JWT token.
func WithToken(ctx context.Context, token string) context.Context {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}
	md[AuthorizationKey] = bearerPrefix + token
	return metadata.NewContext(ctx, md)
}

// Authorizer provides convenience method for extracting username from request context metadata.
// When it returns error it is micro framework compatible and have Unauthorized status.
type Authorizer interface {
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
