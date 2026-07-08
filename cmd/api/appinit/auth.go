package appinit

import (
	"context"
	"strings"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"gorm.io/gorm"

	"github.com/geniusrabbit/blaze-api/pkg/auth/jwt"
	"github.com/geniusrabbit/blaze-api/pkg/auth/oauth2/serverprovider"
	"github.com/geniusrabbit/blaze-api/pkg/cache"
	"github.com/geniusrabbit/blaze-api/pkg/cache/dummy"
	"github.com/geniusrabbit/blaze-api/pkg/cache/memory"
	"github.com/geniusrabbit/blaze-api/pkg/cache/redis"
	"github.com/geniusrabbit/blaze-api/repository/user"
	user_repository "github.com/geniusrabbit/blaze-api/repository/user/repository"

	"github.com/sspserver/api/cmd/api/appcontext"
	"github.com/sspserver/api/pkg/models"
)

type oauth2UserAccessor struct {
	emailRepo    user.EmailRepository[*models.User]
	passwordRepo user.PasswordRepository[*models.User]
}

func (a *oauth2UserAccessor) GetByEmail(ctx context.Context, email string) (user.Model, error) {
	return a.emailRepo.GetByEmail(ctx, email)
}

func (a *oauth2UserAccessor) GetByPassword(ctx context.Context, userID uint64, password string) (user.Model, error) {
	return a.passwordRepo.GetByPassword(ctx, userID, password)
}

// Auth new provider
func Auth(ctx context.Context, conf *appcontext.ConfigType, masterDatabase *gorm.DB) (fosite.OAuth2Provider, *jwt.Provider) {
	oauth2config := &fosite.Config{
		AccessTokenLifespan:           conf.OAuth2.AccessTokenLifespan,
		RefreshTokenLifespan:          conf.OAuth2.RefreshTokenLifespan,
		AuthorizeCodeLifespan:         conf.OAuth2.AuthorizeCodeLifespan,
		HashCost:                      conf.OAuth2.HashCost,
		DisableRefreshTokenValidation: conf.OAuth2.DisableRefreshTokenValidation,
		SendDebugMessagesToClients:    conf.OAuth2.SendDebugMessagesToClients,
	}
	sessionCache := newCache(ctx, conf.OAuth2.CacheConnect, conf.OAuth2.CacheLifetime)
	newUser := func() *models.User { return &models.User{} }
	userRepository := user_repository.NewRepository(newUser)
	userAccessor := &oauth2UserAccessor{
		emailRepo:    user_repository.NewEmailRepository(userRepository, newUser),
		passwordRepo: user_repository.NewPasswordRepository(userRepository, newUser),
	}
	oauth2storage := serverprovider.NewDatabaseStorage(
		masterDatabase,
		userAccessor,
		sessionCache,
		conf.OAuth2.CacheLifetime,
	)
	oauth2provider := serverprovider.NewProvider(
		oauth2config,
		oauth2storage,
		&compose.CommonStrategy{
			CoreStrategy: compose.NewOAuth2HMACStrategy(oauth2config),
		},
		nil,
	)
	jwtProvider := jwt.NewDefaultProvider(
		conf.OAuth2.Secret,
		conf.OAuth2.AccessTokenLifespan,
		conf.IsDebug(),
	)
	return oauth2provider, jwtProvider
}

func newCache(ctx context.Context, connect string, lifetime time.Duration) cache.Client {
	switch {
	case connect == ":memory:":
		cacheObj, err := memory.NewTimeout(ctx, lifetime)
		fatalError(err, "memory cache")
		return cacheObj
	case connect == ":dummy:" || connect == "":
		return dummy.New()
	case strings.HasPrefix(connect, "redis://"):
		cli, err := redis.NewByURL(connect)
		fatalError(err, "redis cache")
		return cli
	default:
		return dummy.New()
	}
}
