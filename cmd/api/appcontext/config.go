// Package appcontext provides config options
package appcontext

/**
 ██████╗ ██████╗ ███╗   ██╗███████╗██╗ ██████╗
██╔════╝██╔═══██╗████╗  ██║██╔════╝██║██╔════╝
██║     ██║   ██║██╔██╗ ██║█████╗  ██║██║  ███╗
██║     ██║   ██║██║╚██╗██║██╔══╝  ██║██║   ██║
╚██████╗╚██████╔╝██║ ╚████║██║     ██║╚██████╔╝
 ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝     ╚═╝ ╚═════╝
*/

import (
	"encoding/json"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
)

type serverConfig struct {
	HTTP struct {
		Listen       string        `default:":8080" field:"listen" json:"listen" yaml:"listen" cli:"http-listen" env:"SERVER_HTTP_LISTEN"`
		ReadTimeout  time.Duration `default:"120s" field:"read_timeout" json:"read_timeout" yaml:"read_timeout" env:"SERVER_HTTP_READ_TIMEOUT"`
		WriteTimeout time.Duration `default:"120s" field:"write_timeout" json:"write_timeout" yaml:"write_timeout" env:"SERVER_HTTP_WRITE_TIMEOUT"`
	}
	Profile struct {
		Mode   string `json:"mode" yaml:"mode" default:"net" env:"SERVER_PROFILE_MODE"`
		Listen string `json:"listen" yaml:"listen" default:"" env:"SERVER_PROFILE_LISTEN"`
	} `json:"profile" yaml:"profile"`
}

type sessionConfig struct {
	CookieName string        `json:"cookie_name" yaml:"cookie_name" default:"sessid" env:"SESSION_COOKIE_NAME"`
	Lifetime   time.Duration `json:"lifetime" yaml:"lifetime" default:"1h" env:"SESSION_LIFETIME"`
	// DevToken is the permanent token which can be used to API access in develop mode
	DevToken     string `json:"dev_token" yaml:"dev_token" env:"SESSION_DEV_TOKEN"`
	DevUserID    uint64 `json:"dev_user_id" yaml:"dev_user_id" env:"SESSION_DEV_USER_ID"`
	DevAccountID uint64 `json:"dev_account_id" yaml:"dev_account_id" env:"SESSION_DEV_ACCOUNT_ID"`
}

type storageConfig struct {
	MasterConnect string `json:"master_connect" yaml:"master_connect" env:"SYSTEM_STORAGE_DATABASE_MASTER_CONNECT"`
	SlaveConnect  string `json:"slave_connect" yaml:"slave_connect" env:"SYSTEM_STORAGE_DATABASE_SLAVE_CONNECT"`
}

// StatisticConfig contains statistic configuration options
type statisticConfig struct {
	Connect string `json:"connect" yaml:"connect" env:"SYSTEM_STATISTIC_CONNECT"`
}

type socialAuthProviderEndpoint struct {
	AuthURL       string `json:"auth_url" yaml:"auth_url" env:"AUTH_URL"`
	DeviceAuthURL string `json:"device_auth_url" yaml:"device_auth_url" env:"DEVICE_AUTH_URL"`
	TokenURL      string `json:"token_url" yaml:"token_url" env:"TOKEN_URL"`

	// AuthStyle optionally specifies how the endpoint wants the
	// client ID & client secret sent. The zero value means to
	// auto-detect.
	AuthStyle oauth2.AuthStyle `json:"auth_style" yaml:"auth_style" env:"AUTH_STYLE"`
}

func (en *socialAuthProviderEndpoint) IsEmpty() bool {
	return en.AuthURL == "" && en.DeviceAuthURL == "" && en.TokenURL == ""
}

func (en *socialAuthProviderEndpoint) OAuth2(provider string) oauth2.Endpoint {
	if en.IsEmpty() {
		switch strings.ToLower(provider) {
		case "google":
			return google.Endpoint
		case "facebook":
			return facebook.Endpoint
		case "linkedin":
			return linkedin.Endpoint
		}
	}
	return oauth2.Endpoint{
		AuthURL:       en.AuthURL,
		DeviceAuthURL: en.DeviceAuthURL,
		TokenURL:      en.TokenURL,
		AuthStyle:     en.AuthStyle,
	}
}

