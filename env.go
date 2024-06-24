// Package env provides functions for conveniently accessing environment variables.
package env

import (
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// String type constraint represents any string value.
type String interface {
	~string
}

// Boolean type constraint represents any bool value.
type Boolean interface {
	~bool
}

// Float type constraint represents any float value.
type Float interface {
	~float32 | ~float64
}

// Signed type constraint represents any signed int value.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned type constraint represents any unsigned int value.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// GetString returns the associated [String] value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present. A variable defined with an empty string wont return a default
// value.
func GetString[V String](key string, defaultValue V) V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return V(val)
}

// GetBool returns the associated [Boolean] value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseBool]
// for supported values.
func GetBool[V Boolean](key string, defaultValue V) V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedBool, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}

	return V(parsedBool)
}

// GetInt returns the associated [Signed] value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetInt[V Signed](key string, base int, defaultValue V) V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var parsedInt int64
	var err error

	switch any(defaultValue).(type) {
	case int:
		parsedInt, err = strconv.ParseInt(val, base, 0)
	case int8:
		parsedInt, err = strconv.ParseInt(val, base, 8)
	case int16:
		parsedInt, err = strconv.ParseInt(val, base, 16)
	case int32:
		parsedInt, err = strconv.ParseInt(val, base, 32)
	case int64:
		parsedInt, err = strconv.ParseInt(val, base, 64)
	default:
		parsedInt, err = strconv.ParseInt(val, base, 64)
	}
	if err != nil {
		return defaultValue
	}

	return V(parsedInt)
}

// GetUint returns the associated [Unsigned] value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUint[V Unsigned](key string, base int, defaultValue V) V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var parsedUint uint64
	var err error

	switch any(defaultValue).(type) {
	case uint:
		parsedUint, err = strconv.ParseUint(val, base, 0)
	case uint8:
		parsedUint, err = strconv.ParseUint(val, base, 8)
	case uint16:
		parsedUint, err = strconv.ParseUint(val, base, 16)
	case uint32:
		parsedUint, err = strconv.ParseUint(val, base, 32)
	case uint64:
		parsedUint, err = strconv.ParseUint(val, base, 64)
	default:
		parsedUint, err = strconv.ParseUint(val, base, 64)
	}
	if err != nil {
		return defaultValue
	}

	return V(parsedUint)
}

// GetFloat returns the associated [Float] value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseFloat]
// for supported values.
func GetFloat[V Float](key string, defaultValue V) V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var parsedFloat float64
	var err error

	switch any(defaultValue).(type) {
	case float32:
		parsedFloat, err = strconv.ParseFloat(val, 32)
	case float64:
		parsedFloat, err = strconv.ParseFloat(val, 64)
	default:
		parsedFloat, err = strconv.ParseFloat(val, 64)
	}
	if err != nil {
		return defaultValue
	}

	return V(parsedFloat)
}

// GetDuration returns the associated [time.Duration] value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [time.ParseDuration]
// for supported values.
func GetDuration(key string, defaultValue time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedDuration, err := time.ParseDuration(val)
	if err != nil {
		return defaultValue
	}

	return parsedDuration
}

// GetURL returns the associated [net/url.URL] value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [net/url.ParseRequestURI]
// for supported values.
func GetURL(key string, defaultValue url.URL) url.URL {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedURL, err := url.ParseRequestURI(val)
	if err != nil {
		return defaultValue
	}

	return *parsedURL
}

// GetStringSlice returns the associated [][String] values for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variables
// is not present.
func GetStringSlice[V String](key, separator string, defaultValue []V) []V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	stringVals := strings.Split(val, separator)
	stringSlice := make([]V, 0, len(stringVals))

	for _, strVal := range stringVals {
		stringSlice = append(stringSlice, V(strVal))
	}

	return stringSlice
}

// GetBoolSlice returns the associated [][Boolean] values for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variables
// is not present or any of the associated values could not be parsed. Refer to [strconv.ParseBool]
// for supported values.
func GetBoolSlice[V Boolean](key, separator string, defaultValue []V) []V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	stringVals := strings.Split(val, separator)

	boolSlice := make([]V, 0, len(stringVals))

	for _, boolStrVal := range stringVals {
		parsedBool, err := strconv.ParseBool(boolStrVal)
		if err != nil {
			return defaultValue
		}

		boolSlice = append(boolSlice, V(parsedBool))
	}

	return boolSlice
}

