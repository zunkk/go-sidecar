package repo

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"

	glog "github.com/zunkk/go-sidecar/log"
)

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Duration(d).String() + "\""), nil
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	// remove "
	str := string(b)
	str = strings.TrimPrefix(str, "\"")
	str = strings.TrimSuffix(str, "\"")
	x, err := time.ParseDuration(str)
	if err != nil {
		return err
	}
	*d = Duration(x)
	return nil
}

func (d *Duration) MarshalText() (text []byte, err error) {
	return []byte(time.Duration(*d).String()), nil
}

func (d *Duration) UnmarshalText(b []byte) error {
	x, err := time.ParseDuration(string(b))
	if err != nil {
		return err
	}
	*d = Duration(x)
	return nil
}

func (d *Duration) ToDuration() time.Duration {
	return time.Duration(*d)
}

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d Duration) FormatToMinutes() string {
	totalMinutes := int64(time.Duration(d).Minutes())
	days := totalMinutes / (60 * 24)
	hours := (totalMinutes % (60 * 24)) / 60
	minutes := totalMinutes % 60

	var result string
	if days > 0 {
		result += fmt.Sprintf("%dd", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%dh", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%dm", minutes)
	}

	if result == "" {
		return "0m"
	}
	return result
}

func StringToTimeDurationHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(Duration(5)) {
			return data, nil
		}

		d, err := time.ParseDuration(data.(string))
		if err != nil {
			return nil, err
		}
		return Duration(d), nil
	}
}

func StringToLevelHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(glog.Level(0)) {
			return data, nil
		}

		var l glog.Level
		if err := l.UnmarshalText([]byte(data.(string))); err != nil {
			return nil, err
		}
		return l, nil
	}
}

type HTTP struct {
	Enable                bool     `mapstructure:"enable" toml:"enable"`
	Port                  int      `mapstructure:"port" toml:"port"`
	MultipartMemory       int64    `mapstructure:"-" toml:"multipart_memory"`
	ReadTimeout           Duration `mapstructure:"read_timeout" toml:"read_timeout"`
	WriteTimeout          Duration `mapstructure:"write_timeout" toml:"write_timeout"`
	TLSEnable             bool     `mapstructure:"tls_enable" toml:"tls_enable"`
	TLSCertFilePath       string   `mapstructure:"tls_cert_file_path" toml:"tls_cert_file_path"`
	TLSKeyFilePath        string   `mapstructure:"tls_key_file_path" toml:"tls_key_file_path"`
	JWTTokenValidDuration Duration `mapstructure:"jwt_token_valid_duration" toml:"jwt_token_valid_duration"`
	JWTTokenHMACKey       string   `mapstructure:"jwt_token_hmac_key" toml:"jwt_token_hmac_key"`
}

type DBInfo struct {
	Host     string `mapstructure:"host" toml:"host"`
	Port     uint32 `mapstructure:"port" toml:"port"`
	User     string `mapstructure:"user" toml:"user"`
	Password string `mapstructure:"password" toml:"password"`
	DBName   string `mapstructure:"db_name" toml:"db_name"`
	Schema   string `mapstructure:"schema" toml:"schema"`
	SSLMode  string `mapstructure:"ssl_mode" toml:"ssl_mode"`
}

type Mongodb struct {
	DBInfo          `mapstructure:",squash"`
	ConnectTimeout  Duration `mapstructure:"connect_timeout" toml:"connect_timeout"`
	MaxPoolSize     int      `mapstructure:"max_pool_size" toml:"max_pool_size"`
	MaxConnIdleTime Duration `mapstructure:"max_conn_idle_time" toml:"max_conn_idle_time"`
}

type Log struct {
	Level            glog.Level            `mapstructure:"level" toml:"level"`
	Filename         string                `mapstructure:"file_name" toml:"file_name"`
	MaxAge           Duration              `mapstructure:"max_age" toml:"max_age"`
	MaxSizeStr       string                `mapstructure:"max_size" toml:"max_size"`
	MaxSize          int64                 `mapstructure:"-" toml:"-"`
	RotationTime     Duration              `mapstructure:"rotation_time" toml:"rotation_time"`
	EnableColor      bool                  `mapstructure:"enable_color" toml:"enable_color"`
	EnableCaller     bool                  `mapstructure:"enable_caller" toml:"enable_caller"`
	DisableTimestamp bool                  `mapstructure:"disable_timestamp" toml:"disable_timestamp"`
	ModuleLevelMap   map[string]glog.Level `mapstructure:"module_level_map" toml:"module_level_map"`
}
