package profile

import (
	"os"
	"strings"
	"testing"
)

// TestGetStringValue test
func TestGetStringValue(t *testing.T) {
	value := GetStringValue("test.data.string")
	if value != "testValue" {
		t.Fatal("wrong value retured")
	}
}

// TestGetIntValue test
func TestGetIntValue(t *testing.T) {
	value := GetIntValue("test.data.int")
	if value != 123 {
		t.Fatal("wrong value retured")
	}
}

// TestGetFloatValue test
func TestGetFloatValue(t *testing.T) {
	value := GetFloatValue("test.data.float")
	if value != 3.6 {
		t.Fatal("wrong value retured")
	}
}

// TestGetArray test
func TestGetArray(t *testing.T) {
	values := GetArrayValues("test.array")
	if len(values) != 3 {
		t.Fatal("wrong value retured")
	}

	for _, value := range values {
		if !strings.HasPrefix(value.(string), "value") {
			t.Fatal("wrong value returned - ")
		}
	}
}

// TestGetBoolean test
func TestGetBoolean(t *testing.T) {
	value := GetBooleanValue("test.data.boolean")
	if !value {
		t.Fatal("wrong value retured")
	}
}

// TestGetNotExistingStringValue test
func TestGetNotExistingStringValue(t *testing.T) {
	value := GetStringValue("not.exists")
	if len(value) > 0 {
		t.Fatal("wrong value retured")
	}
}

// TestGetNotExistingIntValue test
func TestGetNotExistingIntValue(t *testing.T) {
	value := GetIntValue("not.exists")
	if value != 0 {
		t.Fatal("wrong value retured")
	}
}

// TestGetNotExistingFloatValue test
func TestGetNotExistingFloatValue(t *testing.T) {
	value := GetFloatValue("not.exists")
	if value != 0.0 {
		t.Fatal("wrong value retured")
	}
}

// TestGetotExistingArray test
func TestGetNotExistingArray(t *testing.T) {
	value := GetArrayValues("not.exists")
	if len(value) != 0 {
		t.Fatal("wrong value retured")
	}
}

// TestGetNotExistingBoolean test
func TestGetNotExistingBoolean(t *testing.T) {
	value := GetBooleanValue("not.exists")
	if value {
		t.Fatal("wrong value retured")
	}
}

// TestGetEmptyString test
func TestGetEmptyString(t *testing.T) {
	value := GetStringValue("test.data.emptyString")
	if len(value) > 0 {
		t.Fatal("wrong value retured")
	}
}

// TestGetDefaultValueWithString test
func TestGetDefaultValueWithString(t *testing.T) {
	value := GetValueWithDefault("not.exists", "defaultValue")
	if value != "defaultValue" {
		t.Fatalf("wrong value retured")
	}
}

// TestGetDefaultValueWithString test
func TestGetDefaultValueWithEmptyString(t *testing.T) {
	value := GetValueWithDefault("test.data.emptyString", "defaultValue")
	if value != "defaultValue" {
		t.Fatalf("wrong value retured")
	}
}

// TestGetDefaultValueWithString test
func TestGetDefaultValueWithUnkownType(t *testing.T) {
	value := GetValueWithDefault("test.data.emptyString", true)
	if !value.(bool) {
		t.Fatalf("wrong value retured")
	}
}

// TestGetDefaultValueWithInt test
func TestGetDefaultValueWithInt(t *testing.T) {
	value := GetValueWithDefault("not.exists", 5)
	if value != 5 {
		t.Fatalf("wrong value retured")
	}
}

// TestGetDefaultValueWithFloat test
func TestGetDefaultValueWithFloat(t *testing.T) {
	value := GetValueWithDefault("not.exists", 4.5)
	if value != 4.5 {
		t.Fatalf("wrong value retured")
	}
}

// TestGetPath test
func TestGetPath(t *testing.T) {
	value := GetStringValue("test.data.path")
	if value != "/etc/hosts" {
		t.Fatalf("wrong value retured")
	}
}

// TestGetURL test
func TestGetURL(t *testing.T) {
	value := GetStringValue("test.data.url")
	if value != "http://127.0.0.1:8080/test" {
		t.Fatalf("wrong value retured")
	}
}

// TestClearData test
func TestClearData(t *testing.T) {
	ClearData()
	GetStringValue("test.data.string")
}

// TestSetLogLevel test
func TestSetLogLevel(t *testing.T) {
	SetLogLevel("debug")
	SetLogLevel("info")
	SetLogLevel("warn")
	SetLogLevel("error")
	ClearData()
	GetStringValue("not.exists")
}

// TestUseDifferentProfile test
func TestUseDifferentProfile(t *testing.T) {
	defaultProfile := GetStringValue("profile.name")
	if defaultProfile != "default" {
		t.Fatal("start with wrong profile")
	}
	os.Setenv("PROFILE", "dev")
	ClearData()
	devProfile := GetStringValue("profile.name")
	if devProfile != "dev" {
		t.Fatal("profile name not changed")
	}
	os.Unsetenv("PROFILE")
}

// TestUseNotExistingProfile test
func TestUseNotExistingProfile(t *testing.T) {
	os.Setenv("PROFILE", "qs")
	ClearData()
	value := GetStringValue("profile.name")
	if len(value) > 0 {
		t.Fatal("should not return a value")
	}
}

// TestUseNotExistingProfile test
func TestUsePathVariable(t *testing.T) {
	os.Setenv("CONFIG_FOLDER_PATH", "./")
	ClearData()
	value := GetStringValue("profile.name")
	if len(value) < 0 {
		t.Fatal("should not return a value")
	}
}
