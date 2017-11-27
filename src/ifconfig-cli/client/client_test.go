package client

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"strconv"
)

const testVersion = "v1"

var getVersionHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"version": "` + testVersion +  `"}`)
}

func TestGetVersionSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getVersionHandler))
	defer ts.Close()

	v, err := GetVersion(ts.URL)

	if (err != nil) {
		t.Error(err)
	}
	if (v != testVersion) {
		t.Errorf("got: %v, expected: %v", v, testVersion)
	}
}

func TestGetVersionBadUrl(t *testing.T) {
	urls := []string{"", "not a url", "google.com", "12355"}
	for _,u := range urls {
		_, err := GetVersion(u)
		if (err == nil) {
			t.Errorf("Expected error, URL: %v", u)
		}
	}
}


const testInterface = "lo"

var getInterfacesHandler = func(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path == "/version") {
		fmt.Fprintf(w, `{"version": "` + testVersion +  `"}`)
	} else if (r.URL.Path == "/" + testVersion + "/interfaces") {
		fmt.Fprintf(w, `{"interfaces": ["` + testInterface +  `"]}`)
	}
}

func TestListInterfacesSuccess(t *testing.T) {
	ts:= httptest.NewServer(http.HandlerFunc(getInterfacesHandler))
	defer ts.Close()

	list, err := ListInterfaces(ts.URL)

	if (err != nil) {
		t.Error(err)
	}
	if (len(list) != 1 && list[0] != testInterface) {
		t.Errorf("got: %v, expected: %v", list, `["` + testInterface + `"]`)
	}
}

func TestListInterfacesBadUrl(t *testing.T) {
	urls := []string{"", "not a url", "google.com", "12355"}
	for _,u := range urls {
		_, err := ListInterfaces(u)
		if (err == nil) {
			t.Errorf("Expected error, URL: %v", u)
		}
	}
}

const testMTU = 65536

var getInterfaceInfoHandler = func(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path == "/version") {
		fmt.Fprintf(w, `{"version": "` + testVersion +  `"}`)
	} else if (r.URL.Path == "/" + testVersion + "/interface/" + testInterface) {
		fmt.Fprintf(w, `{"name":"` + testInterface + `", "MTU":` + strconv.Itoa(testMTU)  + `}`)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}


func TestShowInterfaceSuccess(t *testing.T) {
	ts:= httptest.NewServer(http.HandlerFunc(getInterfaceInfoHandler))
	defer ts.Close()

	info, err := ShowInterface(ts.URL, testInterface)

	if (err != nil) {
		t.Error(err)
	}
	if (info.Name != testInterface) {
		t.Errorf("Name: got: %v, expected: %v", info.Name, testInterface)
	}
	if (info.MTU != testMTU) {
		t.Errorf("MTU: fot: %v, expected: %v", info.MTU, testMTU)
	}
}

func TestShowInterfaceBadURL(t *testing.T) {
	urls := []string{"", "not a url", "google.com", "12355"}
	for _,u := range urls {
		_, err := ShowInterface(u, testInterface)
		if (err == nil) {
			t.Errorf("Expected error, URL: %v", u)
		}
	}
}

func TestShowInterfaceBadName(t *testing.T) {
	interfaces := []string{"", "not an interface", "32132143215", "''''''''''''''"}
	ts:= httptest.NewServer(http.HandlerFunc(getInterfaceInfoHandler))
	defer ts.Close()

	for _,i := range interfaces {
		_, err := ShowInterface(ts.URL, i)
		if (err == nil) {
			t.Errorf("Expected error, interface name: %v", i)
		}
	}
}
