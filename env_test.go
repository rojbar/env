package env

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

func equalSlices[T comparable](a []T, b []T) error {
	if len(a) != len(b) {
		return fmt.Errorf("not equal length %d %d", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return fmt.Errorf("not equal values in position %d", i)
		}
	}

	return nil
}

func TestGetString(t *testing.T) {
	defaultVal := "default"
	envKey := "KEY_STRING"

	val := GetString(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %s got %s", defaultVal, val)
	}

	newVal := "val"
	t.Setenv(envKey, newVal)

	val = GetString(envKey, defaultVal)
	if val != newVal {
		t.Errorf("expected env var %s value %s got %s", envKey, defaultVal, val)
	}
}

func TestGetBool(t *testing.T) {
	defaultVal := false
	envKey := "KEY_BOOL"

	val := GetBool(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %t got %t", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetBool(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %t got %t", defaultVal, val)
	}

	newVal := "t"
	t.Setenv(envKey, newVal)

	val = GetBool(envKey, defaultVal)
	if !val {
		t.Errorf("expected env var %s to have value %t got %t", envKey, defaultVal, val)
	}
}

func TestGetInt(t *testing.T) {
	defaultVal := 2
	base := 10
	envKey := "KEY_INT"

	val := GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt(envKey, base, defaultVal)
	if val != 3 {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt8(t *testing.T) {
	defaultVal := int8(2)
	base := 10
	envKey := "KEY_INT_8"

	val := GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt(envKey, base, defaultVal)
	if val != int8(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt16(t *testing.T) {
	defaultVal := int16(2)
	base := 10
	envKey := "KEY_INT_16"

	val := GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt(envKey, base, defaultVal)
	if val != int16(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt32(t *testing.T) {
	defaultVal := int32(2)
	base := 10
	envKey := "KEY_INT_32"

	val := GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt(envKey, base, defaultVal)
	if val != int32(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt64(t *testing.T) {
	defaultVal := int64(2)
	base := 10
	envKey := "KEY_INT_64"

	val := GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt(envKey, base, defaultVal)
	if val != int64(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

type customInt int

func TestGetIntCustom(t *testing.T) {
	defaultVal := customInt(2)
	base := 10
	envKey := "KEY_INT_CUSTOM"

	val := GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt(envKey, base, defaultVal)
	if val != customInt(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint(t *testing.T) {
	defaultVal := uint(2)
	base := 10
	envKey := "KEY_UINT"

	val := GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint(envKey, base, defaultVal)
	if val != uint(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint8(t *testing.T) {
	defaultVal := uint8(2)
	base := 10
	envKey := "KEY_UINT_8"

	val := GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint(envKey, base, defaultVal)
	if val != uint8(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint16(t *testing.T) {
	defaultVal := uint16(2)
	base := 10
	envKey := "KEY_UINT_16"

	val := GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint(envKey, base, defaultVal)
	if val != uint16(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint32(t *testing.T) {
	defaultVal := uint32(2)
	base := 10
	envKey := "KEY_UINT_32"

	val := GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint(envKey, base, defaultVal)
	if val != uint32(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint64(t *testing.T) {
	defaultVal := uint64(2)
	base := 10
	envKey := "KEY_UINT_64"

	val := GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint(envKey, base, defaultVal)
	if val != uint64(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

type customUint uint

func TestGetUintCustom(t *testing.T) {
	defaultVal := customUint(2)
	base := 10
	envKey := "KEY_UINT_CUSTOM"

	val := GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint(envKey, base, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint(envKey, base, defaultVal)
	if val != customUint(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetFloat32(t *testing.T) {
	defaultVal := float32(2)
	envKey := "KEY_FLOAT_32"

	val := GetFloat(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetFloat(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetFloat(envKey, defaultVal)
	if val != float32(3) {
		t.Errorf("expected env var %s value %f got %f", envKey, defaultVal, val)
	}
}

func TestGetFloat64(t *testing.T) {
	defaultVal := float64(2)
	envKey := "KEY_FLOAT_64"

	val := GetFloat(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetFloat(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetFloat(envKey, defaultVal)
	if val != float64(3) {
		t.Errorf("expected env var %s value %f got %f", envKey, defaultVal, val)
	}
}

type customFloat float32

func TestGetFloatCustom(t *testing.T) {
	defaultVal := customFloat(2)
	envKey := "KEY_FLOAT_CUSTOM"

	val := GetFloat(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetFloat(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetFloat(envKey, defaultVal)
	if val != customFloat(3) {
		t.Errorf("expected env var %s value %f got %f", envKey, defaultVal, val)
	}
}

func TestGetDuration(t *testing.T) {
	defaultVal := time.Second
	envKey := "KEY_DURATION"

	val := GetDuration(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetDuration(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "70s"
	t.Setenv(envKey, newVal)

	val = GetDuration(envKey, defaultVal)
	if val != time.Second*70 {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetURL(t *testing.T) {
	defaultVal, err := url.Parse("https://default.com/")
	if err != nil {
		t.Errorf("parse url failed %s", err.Error())
	}

	envKey := "KEY_URL"

	val := GetURL(envKey, *defaultVal)
	if val != *defaultVal {
		t.Errorf("expected default value %s got %s", defaultVal.String(), val.String())
	}

	t.Setenv(envKey, "invalid")

	val = GetURL(envKey, *defaultVal)
	if val != *defaultVal {
		t.Errorf("expected default value %s got %s", defaultVal.String(), val.String())
	}

	newVal := "https://rojbar.com/"
	t.Setenv(envKey, newVal)

	newValURL, err := url.Parse(newVal)
	if err != nil {
		t.Errorf("parse url failed %s", err.Error())
	}

	val = GetURL(envKey, *defaultVal)
	if val != *newValURL {
		t.Errorf("expected env var %s value %s got %s", envKey, defaultVal.String(), val.String())
	}
}
func TestGetStringSlice(t *testing.T) {
	defaultVal := []string{"1", "2", "3"}
	envKey := "KEY_STRING_SLICE"
	separator := ":"

	val := GetStringSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetStringSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, []string{"1", "4", "3"}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetBoolSlice(t *testing.T) {
	defaultVal := []bool{true, false, true}
	envKey := "KEY_BOOL_SLICE"
	separator := ":"

	val := GetBoolSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")
	val = GetBoolSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}

	newVal := "t:t:t"
	t.Setenv(envKey, newVal)

	val = GetBoolSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, []bool{true, true, true}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetIntSlice(t *testing.T) {
	defaultVal := []int{1, 2, 3}
	base := 10
	envKey := "KEY_INT_SLICE"
	separator := ":"

	val := GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []int{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetInt8Slice(t *testing.T) {
	defaultVal := []int8{1, 2, 3}
	base := 10
	envKey := "KEY_INT8_SLICE"
	separator := ":"

	val := GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []int8{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetInt16Slice(t *testing.T) {
	defaultVal := []int16{1, 2, 3}
	base := 10
	envKey := "KEY_INT16_SLICE"
	separator := ":"

	val := GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []int16{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetInt32Slice(t *testing.T) {
	defaultVal := []int32{1, 2, 3}
	base := 10
	envKey := "KEY_INT32_SLICE"
	separator := ":"

	val := GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []int32{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetInt64Slice(t *testing.T) {
	defaultVal := []int64{1, 2, 3}
	base := 10
	envKey := "KEY_INT64_SLICE"
	separator := ":"

	val := GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []int64{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetIntCustomSlice(t *testing.T) {
	defaultVal := []customInt{customInt(1), customInt(2), customInt(3)}
	base := 10
	envKey := "KEY_INT64_SLICE"
	separator := ":"

	val := GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetIntSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []customInt{customInt(1), customInt(4), customInt(3)}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetUintSlice(t *testing.T) {
	defaultVal := []uint{1, 2, 3}
	base := 10
	envKey := "KEY_UINT_SLICE"
	separator := ":"

	val := GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []uint{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetUint8Slice(t *testing.T) {
	defaultVal := []uint8{1, 2, 3}
	base := 10
	envKey := "KEY_UINT8_SLICE"
	separator := ":"

	val := GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []uint8{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetUint16Slice(t *testing.T) {
	defaultVal := []uint16{1, 2, 3}
	base := 10
	envKey := "KEY_UINT16_SLICE"
	separator := ":"

	val := GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []uint16{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetUint32Slice(t *testing.T) {
	defaultVal := []uint32{1, 2, 3}
	base := 10
	envKey := "KEY_UINT32_SLICE"
	separator := ":"

	val := GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []uint32{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetUint64Slice(t *testing.T) {
	defaultVal := []uint64{1, 2, 3}
	base := 10
	envKey := "KEY_UINT64_SLICE"
	separator := ":"

	val := GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []uint64{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetUintCustomSlice(t *testing.T) {
	defaultVal := []customUint{1, 2, 3}
	base := 10
	envKey := "KEY_UINT_CUSTOM_SLICE"
	separator := ":"

	val := GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	newVal := "1:4:3"
	t.Setenv(envKey, newVal)

	val = GetUintSlice(envKey, separator, base, defaultVal)
	if err := equalSlices(val, []customUint{1, 4, 3}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetDurationSlice(t *testing.T) {
	defaultVal := []time.Duration{time.Hour, time.Microsecond, time.Second}
	envKey := "KEY_TIME_DURATION_SLICE"
	separator := ":"

	val := GetDurationSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")
	val = GetDurationSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}

	newVal := "10s:20µs:5h"
	t.Setenv(envKey, newVal)

	val = GetDurationSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, []time.Duration{time.Second * 10, time.Microsecond * 20, time.Hour * 5}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func TestGetURLSlice(t *testing.T) {
	rojbarURl, err := url.ParseRequestURI("https://rojbar.com/")
	if err != nil {
		t.Errorf("parse request failed %s", err.Error())
	}

	exampleURL, err := url.ParseRequestURI("https://example.com/")
	if err != nil {
		t.Errorf("parse request failed %s", err.Error())
	}

	ex, err := url.ParseRequestURI("https://ex.com/")
	if err != nil {
		t.Errorf("parse request failed %s", err.Error())
	}

	am, err := url.ParseRequestURI("https://am.com/")
	if err != nil {
		t.Errorf("parse request failed %s", err.Error())
	}

	defaultVal := []url.URL{*rojbarURl, *exampleURL}
	envKey := "KEY_URL_SLICE"
	separator := ";"

	val := GetURLSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected default value %s", err.Error())
	}

	t.Setenv(envKey, "invalid")
	val = GetURLSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, defaultVal); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}

	newVal := "https://ex.com/;https://am.com/"
	t.Setenv(envKey, newVal)

	val = GetURLSlice(envKey, separator, defaultVal)
	if err := equalSlices(val, []url.URL{*ex, *am}); err != nil {
		t.Errorf("expected env var value %s", err.Error())
	}
}

func BenchmarkGetString(b *testing.B) {
	defaultVal := "default"
	envKey := "KEY_STRING"

	for i := 0; i < b.N; i++ {
		val := GetString(envKey, defaultVal)
		if val != defaultVal {
			b.Errorf("expected default value %s got %s", defaultVal, val)
		}
	}
}