type socialAuthProviderConfig struct {
	// ClientID is the application's ID.
	ClientID string `json:"client_id" yaml:"client_id" env:"CLIENT_ID"`

	// ClientSecret is the application's secret.
	ClientSecret string `json:"client_secret" yaml:"client_secret" env:"CLIENT_SECRET"`

	// Endpoint contains the resource server's token endpoint
	// URLs. These are constants specific to each server and are
	// often available via site-specific packages, such as
	// google.Endpoint or github.Endpoint.
	Endpoint socialAuthProviderEndpoint `json:"endpoint" yaml:"endpoint" envPrefix:"ENDPOINT_"`

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string `json:"redirect_url" yaml:"redirect_url" env:"REDIRECT_URL"`

	// Scope specifies optional requested permissions.
	Scopes []string `json:"scopes" yaml:"scopes" env:"SCOPES"`
}

func (s *socialAuthProviderConfig) OAuth2Config(provider string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     s.ClientID,
		ClientSecret: s.ClientSecret,
		Endpoint:     s.Endpoint.OAuth2(provider),
		RedirectURL:  s.RedirectURL,
		Scopes:       s.Scopes,
	}
}

func (s *socialAuthProviderConfig) IsValid() bool {
	return s != nil && s.ClientID != "" && s.ClientSecret != "" && s.RedirectURL != ""
}

type socialAuthConfig struct {
	Google   socialAuthProviderConfig `json:"google" yaml:"google" envPrefix:"GOOGLE_"`
	Facebook socialAuthProviderConfig `json:"facebook" yaml:"facebook" envPrefix:"FACEBOOK_"`
	XCOM     socialAuthProviderConfig `json:"xcom" yaml:"xcom" envPrefix:"XCOM_"` // Ex Twitter
	LinkedIn socialAuthProviderConfig `json:"linkedin" yaml:"linkedin" envPrefix:"LINKEDIN_"`
}

type oauth2Config struct {
	// Secret used by server to preprocess the secrets. Minimal size is 32 symbols
	Secret string `json:"secret" yaml:"secret" env:"OAUTH2_SECRET"`

	// AccessTokenLifespan sets how long an access token is going to be valid. Defaults to one hour.
	AccessTokenLifespan time.Duration `json:"access_token_lifespan" yaml:"access_token_lifespan" env:"OAUTH2_ACCESS_TOKEN_LIFESPAN" default:"1h"`

	// RefreshTokenLifespan sets how long a refresh token is going to be valid. Defaults to 30 days. Set to -1 for
	// refresh tokens that never expire.
	RefreshTokenLifespan time.Duration `json:"refresh_token_lifespan" yaml:"refresh_token_lifespan" env:"OAUTH2_REFRESH_TOKEN_LIFESPAN" default:"720h"`

	// AuthorizeCodeLifespan sets how long an authorize code is going to be valid. Defaults to fifteen minutes.
	AuthorizeCodeLifespan time.Duration `json:"authorize_code_lifespan" yaml:"authorize_code_lifespan" env:"OAUTH2_AUTHORIZE_CODE_LIFESPAN" default:"15m"`

	// HashCost sets the cost of the password hashing cost. Defaults to 12.
	HashCost int `json:"hash_cost" yaml:"hash_cost" env:"OAUTH2_HASH_COST"`

	// DisableRefreshTokenValidation sets the introspection endpoint to disable refresh token validation.
	DisableRefreshTokenValidation bool `json:"disable_refresh_token_validation" yaml:"disable_refresh_token_validation" env:"OAUTH2_DISABLE_REFRESH_TOKEN_VALIDATION"`

	// SendDebugMessagesToClients if set to true, includes error debug messages in response payloads. Be aware that sensitive
	// data may be exposed, depending on your implementation of Fosite. Such sensitive data might include database error
	// codes or other information. Proceed with caution!
	SendDebugMessagesToClients bool `json:"send_debug_messages_to_clients" yaml:"send_debug_messages_to_clients" env:"OAUTH2_SEND_DEBUG_MESSAGES_TO_CLIENTS"`

	// CacheConnect provides functionality of session cache to reduce amount of requests to the database
	// Supports: redis://host:port/dbNum, :memory:, :dummy:
	CacheConnect string `json:"cache_connect" yaml:"cache_connect" env:"OAUTH2_CACHE_CONNECT"`

	// CacheLifetime define the lifetime of elements in the cache
	CacheLifetime time.Duration `json:"cache_lifetime" yaml:"cache_lifetime" env:"OAUTH2_CACHE_LIFETIME"`
}

type permissionConfig struct {
	RoleCacheLifetime time.Duration `json:"role_cache_lifetime" yaml:"role_cache_lifetime" env:"PERMISSIONS_CACHE_LIFETIME" default:"10s"`
}

