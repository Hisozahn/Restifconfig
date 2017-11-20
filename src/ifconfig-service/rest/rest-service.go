package rest

import  (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"ifconfig-service/api"
)

type Response struct {
	Code int `json:Code`
	Content interface{} `json:Content`
}

const addrToListen string = ":55555"

func NewErrorResponse(err api.ApiError) (r Response) {
	content := map[string]string{"error": err.Info().Message}
	r = Response{Code: err.Info().Code, Content: content}
	return
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	v := api.GetVersion()
	content := map[string]string{"version": v}
	msg := Response{Code: 200, Content: content}
	json.NewEncoder(w).Encode(msg)
}

func GetInterfaces(w http.ResponseWriter, r *http.Request) {
	names,err := api.GetInterfaces(mux.Vars(r)["api_version"])
	var msg Response

	if err != nil {
		msg = NewErrorResponse(err)
	} else {
		content := map[string][]string{"interfaces": names}
		msg = Response{Code: 200, Content: content}
	}
	json.NewEncoder(w).Encode(msg)
}

func GetInterfaceInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	info,err := api.GetInterfaceInfo(vars["api_version"], vars["name"])
	var msg Response

	if err != nil {
		msg = NewErrorResponse(err)
	} else {
		msg = Response{Code: 200, Content: info}
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
