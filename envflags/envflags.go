package envflags

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var (
	EnvPrefix = ""
)

type ValueType interface {
	// string | int | time.Duration | bool | []string
	any
}

type Value[T ValueType] struct {
	flagValue  T
	converter  func(string) (T, error)
	isBoolFlag bool
}

func NewEnvFlagValue[T ValueType](
	envName string,
	defaultValue T,
	converter func(string) (T, error),
) *Value[T] {
	envFlag := &Value[T]{
		converter: converter,
	}
	envFlag.setDefault(envName, defaultValue)
	return envFlag
}

func (p Value[T]) IsBoolFlag() bool {
	return p.isBoolFlag
}

func (p *Value[T]) setDefault(envName string, defaultValue T) {
	if env := os.Getenv(EnvPrefix + envName); env != "" {
		converted, err := p.converter(env)
		if err == nil {
			p.flagValue = converted
			return
		} else {
			slog.Warn("error converting environment variable, ignoring", "env", EnvPrefix+envName, "error", err)
		}
	}
	p.flagValue = defaultValue
}

func (p Value[T]) String() string {
	return fmt.Sprintf("%v", p.flagValue)
}

func (p Value[T]) Get() T {
	return p.flagValue
}

func (p *Value[T]) Set(value string) error {
	if p.converter == nil {
		return fmt.Errorf("no converter for type %T", p.flagValue)
	}
	converted, err := p.converter(value)
	if err != nil {
		return err
	}
	p.flagValue = converted
	return nil
}

func NewString(env, defaultValue string) *Value[string] {
	converter := func(s string) (string, error) {
		return s, nil
	}
	pflag := NewEnvFlagValue(env, defaultValue, converter)
	return pflag
}

func NewBool(env string, defaultValue bool) *Value[bool] {
	converter := func(s string) (bool, error) {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return false, err
		}
		return b, nil
	}
	eflag := NewEnvFlagValue(env, defaultValue, converter)
	eflag.isBoolFlag = true
	return eflag
}

func NewInt(env string, defaultValue int) *Value[int] {
	converter := func(s string) (int, error) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	pflag := NewEnvFlagValue(env, defaultValue, converter)
	return pflag
}

func NewLogLevel(env string, defaultValue slog.Level) *Value[slog.Level] {
	converter := func(s string) (slog.Level, error) {
		switch strings.ToLower(s) {
		case "debug":
			return slog.LevelDebug, nil
		case "info":
			return slog.LevelInfo, nil
		case "warn":
			return slog.LevelWarn, nil
		case "error":
			return slog.LevelError, nil
		default:
			return slog.LevelInfo, fmt.Errorf("invalid log level: %s", s)
		}
	}
	pflag := NewEnvFlagValue(env, defaultValue, converter)
	return pflag
}