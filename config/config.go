package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	Version   = "dev"
	Commit    = "dev"
	BuildDate = "dev"
)

type Config struct {
	Name                   string              `mapstructure:"name" yaml:"name"`
	Telegram               TelegramConfig      `mapstructure:"telegram" yaml:"telegram"`
	Pushover               PushoverConfig      `mapstructure:"pushover" yaml:"pushover"`
	Database               DatabaseConfig      `mapstructure:"database" yaml:"database"`
	Jwt                    JwtConfig           `mapstructure:"jwt" yaml:"jwt"`
	Server                 ServerConfig        `mapstructure:"server" yaml:"server"`
	SchedulerJobs          SchedulerConfig     `mapstructure:"scheduler_jobs" yaml:"scheduler_jobs"`
	EmailConfig            EmailConfig         `mapstructure:"email" yaml:"email"`
	StripeConfig           StripeConfig        `mapstructure:"stripe" yaml:"stripe"`
	OAuth2Config           OAuth2Config        `mapstructure:"oauth2" yaml:"oauth2"`
	WebhookConfig          WebhookConfig       `mapstructure:"webhook" yaml:"webhook"`
	MFAConfig              MFAConfig           `mapstructure:"mfa" yaml:"mfa"`
	IsDoneTickDotCom       bool                `mapstructure:"is_done_tick_dot_com" yaml:"is_done_tick_dot_com"`
	IsUserCreationDisabled bool                `mapstructure:"is_user_creation_disabled" yaml:"is_user_creation_disabled"`
	MinVersion             string              `mapstructure:"min_version" yaml:"min_version"`
	DonetickCloudConfig    DonetickCloudConfig `mapstructure:"donetick_cloud" yaml:"donetick_cloud"`
	Storage                StorageConfig       `mapstructure:"storage" yaml:"storage"`
	Info                   Info
}

type Info struct {
	Version   string
	Commit    string
	BuildDate string
}
type StorageConfig struct {
	StorageType string `mapstructure:"storage_type" yaml:"storage_type"`
	// CloudStorage:
	BucketName     string `mapstructure:"bucket_name" yaml:"bucket_name"`
	Region         string `mapstructure:"region" yaml:"region"`
	BasePath       string `mapstructure:"base_path" yaml:"base_path"`
	AccessKey      string `mapstructure:"access_key" yaml:"access_key"`
	SecretKey      string `mapstructure:"secret_key" yaml:"secret_key"`
	Endpoint       string `mapstructure:"endpoint" yaml:"endpoint"`
	MaxUserStorage int    `mapstructure:"max_user_storage" yaml:"max_user_storage"`
	MaxFileSize    int64  `mapstructure:"max_file_size" yaml:"max_file_size"`
	PublicHost     string `mapstructure:"public_host" yaml:"public_host"`
}
type DonetickCloudConfig struct {
	GoogleClientID        string `mapstructure:"google_client_id" yaml:"google_client_id"`
	GoogleAndroidClientID string `mapstructure:"google_android_client_id" yaml:"google_android_client_id"`
	GoogleIOSClientID     string `mapstructure:"google_ios_client_id" yaml:"google_ios_client_id"`
}

type TelegramConfig struct {
	Token string `mapstructure:"token" yaml:"token"`
}

type PushoverConfig struct {
	Token string `mapstructure:"token" yaml:"token"`
}

type DatabaseConfig struct {
	Type      string `mapstructure:"type" yaml:"type"`
	Host      string `mapstructure:"host" yaml:"host"`
	Port      int    `mapstructure:"port" yaml:"port"`
	User      string `mapstructure:"user" yaml:"user"`
	Password  string `mapstructure:"password" yaml:"password"`
	Name      string `mapstructure:"name" yaml:"name"`
	Migration bool   `mapstructure:"migration" yaml:"migration" default:"true"`
	LogLevel  int    `mapstructure:"logger" yaml:"logger"`
}

type JwtConfig struct {
	Secret      string        `mapstructure:"secret" yaml:"secret"`
	SessionTime time.Duration `mapstructure:"session_time" yaml:"session_time"`
	MaxRefresh  time.Duration `mapstructure:"max_refresh" yaml:"max_refresh"`
}

