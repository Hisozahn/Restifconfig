package rest

import  (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"ifconfig-service/api"
)

const addrToListen string = ":55555"

func ApiToHttpCode(apiCode int) int {
	switch apiCode {
	case 404:
		return http.StatusNotFound
	case 500:
		return http.StatusInternalServerError
	case 200:
		return http.StatusOK
	default:
		return http.StatusNotImplemented
	}
}

func NewErrorResponse(err api.ApiError) (r interface{}) {
	r = ErrorMsg{Error: err.Info().Message}
	return
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	v := api.GetVersion()
	msg := VersionMsg{Version: v}
	json.NewEncoder(w).Encode(msg)
}

func GetInterfaces(w http.ResponseWriter, r *http.Request) {
	names,err := api.GetInterfaces(mux.Vars(r)["api_version"])
	var msg interface{}

	if err != nil {
		msg = NewErrorResponse(err)
		w.WriteHeader(ApiToHttpCode(err.Info().Code))
	} else {
		msg = InterfaceListMsg{Interfaces: names}
	}
	json.NewEncoder(w).Encode(msg)
}

func GetInterfaceInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	info,err := api.GetInterfaceInfo(vars["api_version"], vars["name"])
	var msg interface{}

	if err != nil {
		msg = NewErrorResponse(err)
		w.WriteHeader(ApiToHttpCode(err.Info().Code))
	} else {
		msg = InterfaceInfoMsg{Name: info.Name, HwAddr: info.HwAddr, InetAddrs: info.InetAddrs, MTU: info.MTU}
	}
	json.NewEncoder(w).Encode(msg)
}

func Listen() {
	router := mux.NewRouter()
	router.HandleFunc("/service/version", GetVersion).Methods("GET")
	router.HandleFunc("/service/{api_version}/interfaces", GetInterfaces).Methods("GET")
	router.HandleFunc("/service/{api_version}/interface/{name}", GetInterfaceInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(addrToListen, router))
}
