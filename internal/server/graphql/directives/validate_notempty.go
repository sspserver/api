package directives

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/demdxx/gocast/v2"
)

func ValidateNotEmpty(ctx context.Context, obj any, next graphql.Resolver, trim, ornil bool) (res any, err error) {
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
		return nil, errors.New("value is nil")
	case string:
		str = v
	case *string:
		str = *v
		canBeNil = true
	default:
		if gocast.IsEmpty(res) {
			if ornil && reflect.ValueOf(res).Kind() == reflect.Ptr {
				return nil, nil
			}
			return nil, errors.New("value is empty")
		}
	}

	// Trim the string if needed
	if trim {
		str = strings.TrimSpace(str)
	}

	// Check if the string is empty
	if str == "" && ornil && canBeNil {
		return nil, nil
	}

	// Return the error if the string is empty
	if str == "" {
		return nil, errors.New("value is empty")
	}

	if canBeNil {
		return &str, nil
	}
	return str, nil
}