// GetIntSlice returns the associated [][Signed] values for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or any of the associated values could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetIntSlice[V Signed](key, separator string, base int, defaultValue []V) []V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	stringVals := strings.Split(val, separator)
	intSlice := make([]V, 0, len(stringVals))

	var parser func(string) (int64, error)

	var h V
	switch any(h).(type) {
	case int:
		parser = func(val string) (int64, error) {
			return strconv.ParseInt(val, base, 0)
		}
	case int8:
		parser = func(val string) (int64, error) {
			return strconv.ParseInt(val, base, 8)
		}
	case int16:
		parser = func(val string) (int64, error) {
			return strconv.ParseInt(val, base, 16)
		}
	case int32:
		parser = func(val string) (int64, error) {
			return strconv.ParseInt(val, base, 32)
		}
	case int64:
		parser = func(val string) (int64, error) {
			return strconv.ParseInt(val, base, 64)
		}
	default:
		parser = func(val string) (int64, error) {
			return strconv.ParseInt(val, base, 64)
		}
	}

	for _, intStrVal := range stringVals {
		parsedInt, err := parser(intStrVal)
		if err != nil {
			return defaultValue
		}

		intSlice = append(intSlice, V(parsedInt))
	}

	return intSlice
}

// GetUintSlice returns the associated [][Unsigned] values for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or any of the associated values could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUintSlice[V Unsigned](key, separator string, base int, defaultValue []V) []V {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	stringVals := strings.Split(val, separator)
	uintSlice := make([]V, 0, len(stringVals))

	var parser func(string) (uint64, error)

	var h V
	switch any(h).(type) {
	case uint:
		parser = func(val string) (uint64, error) {
			return strconv.ParseUint(val, base, 0)
		}
	case uint8:
		parser = func(val string) (uint64, error) {
			return strconv.ParseUint(val, base, 8)
		}
	case uint16:
		parser = func(val string) (uint64, error) {
			return strconv.ParseUint(val, base, 16)
		}
	case uint32:
		parser = func(val string) (uint64, error) {
			return strconv.ParseUint(val, base, 32)
		}
	case uint64:
		parser = func(val string) (uint64, error) {
			return strconv.ParseUint(val, base, 64)
		}
	default:
		parser = func(val string) (uint64, error) {
			return strconv.ParseUint(val, base, 64)
		}
	}

	for _, uintStrVal := range stringVals {
		parsedInt, err := parser(uintStrVal)
		if err != nil {
			return defaultValue
		}

		uintSlice = append(uintSlice, V(parsedInt))
	}

	return uintSlice
}

// GetDurationSlice returns the associated []time.Duration values for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variables
// is not present or any of the associated values could not be parsed. Refer to [time.ParseDuration]
// for supported values.
func GetDurationSlice(key, separator string, defaultValue []time.Duration) []time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	stringVals := strings.Split(val, separator)

	durationSlice := make([]time.Duration, 0, len(stringVals))

	for _, durationStrVal := range stringVals {
		parsedDuration, err := time.ParseDuration(durationStrVal)
		if err != nil {
			return defaultValue
		}

		durationSlice = append(durationSlice, parsedDuration)
	}

	return durationSlice
}

// GetURLSlice returns the associated [net/url.URL] values for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variables
// is not present or any of the associated values could not be parsed. Refer to [net/url.ParseRequestURI]
// for supported values.
func GetURLSlice(key, separator string, defaultValue []url.URL) []url.URL {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	stringVals := strings.Split(val, separator)

	urlSlice := make([]url.URL, 0, len(stringVals))

	for _, urlStrVal := range stringVals {
		parsedDuration, err := url.ParseRequestURI(urlStrVal)
		if err != nil {
			return defaultValue
		}

		urlSlice = append(urlSlice, *parsedDuration)
	}

	return urlSlice
}
