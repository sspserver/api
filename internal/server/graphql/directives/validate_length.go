package directives

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func ValidateLength(ctx context.Context, obj any, next graphql.Resolver, min int, max int, trim bool) (res any, err error) {
	if res, err = next(ctx); err != nil {
		return nil, err
	}
	str := res.(string)
	if trim {
		str = strings.TrimSpace(str)
	}
	if len(str) < min {
		return nil, fmt.Errorf("value is too short, minimum length is %d", min)
	}
	if max > min && len(str) > max {
		return nil, fmt.Errorf("value is too long, maximum length is %d", max)
	}
	return str, nil
}
