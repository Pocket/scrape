package envflags

import (
	"os"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue string
		setValue     string
		want         string
		expectError  bool
	}{
		{
			name:         "env, no explicit value",
			envValue:     "foo",
			defaultValue: "bar",
			setValue:     "_",
			want:         "foo",
			expectError:  false,
		},
		{
			name:         "env not set, no explicit value",
			envValue:     "",
			defaultValue: "bar",
			setValue:     "_",
			want:         "bar",
			expectError:  false,
		},
		{
			name:         "env not set, explicit value",
			envValue:     "",
			defaultValue: "bar",
			setValue:     "baz",
			want:         "baz",
			expectError:  false,
		},
		{
			name:         "env, explicit value",
			envValue:     "foo",
			defaultValue: "bar",
			setValue:     "baz",
			want:         "baz",
			expectError:  false,
		},
		{
			name:         "env, empty explicit value",
			envValue:     "foo",
			defaultValue: "bar",
			setValue:     "",
			want:         "",
			expectError:  false,
		},
	}
	envKey := "ENVFLAGS_TEST"
	for _, test := range tests {
		os.Setenv(envKey, test.envValue)
		pflag := NewString(envKey, test.defaultValue)
		if test.setValue != "_" {
			err := pflag.Set(test.setValue)
			if (err != nil) != test.expectError {
				t.Errorf(
					"%s: Set(%s) returned error %v, expected error %t",
					test.name,
					test.setValue,
					err,
					test.expectError,
				)
			}
		}
		got := pflag.Get()
		if got != test.want {
			t.Errorf("%s: String() = %q, want %q", test.name, got, test.want)
		}
		os.Setenv("ENVFLAGS_TEST", "")
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue bool
		setValue     string
		want         bool
		expectError  bool
	}{
		{
			name:         "env, no explicit value",
			envValue:     "true",
			defaultValue: false,
			setValue:     "_",
			want:         true,
			expectError:  false,
		},
		{
			name:         "env not set, no explicit value",
			envValue:     "",
			defaultValue: false,
			setValue:     "_",
			want:         false,
			expectError:  false,
		},
		{
			name:         "env not set, explicit value",
			envValue:     "",
			defaultValue: false,
			setValue:     "true",
			want:         true,
			expectError:  false,
		},
		{
			name:         "env, explicit value",
			envValue:     "true",
			defaultValue: false,
			setValue:     "false",
			want:         false,
			expectError:  false,
		},
		{
			name:         "env, empty explicit value",
			envValue:     "true",
			defaultValue: false,
			setValue:     "",
			want:         true,
			expectError:  true,
		},
		{
			name:         "malformed env, no explicit value",
			envValue:     "xyzabc",
			defaultValue: false,
			setValue:     "_",
			want:         false,
			expectError:  false,
		},
	}
	envKey := "ENVFLAGS_TEST"
	for _, test := range tests {
		os.Setenv(envKey, test.envValue)
		pflag := NewBool(envKey, test.defaultValue)
		if test.setValue != "_" {
			err := pflag.Set(test.setValue)
			if (err != nil) != test.expectError {
				t.Errorf(
					"%s: Set(%s) returned error %v, expected error %t",
					test.name,
					test.setValue,
					err,
					test.expectError,
				)
			}
		}
		got := pflag.Get()
		if got != test.want {
			t.Errorf("%s: Get() = %t, want %t", test.name, got, test.want)
		}
		os.Setenv("ENVFLAGS_TEST", "")
	}
}

func TestIsBoolFlag(t *testing.T) {
	envKey := "ENVFLAGS_TEST"
	pflag := NewBool(envKey, false)
	if !pflag.IsBoolFlag() {
		t.Errorf("IsBoolFlag() = false, want true")
	}
}

func TestSetDefault(t *testing.T) {
	envKey := "ENVFLAGS_TEST"
	pflag := NewBool(envKey, false)
	pflag.setDefault(envKey, true)
	if pflag.Get() != true {
		t.Errorf("setDefault() did not set value")
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue int
		setValue     string
		want         int
		expectError  bool
	}{
		{
			name:         "env, no explicit value",
			envValue:     "42",
			defaultValue: 0,
			setValue:     "_",
			want:         42,
			expectError:  false,
		},
		{
			name:         "env not set, no explicit value",
			envValue:     "",
			defaultValue: 0,
			setValue:     "_",
			want:         0,
			expectError:  false,
		},
		{
			name:         "env not set, explicit value",
			envValue:     "",
			defaultValue: 0,
			setValue:     "42",
			want:         42,
			expectError:  false,
		},
		{
			name:         "env, explicit value",
			envValue:     "42",
			defaultValue: 0,
			setValue:     "0",
			want:         0,
			expectError:  false,
		},
		{
			name:         "env, empty explicit value",
			envValue:     "42",
			defaultValue: 0,
			setValue:     "",
			want:         42,
			expectError:  true,
		},
		{
			name:         "malformed env, no explicit value",
			envValue:     "xyzabc",
			defaultValue: 0,
			setValue:     "_",
			want:         0,
			expectError:  true,
		},
	}
	envKey := "ENVFLAGS_TEST"
	for _, test := range tests {
		os.Setenv(envKey, test.envValue)
		pflag := NewInt(envKey, test.defaultValue)
		if test.setValue != "_" {
			err := pflag.Set(test.setValue)
			if (err != nil) != test.expectError {
				t.Errorf(
					"%s: Set(%s) returned error %v, expected error %t",
					test.name,
					test.setValue,
					err,
					test.expectError,
				)
			}
		}
		got := pflag.Get()
		if got != test.want {
			t.Errorf("%s: Get() = %d, want %d", test.name, got, test.want)
		}
		os.Setenv("ENVFLAGS_TEST", "")
	}
}

