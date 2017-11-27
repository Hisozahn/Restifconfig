package rest

import (
    "net/http"
    "net/http/httptest"
    "testing"
	"fmt"
	"encoding/json"
	"reflect"
	"ifconfig-service/api"
	"github.com/gorilla/mux"
)


func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/service/version", GetVersion).Methods("GET")
	r.HandleFunc("/service/{api_version}/interfaces", GetInterfaces).Methods("GET")
	r.HandleFunc("/service/{api_version}/interface/{name}", GetInterfaceInfo).Methods("GET")
	r.SkipClean(true)
	return r
}

func HandlerTestHelper(t *testing.T, url string, expectedStatus int, expectedBody string) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

    if status := rr.Code; status != expectedStatus {
        t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
    }
	if (expectedBody != "") {
		equal, err := AreEqualJSON(rr.Body.String(), expectedBody)

		if (err != nil) {
			t.Error(err)
		}
		if (!equal) {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
		}
	}
}

func TestGetVersionHandler(t *testing.T) {
	HandlerTestHelper(t, "/service/version", http.StatusOK, `{"version":"v1"}`)
}

func TestGetInterfacesBadVersion(t *testing.T) {
	versions := []string{"v", "vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv", "123", "\"", "'"}
	for _, v := range versions {
		HandlerTestHelper(t, "/service/" + v + "/interfaces", http.StatusInternalServerError, "")
	}
	HandlerTestHelper(t, "/service//interfaces", http.StatusNotFound, "")
}

func TestGetInterfacesSuccess(t *testing.T) {
	v := api.GetVersion()
	HandlerTestHelper(t, "/service/" + v + "/interfaces", http.StatusOK, "")
}

func TestGetInterfaceInfoBadVersion(t *testing.T) {
	versions := []string{"v", "vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv", "123", "\"", "'"}
	for _, v := range versions {
		HandlerTestHelper(t, "/service/" + v + "/interface/no such interface", http.StatusInternalServerError, "")
		HandlerTestHelper(t, "/service/" + v + "/interface/lo", http.StatusInternalServerError, "")
	}
	HandlerTestHelper(t, "/service//interface/no such interface", http.StatusNotFound, "")
	HandlerTestHelper(t, "/service//interface/lo", http.StatusNotFound, "")
}

func TestGetInterfaceNotFound(t *testing.T) {
	v := api.GetVersion()
	HandlerTestHelper(t, "/service/" + v + "/interface/no such interface", http.StatusNotFound, "")
}

func TestGetInterfaceSuccess(t *testing.T) {
	v := api.GetVersion()
	HandlerTestHelper(t, "/service/" + v + "/interface/lo", http.StatusOK, "")
}
