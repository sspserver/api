package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/geniusrabbit/blaze-api/pkg/auth"
	"github.com/geniusrabbit/blaze-api/pkg/context/ctxlogger"
	"github.com/geniusrabbit/blaze-api/pkg/middleware"
	"github.com/geniusrabbit/blaze-api/pkg/profiler"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"
	"github.com/go-chi/chi/v5"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type (
	contextWrapper func(context.Context) context.Context
	muxInitWrapper func(mux *chi.Mux)
)

// HTTPServer wrapper object
type HTTPServer struct {
	RequestTimeout time.Duration
	ContextWrap    contextWrapper
	InitWrap       muxInitWrapper
	Authorizers    []auth.Authorizer
	SessionManager *scs.SessionManager
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func (s *HTTPServer) Run(ctx context.Context, address string) (err error) {
	ctxlogger.Get(ctx).Info("Start server HTTP API: " + address)

	mux := chi.NewRouter()
	mux.Handle("/healthcheck", http.HandlerFunc(profiler.HealthCheckHandler))
	mux.Handle("/metrics", promhttp.Handler())

	if s.InitWrap != nil {
		s.InitWrap(mux)
	}

	h := http.Handler(mux)

	// Add middleware's
	h = auth.Middelware(h, s.Authorizers...)
	h = middleware.HTTPContextWrapper(h, s.ContextWrap)
	h = middleware.HTTPSession(h, s.SessionManager)
	h = middleware.RealIP(h)
	h = middleware.AllowCORS(h)
	h = middleware.RequestID(h, requestid.WithCloudflareRequestID())
	h = nethttp.Middleware(opentracing.GlobalTracer(), h)

	srv := &http.Server{
		Addr:         address,
		Handler:      h,
		ReadTimeout:  s.RequestTimeout,
		WriteTimeout: s.RequestTimeout,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
	}
	go func() {
		<-ctx.Done()
		ctxlogger.Get(ctx).Info("Shutting down the http server")
		if err := srv.Shutdown(context.Background()); err != nil {
			ctxlogger.Get(ctx).Error("Failed to shutdown http server", zap.Error(err))
		}
	}()

	ctxlogger.Get(ctx).Info(fmt.Sprintf("Starting listening at %s", address))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		ctxlogger.Get(ctx).Error("Failed to listen and serve", zap.Error(err))
		return err
	}
	return nil
}
