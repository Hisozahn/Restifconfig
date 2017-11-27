package rest

type InterfaceInfoMsg struct {
	Name string `json:"name"`
	HwAddr string `json:"hw_addr,omitempty"`
	InetAddrs []string `json:"inet_addr,omitempty"`
	MTU int `json:"MTU"`
}

type VersionMsg struct {
	Version string `json:"version"`
}

type InterfaceListMsg struct {
	Interfaces []string `json:"interfaces"`
}

type ErrorMsg struct {
	Error string `json:"error"`
}