type systemConfig struct {
	Storage   storageConfig   `json:"storage" yaml:"storage"`
	Statistic statisticConfig `json:"statistic" yaml:"statistic"`
}

// MessangerConfig contains email configuration options for messanger
type messangerConfig struct {
	Email struct {
		Mailer      string `json:"mailer" yaml:"mailer" env:"MESSANGER_EMAIL_MAILER" default:"smtp"`
		URL         string `json:"url" yaml:"url" env:"MESSANGER_EMAIL_URL"`
		Port        int    `json:"port" yaml:"port" env:"MESSANGER_EMAIL_PORT" default:"587"`
		APIKey      string `json:"api_key" yaml:"api_key" env:"MESSANGER_EMAIL_API_KEY"`
		Domain      string `json:"domain" yaml:"domain" env:"MESSANGER_EMAIL_DOMAIN"`
		FromAddress string `json:"from_address" yaml:"from_address" env:"MESSANGER_EMAIL_FROM_ADDRESS"`
		FromName    string `json:"from_name" yaml:"from_name" env:"MESSANGER_EMAIL_FROM_NAME"`
		Password    string `json:"password" yaml:"password" env:"MESSANGER_EMAIL_PASSWORD"`
	} `json:"email" yaml:"email"`

	EmailDefaults struct {
		Name         string `json:"name" yaml:"name" env:"MESSANGER_EMAIL_DEFAULT_VAR_NAME"`
		Position     string `json:"position" yaml:"position" env:"MESSANGER_EMAIL_DEFAULT_VAR_POSITION"`
		Contact      string `json:"contact" yaml:"contact" env:"MESSANGER_EMAIL_DEFAULT_VAR_CONTACT"`
		SupportEmail string `json:"support_email" yaml:"support_email" env:"MESSANGER_EMAIL_DEFAULT_VAR_SUPPORT_EMAIL"`
	} `json:"email_defaults" yaml:"email_defaults"`
}

type optionsConfig struct {
	RTBServerDomain      string `json:"rtb_server_domain" yaml:"rtb_server_domain" env:"OPTION_RTB_SERVER_DOMAIN"`
	AdTemplateCode       string `json:"ad_template_code" yaml:"ad_template_code" env:"OPTION_AD_TEMPLATE_CODE"`
	AdDirectTemplateURL  string `json:"ad_direct_template_url" yaml:"ad_direct_template_url" env:"OPTION_AD_DIRECT_TEMPLATE_URL"`
	AdDirectTemplateCode string `json:"ad_direct_template_code" yaml:"ad_direct_template_code" env:"OPTION_AD_DIRECT_TEMPLATE_CODE"`
}

// ConfigType contains all application options
type ConfigType struct {
	ServiceName    string `json:"service_name" yaml:"service_name" env:"SERVICE_NAME" default:"adnet.api"`
	DatacenterName string `json:"datacenter_name" yaml:"datacenter_name" env:"DC_NAME" default:"??"`
	Hostname       string `json:"hostname" yaml:"hostname" env:"HOSTNAME"`
	Hostcode       string `json:"hostcode" yaml:"hostcode" env:"HOSTCODE"`

	LogAddr    string `json:"log_addr" default:"" env:"LOG_ADDR"`
	LogLevel   string `json:"log_level" default:"debug" env:"LOG_LEVEL"`
	LogEncoder string `json:"log_encoder" env:"LOG_ENCODER"`

	Server      serverConfig     `json:"server" yaml:"server"`
	Session     sessionConfig    `json:"session" yaml:"session"`
	System      systemConfig     `json:"system" yaml:"system"`
	SocialAuth  socialAuthConfig `json:"social_auth" yaml:"social_auth"`
	OAuth2      oauth2Config     `json:"oauth2" yaml:"oauth2"`
	Messanger   messangerConfig  `json:"messanger" yaml:"messanger"`
	Permissions permissionConfig `json:"permissions" yaml:"permissions"`
	Options     optionsConfig    `json:"options" yaml:"options"`
}

// String implementation of Stringer interface
func (cfg *ConfigType) String() (res string) {
	if data, err := json.MarshalIndent(cfg, "", "  "); err != nil {
		res = `{"error":"` + err.Error() + `"}`
	} else {
		res = string(data)
	}
	return res
}

// IsDebug mode
func (cfg *ConfigType) IsDebug() bool {
	return strings.EqualFold(cfg.LogLevel, "debug")
}

func (cfg *ConfigType) IsInfo() bool {
	return strings.EqualFold(cfg.LogLevel, "info")
}

// Config global value
var Config ConfigType