func TestDuration(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue time.Duration
		setValue     string
		want         time.Duration
		expectError  bool
	}{
		{
			name:         "env, no explicit value",
			envValue:     "42s",
			defaultValue: 10 * time.Second,
			setValue:     "_",
			want:         42 * time.Second,
			expectError:  false,
		},
		{
			name:         "env not set, no explicit value",
			envValue:     "",
			defaultValue: 30 * time.Second,
			setValue:     "_",
			want:         30 * time.Second,
			expectError:  false,
		},
		{
			name:         "env not set, explicit value",
			envValue:     "",
			defaultValue: 0 * time.Second,
			setValue:     "42s",
			want:         42 * time.Second,
			expectError:  false,
		},
		{
			name:         "env, explicit value",
			envValue:     "42s",
			defaultValue: 0 * time.Second,
			setValue:     "10s",
			want:         10 * time.Second,
			expectError:  false,
		},
		{
			name:         "env, empty explicit value",
			envValue:     "42s",
			defaultValue: 0 * time.Second,
			setValue:     "",
			want:         42 * time.Second,
			expectError:  true,
		},
		{
			name:         "malformed env, no explicit value",
			envValue:     "xyzabc",
			defaultValue: 30 * time.Second,
			setValue:     "_",
			want:         30 * time.Second,
			expectError:  true,
		},
	}
	envKey := "ENVFLAGS_TEST"
	for _, test := range tests {
		os.Setenv(envKey, test.envValue)
		pflag := NewDuration(envKey, test.defaultValue)
		if test.setValue != "_" {
			err := pflag.Set(test.setValue)
			if (err != nil) != test.expectError {
				t.Errorf(
					"%s: Set(%s) returned error %v, expected error %t",
					test.name,
					test.setValue,
					err,
					test.expectError,
				)
			}
		}
		got := pflag.Get()
		if got != test.want {
			t.Errorf("%s: Get() = %s, want %s", test.name, got, test.want)
		}
		os.Setenv("ENVFLAGS_TEST", "")
	}
}

func TestNewUint64(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue uint64
		setValue     string
		want         uint64
		expectError  bool
	}{
		{
			name:         "env, no explicit value",
			envValue:     "42",
			defaultValue: 0,
			setValue:     "_",
			want:         42,
			expectError:  false,
		},
		{
			name:         "env not set, no explicit value",
			envValue:     "",
			defaultValue: 10,
			setValue:     "_",
			want:         10,
			expectError:  false,
		},
		{
			name:         "env not set, explicit value",
			envValue:     "",
			defaultValue: 10,
			setValue:     "42",
			want:         42,
			expectError:  false,
		},
		{
			name:         "env, explicit value",
			envValue:     "42",
			defaultValue: 10,
			setValue:     "100",
			want:         100,
			expectError:  false,
		},
		{
			name:         "env, empty explicit value",
			envValue:     "42",
			defaultValue: 10,
			setValue:     "",
			want:         42,
			expectError:  true,
		},
		{
			name:         "malformed env, no explicit value",
			envValue:     "xyzabc",
			defaultValue: 10,
			setValue:     "_",
			want:         10,
			expectError:  true,
		},
	}
	envKey := "ENVFLAGS_TEST"
	for _, test := range tests {
		os.Setenv(envKey, test.envValue)
		pflag := NewUint64(envKey, test.defaultValue)
		if test.setValue != "_" {
			err := pflag.Set(test.setValue)
			if (err != nil) != test.expectError {
				t.Errorf(
					"%s: Set(%s) returned error %v, expected error %t",
					test.name,
					test.setValue,
					err,
					test.expectError,
				)
			}
		}
		got := pflag.Get()
		if got != test.want {
			t.Errorf("%s: Get() = %d, want %d", test.name, got, test.want)
		}
		os.Setenv("ENVFLAGS_TEST", "")
	}
}
