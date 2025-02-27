package directives

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func ValidateLength(ctx context.Context, obj any, next graphql.Resolver, min int, max int, trim, ornil bool) (res any, err error) {
	if res, err = next(ctx); err != nil {
		return nil, err
	}

	var (
		str      string
		canBeNil = false
	)

	// Check if the value is a string or a pointer to a string
	switch v := res.(type) {
	case nil:
		if ornil {
			return nil, nil
		}
		return nil, fmt.Errorf("value is nil")
	case string:
		str = v
	case *string:
		str = *v
		canBeNil = true
	default:
		return nil, fmt.Errorf("value is not a string")
	}

	// Trim the string if needed
	if trim {
		str = strings.TrimSpace(str)
	}

	// Check if the string is empty
	if str == "" && ornil && canBeNil {
		return nil, nil
	}

	// Check the min length of the string
	if len(str) < min {
		return nil, fmt.Errorf("value is too short, minimum length is %d", min)
	}

	// Check the max length of the string
	if max > min && len(str) > max {
		return nil, fmt.Errorf("value is too long, maximum length is %d", max)
	}

	if canBeNil {
		return &str, nil
	}
	return str, nil
}