type ServerConfig struct {
	Port             int           `mapstructure:"port" yaml:"port"`
	RatePeriod       time.Duration `mapstructure:"rate_period" yaml:"rate_period"`
	RateLimit        int           `mapstructure:"rate_limit" yaml:"rate_limit"`
	ReadTimeout      time.Duration `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout     time.Duration `mapstructure:"write_timeout" yaml:"write_timeout"`
	CorsAllowOrigins []string      `mapstructure:"cors_allow_origins" yaml:"cors_allow_origins"`
	ServeFrontend    bool          `mapstructure:"serve_frontend" yaml:"serve_frontend"`
}

type SchedulerConfig struct {
	DueJob     time.Duration `mapstructure:"due_job" yaml:"due_job"`
	OverdueJob time.Duration `mapstructure:"overdue_job" yaml:"overdue_job"`
	PreDueJob  time.Duration `mapstructure:"pre_due_job" yaml:"pre_due_job"`
}

type StripeConfig struct {
	APIKey         string         `mapstructure:"api_key"`
	WhitelistedIPs []string       `mapstructure:"whitelisted_ips"`
	Prices         []StripePrices `mapstructure:"prices"`
	SuccessURL     string         `mapstructure:"success_url"`
	CancelURL      string         `mapstructure:"cancel_url"`
}

type StripePrices struct {
	PriceID string `mapstructure:"id"`
	Name    string `mapstructure:"name"`
}

type EmailConfig struct {
	Email   string `mapstructure:"email"`
	Key     string `mapstructure:"key"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	AppHost string `mapstructure:"appHost"`
}

type OAuth2Config struct {
	ClientID     string   `mapstructure:"client_id" yaml:"client_id"`
	ClientSecret string   `mapstructure:"client_secret" yaml:"client_secret"`
	RedirectURL  string   `mapstructure:"redirect_url" yaml:"redirect_url"`
	Scopes       []string `mapstructure:"scopes" yaml:"scopes"`
	AuthURL      string   `mapstructure:"auth_url" yaml:"auth_url"`
	TokenURL     string   `mapstructure:"token_url" yaml:"token_url"`
	UserInfoURL  string   `mapstructure:"user_info_url" yaml:"user_info_url"`
	Name         string   `mapstructure:"name" yaml:"name"`
}

type WebhookConfig struct {
	Timeout   time.Duration `mapstructure:"timeout" yaml:"timeout" default:"5s"`
	QueueSize int           `mapstructure:"queue_size" yaml:"queue_size" default:"100"`
}

type MFAConfig struct {
	Enabled                 bool          `mapstructure:"enabled" yaml:"enabled" default:"true"`
	SessionTimeoutMinutes   int           `mapstructure:"session_timeout_minutes" yaml:"session_timeout_minutes" default:"15"`
	BackupCodeCount         int           `mapstructure:"backup_code_count" yaml:"backup_code_count" default:"10"`
	MaxVerificationAttempts int           `mapstructure:"max_verification_attempts" yaml:"max_verification_attempts" default:"5"`
	RateLimitWindow         time.Duration `mapstructure:"rate_limit_window" yaml:"rate_limit_window" default:"5m"`
}

func NewConfig() *Config {
	return &Config{
		Telegram: TelegramConfig{
			Token: "",
		},
		Database: DatabaseConfig{
			Type:      "sqlite",
			Migration: true,
		},
		Jwt: JwtConfig{
			Secret:      "secret",
			SessionTime: 7 * 24 * time.Hour,
			MaxRefresh:  7 * 24 * time.Hour,
		},
	}
}
func configEnvironmentOverrides(Config *Config) {
	if os.Getenv("DONETICK_TELEGRAM_TOKEN") != "" {
		Config.Telegram.Token = os.Getenv("DONETICK_TELEGRAM_TOKEN")
	}
	if os.Getenv("DONETICK_PUSHOVER_TOKEN") != "" {
		Config.Pushover.Token = os.Getenv("DONETICK_PUSHOVER_TOKEN")
	}
	if os.Getenv("DONETICK_DISABLE_SIGNUP") == "true" {
		Config.IsUserCreationDisabled = true
	}

}
func LoadConfig() *Config {
	// set the config name based on the environment:

	if os.Getenv("DT_ENV") == "local" {
		viper.SetConfigName("local")
	} else if os.Getenv("DT_ENV") == "prod" {
		viper.SetConfigName("prod")
	} else if os.Getenv("DT_ENV") == "selfhosted" {
		viper.SetConfigName("selfhosted")
	} else {
		viper.SetConfigName("local")
	}
	// get logger and log the current environment:
	fmt.Printf("--ConfigLoad config for environment: %s ", os.Getenv("DT_ENV"))
	viper.SetEnvPrefix("DT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	// print a useful error:
	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--ConfigLoad name : %s ", config.Name)

	configEnvironmentOverrides(&config)
	config.Info.Version = Version
	config.Info.Commit = Commit
	config.Info.BuildDate = BuildDate
	return &config

	// return LocalConfig()
}
