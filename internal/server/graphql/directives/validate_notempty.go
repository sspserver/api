package directives

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/demdxx/gocast/v2"
)

func ValidateNotEmpty(ctx context.Context, obj any, next graphql.Resolver, trim bool) (res any, err error) {
	if res, err = next(ctx); err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("value is nil")
	}
	if trim {
		res = strings.TrimSpace(res.(string))
	}
	if gocast.IsEmpty(res) {
		return nil, errors.New("value is empty")
	}
	return res, nil
}
