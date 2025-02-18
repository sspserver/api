package directives

import (
	"context"
	"fmt"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
)

func ValidateRegex(ctx context.Context, obj any, next graphql.Resolver, pattern string, strict bool) (res any, err error) {
	if res, err = next(ctx); err != nil {
		return nil, err
	}
	if !strict && res == nil {
		return nil, nil
	}
	var str string
	switch v := res.(type) {
	case string:
		str = v
	case *string:
		str = *v
	default:
		return nil, fmt.Errorf("value is not a string")
	}
	if !regexp.MustCompile(pattern).MatchString(str) {
		return nil, fmt.Errorf("value does not match pattern `%s`", pattern)
	}
	return str, nil
}
