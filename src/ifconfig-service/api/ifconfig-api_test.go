package api

import "testing"

func TestGetInterfacesBadVersion(t *testing.T) {
	versions := []string{"", "v2", "vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv", "123", "\"", "'"}
	for _, v := range versions {
		_, err := GetInterfaces(v)
		if (err == nil) {
			t.Error("Expected error on version value = '", v, "'")
		}
		if (err.Info().Code != 500) {
			t.Error("Expected 500 error code")
		}
	}
}

func TestGetInterfacesSuccess(t *testing.T) {
	v := GetVersion()
	_, err := GetInterfaces(v)
	if (err != nil) {
		t.Error("Expected nil error, version = '", v, "'")
	}
}

func TestGetIntInfoBadVersion(t *testing.T) {
	versions := []string{"", "v2", "vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv", "123", "\"", "'"}
	existingInterfaceName := "lo"
	badInterfaceName := "I do not exist"
	for _, v := range versions {
		_, err := GetInterfaceInfo(v, existingInterfaceName)
		if (err == nil) {
			t.Error("Expected error passing version = '", v, "' and name = '", existingInterfaceName, "'")
		}
		if (err.Info().Code != 500) {
			t.Error("Expected 500 error code")
		}
		_, err = GetInterfaceInfo(v, badInterfaceName)
		if (err == nil) {
			t.Error("Expected error passing version = '", v, "' and name = '", badInterfaceName, "'")
		}
		if (err.Info().Code != 500) {
			t.Error("Expected 500 error code")
		}
	}
}

func TestGetIntInfoLoopback(t *testing.T) {
	v := GetVersion()
	_, err := GetInterfaceInfo(v, "lo")
	if (err != nil) {
		t.Error("Expected nil error passing loopback interface")
	}
}

func TestGetIntInfoNotFound(t *testing.T) {
	v := GetVersion()
	i := "no such interface"
	_, err := GetInterfaceInfo(v, i)
	if (err == nil) {
		t.Error("Expected error passing interface: ", i)
	}
	if (err.Info().Code != 404) {
		t.Error("Expected 404 error code passing interface: ", i)
	}
}
