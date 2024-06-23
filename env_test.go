package env

import (
	"net/url"
	"testing"
	"time"
)

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
	envKey := "KEY_INT"

	val := GetInt(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt(envKey, defaultVal)
	if val != 3 {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt8(t *testing.T) {
	defaultVal := int8(2)
	envKey := "KEY_INT_8"

	val := GetInt8(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt8(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt8(envKey, defaultVal)
	if val != int8(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt16(t *testing.T) {
	defaultVal := int16(2)
	envKey := "KEY_INT_16"

	val := GetInt16(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt16(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt16(envKey, defaultVal)
	if val != int16(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt32(t *testing.T) {
	defaultVal := int32(2)
	envKey := "KEY_INT_32"

	val := GetInt32(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt32(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt32(envKey, defaultVal)
	if val != int32(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInt64(t *testing.T) {
	defaultVal := int64(2)
	envKey := "KEY_INT_64"

	val := GetInt64(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInt64(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetInt64(envKey, defaultVal)
	if val != int64(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetInteger(t *testing.T) {
	defaultVal := int64(1)
	envKey := "KEY_INTEGER"

	val := GetInteger(envKey, 2, 0, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetInteger(envKey, 2, 0, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "0"
	t.Setenv(envKey, newVal)

	val = GetInteger(envKey, 2, 0, defaultVal)
	if val != 0 {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint(t *testing.T) {
	defaultVal := uint(2)
	envKey := "KEY_UINT"

	val := GetUint(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint(envKey, defaultVal)
	if val != uint(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint8(t *testing.T) {
	defaultVal := uint8(2)
	envKey := "KEY_UINT_8"

	val := GetUint8(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint8(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint8(envKey, defaultVal)
	if val != uint8(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint16(t *testing.T) {
	defaultVal := uint16(2)
	envKey := "KEY_UINT_16"

	val := GetUint16(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint16(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint16(envKey, defaultVal)
	if val != uint16(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint32(t *testing.T) {
	defaultVal := uint32(2)
	envKey := "KEY_UINT_32"

	val := GetUint32(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint32(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint32(envKey, defaultVal)
	if val != uint32(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUint64(t *testing.T) {
	defaultVal := uint64(2)
	envKey := "KEY_UINT_64"

	val := GetUint64(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUint64(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetUint64(envKey, defaultVal)
	if val != uint64(3) {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetUInteger(t *testing.T) {
	defaultVal := uint64(1)
	envKey := "KEY_UINTEGER"

	val := GetUInteger(envKey, 2, 0, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetUInteger(envKey, 2, 0, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %d got %d", defaultVal, val)
	}

	newVal := "0"
	t.Setenv(envKey, newVal)

	val = GetUInteger(envKey, 2, 0, defaultVal)
	if val != 0 {
		t.Errorf("expected env var %s value %d got %d", envKey, defaultVal, val)
	}
}

func TestGetFloat32(t *testing.T) {
	defaultVal := float32(2)
	envKey := "KEY_FLOAT_32"

	val := GetFloat32(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetFloat32(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetFloat32(envKey, defaultVal)
	if val != float32(3) {
		t.Errorf("expected env var %s value %f got %f", envKey, defaultVal, val)
	}
}

func TestGetFloat64(t *testing.T) {
	defaultVal := float64(2)
	envKey := "KEY_FLOAT_64"

	val := GetFloat64(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	t.Setenv(envKey, "invalid")

	val = GetFloat64(envKey, defaultVal)
	if val != defaultVal {
		t.Errorf("expected default value %f got %f", defaultVal, val)
	}

	newVal := "3"
	t.Setenv(envKey, newVal)

	val = GetFloat64(envKey, defaultVal)
	if val != float64(3) {
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
