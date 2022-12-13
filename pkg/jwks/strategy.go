// Package jwks provides a fosite.IssuerJWKSURIStrategy implementation.
package jwks

import (
	"context"
	"fmt"

	"go.infratographer.com/dmv/pkg/fositex"
)

var (
	// ErrUnknownIssuer is returned when the issuer is unknown.
	ErrUnknownIssuer = fmt.Errorf("unknown JWT issuer")
)

type issuerJWKSURIStrategy struct {
	issuerURIs map[string]string
}

// NewIssuerJWKSURIStrategy creates a new fosite.IssuerJWKSURIStrategy.
func NewIssuerJWKSURIStrategy(issuers []fositex.Issuer) fositex.IssuerJWKSURIStrategy {
	issuerURIs := make(map[string]string)
	for _, iss := range issuers {
		issuerURIs[iss.Name] = iss.JWKSURI
	}

	out := issuerJWKSURIStrategy{
		issuerURIs: issuerURIs,
	}

	return out
}

func (s issuerJWKSURIStrategy) GetIssuerJWKSURI(ctx context.Context, iss string) (string, error) {
	jwksURI, ok := s.issuerURIs[iss]
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrUnknownIssuer, iss)
	}

	return jwksURI, nil
}
