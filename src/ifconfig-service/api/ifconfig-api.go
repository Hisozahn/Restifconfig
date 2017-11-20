package api

import (
	"net"
)

type InterfaceInfo struct {
	Name string `json:"name"`
	HwAddr string `json:"hw_addr,omitempty"`
	InetAddrs []string `json:"inet_addr,omitempty"`
	MTU int `json:"MTU"`
}

const version string = "v1"

func VersionSupported(v string) (bool) {
	return v == GetVersion()
}

func GetVersion() (v string) {
	return version
}

func GetInterfaces(v string) ([]string, ApiError) {
	if (!VersionSupported(v)) {
		return nil, NewError("Unsupported api version", 500)
	}
	interfaces,err := net.Interfaces()
	if (err != nil) {
		return nil, NewError(err.Error(), 500)
	}
	names := make([]string, len(interfaces))
	for i := range interfaces {
		names[i] = interfaces[i].Name
	}
	return names,nil
}

func GetInterfaceInfo(v string, name string) (InterfaceInfo, ApiError) {
	var info InterfaceInfo
	if (!VersionSupported(v)) {
		return info, NewError("Unsupported api version", 500)
	}
	interfaceByName,err := net.InterfaceByName(name)
	if err != nil {
		return info, NewError("interface " + name + " was not found", 404)
	}
	addresses,err := interfaceByName.Addrs()
	if err != nil {
		return info, NewError(err.Error(), 500)
	}
	addressStrings := make ([]string, len(addresses))
	for i := range addresses {
		addressStrings[i] = addresses[i].String()
	}
	info.Name = interfaceByName.Name
	info.MTU = interfaceByName.MTU
	info.HwAddr = interfaceByName.HardwareAddr.String()
	info.InetAddrs = addressStrings
	return info, nil
}
