package models

import (
	"github.com/demdxx/xtypes"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ErrorInvalidField(field, msg string, path ...string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: field + ": " + msg,
		Path: ast.Path(
			xtypes.SliceApply(append(path, field),
				func(it string) ast.PathElement { return ast.PathName(it) }),
		),
		Extensions: map[string]any{
			"field": field,
			"code":  "validation",
		},
	}
}

func ErrorRequiredField(field string, path ...string) *gqlerror.Error {
	return ErrorInvalidField(field, "is required", path...)
}
