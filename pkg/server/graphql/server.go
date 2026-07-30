package graphql

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/auth/jwt"
	blazeDirectives "github.com/geniusrabbit/blaze-api/server/graphql/directives"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/sspserver/api/pkg/context/ctxcache"
	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/agreement"
	"github.com/sspserver/api/pkg/server/graphql/directives"
	"github.com/sspserver/api/pkg/server/graphql/generated"
	"github.com/sspserver/api/pkg/server/graphql/resolvers"
)

// GraphQL mux handler
func GraphQL(usecases *resolvers.Usecases, provider *jwt.Provider) http.Handler {
	agreements := agreement.NewUsecase(
		agreement.NewRepositoryOptions(usecases.Options, resolvers.ListAgreements()),
	)
	srv := handler.New(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: resolvers.NewResolver(usecases, provider),
			Directives: generated.DirectiveRoot{
				Auth:              blazeDirectives.Auth,
				Acl:               blazeDirectives.HasPermissions[*models.User, *models.Account],
				HasPermissions:    blazeDirectives.HasPermissions[*models.User, *models.Account],
				SkipNoPermissions: blazeDirectives.SkipNoPermissions,
				RequireAgreements: directives.RequireAgreements(agreements),
				Length:            blazeDirectives.ValidateLength,
				Notempty:          blazeDirectives.ValidateNotEmpty,
				Regex:             blazeDirectives.ValidateRegex,
				Range:             blazeDirectives.ValidateRange,
			},
		}),
	)
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.UrlEncodedForm{})
	srv.AddTransport(transport.GRAPHQL{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	srv.SetRecoverFunc(recoverHandler)

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		span, ctx := opentracing.StartSpanFromContext(r.Context(), "graphql.request")
		defer span.Finish()
		ctx = ctxcache.WithCacheBlock(ctx)
		srv.ServeHTTP(rw, r.WithContext(ctx))
		ctxcache.ReleaseCacheBlock(ctx)
	})
}

func recoverHandler(ctx context.Context, err any) error {
	switch verr := err.(type) {
	case error:
		if errors.Is(verr, acl.ErrNoPermissions) {
			return verr
		}
	}
	return graphql.DefaultRecover(ctx, err)
}
