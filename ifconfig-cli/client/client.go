package client

import (
	"time"
	"net/http"
	"github.com/Hisozahn/Restifconfig/ifconfig-service/rest"
	"errors"
	"encoding/json"
)

const timeout = time.Duration(5 * time.Second)
const supportedApiVersion = "v1"

var httpClient = &http.Client{Timeout: timeout}

func SupportedApiVersion(v string) bool {
	return v == supportedApiVersion
}


func CheckStatus(resp *http.Response) error {
	if (resp.StatusCode != http.StatusOK) {
		errorMsg := new(rest.ErrorMsg)
		err := json.NewDecoder(resp.Body).Decode(errorMsg)
		if (err != nil) {
			return errors.New("Incorrect error message from server")
		}
		return errors.New("API error: " + errorMsg.Error)
	}
	return nil
}

func GetVersion(addr string) (string, error) {
	resp, err := httpClient.Get(addr + "/version")
	if (err != nil) {
		return "", errors.New("Connection error: " + err.Error())
	}
	defer resp.Body.Close()
	err = CheckStatus(resp)
	if (err != nil) {
		return "", err
	}
	msg := new(rest.VersionMsg)
	err = json.NewDecoder(resp.Body).Decode(msg)
	if (err != nil) {
		return "", errors.New("Incorrect version message: " + err.Error())
	}
	if (!SupportedApiVersion(msg.Version)) {
		return "", errors.New("Unsupported api version: " + msg.Version)
	}
	return msg.Version, nil
}

func ListInterfaces(addr string) ([]string, error) {
	v, err := GetVersion(addr)
	if (err != nil) {
		return nil, err
	}
	resp, err := httpClient.Get(addr + "/" + v + "/interfaces")
	if (err != nil) {
		return nil, errors.New("Connection error: " + err.Error())
	}
	defer resp.Body.Close()
	err = CheckStatus(resp)
	if (err != nil) {
		return nil, err
	}
	interfaceListMsg := new(rest.InterfaceListMsg)
	err = json.NewDecoder(resp.Body).Decode(interfaceListMsg)
	if (err != nil) {
		return nil, errors.New("Incorrect interface list message: " + err.Error())
	}
	return interfaceListMsg.Interfaces, nil
}

func ShowInterface(addr string, name string) (*rest.InterfaceInfoMsg, error) {
	v, err := GetVersion(addr)
	if (err != nil) {
		return nil, err
	}
	resp, err := httpClient.Get(addr + "/" + v + "/interface/" + name)
	if (err != nil) {
		return nil, errors.New("Connection error: " + err.Error())
	}
	defer resp.Body.Close()
	err = CheckStatus(resp)
	if (err != nil) {
		return nil, err
	}
	interfaceInfoMsg := new(rest.InterfaceInfoMsg)
	err = json.NewDecoder(resp.Body).Decode(interfaceInfoMsg)
	if (err != nil) {
		return nil, errors.New("Incorrect interface info message: " + err.Error())
	}
	return interfaceInfoMsg, nil
}
