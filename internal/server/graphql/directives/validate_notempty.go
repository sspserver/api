package directives

import (
	"context"
	"errors"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/demdxx/gocast/v2"
)

func ValidateNotEmpty(ctx context.Context, obj any, next graphql.Resolver, trim, ornil bool) (res any, err error) {
	if res, err = next(ctx); err != nil {
		return nil, err
	}

	// Check if the value is a string or a pointer to a string
	switch res.(type) {
	case nil:
		if ornil {
			return nil, nil
		}
		return nil, errors.New("value is nil")
	default:
		if gocast.IsEmpty(res) {
			kind := reflect.ValueOf(res).Kind()
			if ornil && (kind == reflect.Ptr ||
				kind == reflect.Slice ||
				kind == reflect.Map ||
				kind == reflect.Array ||
				kind == reflect.Chan ||
				kind == reflect.Func ||
				kind == reflect.Interface) {
				return nil, nil
			}
			return nil, errors.New("value is empty")
		}
	}

	return _validateLength(res, 1, 0, trim, ornil)
}
