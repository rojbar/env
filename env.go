// Package env provides functions for conveniently accessing environment variables.
package env

import (
	"net/url"
	"os"
	"strconv"
	"time"
)

// GetString returns the associated string value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present. A variable defined with an empty string wont return a default
// value.
func GetString(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}

// GetBool returns the associated bool value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseBool]
// for supported values.
func GetBool(key string, defaultValue bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedBool, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}

	return parsedBool
}

// GetInt returns the associated int value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetInt(key string, defaultValue int) int {
	return int(GetInteger(key, 10, 0, int64(defaultValue)))
}

// GetInt8 returns the associated int8 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetInt8(key string, defaultValue int8) int8 {
	return int8(GetInteger(key, 10, 8, int64(defaultValue)))
}

// GetInt16 returns the associated int16 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetInt16(key string, defaultValue int16) int16 {
	return int16(GetInteger(key, 10, 16, int64(defaultValue)))
}

// GetInt32 returns the associated int32 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetInt32(key string, defaultValue int32) int32 {
	return int32(GetInteger(key, 10, 32, int64(defaultValue)))
}

// GetInt64 returns the associated int64 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetInt64(key string, defaultValue int64) int64 {
	return GetInteger(key, 10, 64, defaultValue)
}

// GetInteger returns the associated integer value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseInt]
// for supported values.
func GetInteger(key string, base, bitSize int, defaultValue int64) int64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedInt, err := strconv.ParseInt(val, base, bitSize)
	if err != nil {
		return defaultValue
	}

	return parsedInt
}

// GetUint returns the associated uint value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUint(key string, defaultValue uint) uint {
	return uint(GetUInteger(key, 10, 0, uint64(defaultValue)))
}

// GetUint8 returns the associated uint8 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUint8(key string, defaultValue uint8) uint8 {
	return uint8(GetUInteger(key, 10, 8, uint64(defaultValue)))
}

// GetUint16 returns the associated uint16 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUint16(key string, defaultValue uint16) uint16 {
	return uint16(GetUInteger(key, 10, 16, uint64(defaultValue)))
}

// GetUint32 returns the associated uint32 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUint32(key string, defaultValue uint32) uint32 {
	return uint32(GetUInteger(key, 10, 32, uint64(defaultValue)))
}

// GetUint64 returns the associated uint64 value for the provided environment
// variable name. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUint64(varName string, defaultValue uint64) uint64 {
	return GetUInteger(varName, 10, 64, defaultValue)
}

// GetUInteger returns the associated unsigned integer value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseUint]
// for supported values.
func GetUInteger(key string, base, bitSize int, defaultValue uint64) uint64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedUint, err := strconv.ParseUint(val, base, bitSize)
	if err != nil {
		return defaultValue
	}

	return parsedUint
}

// GetFloat32 returns the associated float32 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseFloat]
// for supported values.
func GetFloat32(key string, defaultValue float32) float32 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedFloat, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return defaultValue
	}

	return float32(parsedFloat)
}

// GetFloat64 returns the associated float64 value for the provided environment
// variable named by the key. The defaultValue is returned only if the environment variable
// is not present or the associated value could not be parsed. Refer to [strconv.ParseFloat]
// for supported values.
func GetFloat64(key string, defaultValue float64) float64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	parsedFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultValue
	}

	return parsedFloat
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
