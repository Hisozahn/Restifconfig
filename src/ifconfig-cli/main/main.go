package main

import (
	"net"
	"strings"
	"fmt"
	"os"
	"net/http"
	"time"
	"encoding/json"
	"ifconfig-cli/rest"
	"strconv"
)

const (
	commandOptionsCount = 2
	connectionExitCode = 1
	apiExitCode = 2
	parametersExitCode = 3
	successExitCode = 0
	version = "v1"
)

var tr = &http.Transport {
	IdleConnTimeout: 30 * time.Second,
}
var client = &http.Client{Transport: tr}

var server string = "localhost"
var port int = 55555
var printVersion bool = false
var addr string
var command string = "none"
var commandArguments []string

func ParseGlobalOptions(index int) (nextIndex int) {
	if (len(os.Args) <= index) {
		UsageError()
	}
	if (os.Args[index] == "--version" || os.Args[index] == "-version") {
		printVersion = true
		return index + 1
	}
	printVersion = false
	return index
}

func ParseCommand(index int) (nextIndex int) {
	if (len(os.Args) <= index) {
		UsageError()
	}
	switch os.Args[index] {
	case "list":
		command = "list"
		return index + 1
	case "show":
		if (len(os.Args) <= index + 1) {
			UsageError()
		}
		command = "show"
		commandArguments = make([]string, 1)
		commandArguments[0] = os.Args[index + 1]
		return index + 2
	default:
		UsageError()
	}
	return index
}

func ParseCommandOptions(index int) (nextIndex int) {
	for i := 0; i < commandOptionsCount ; i++ {
		if (len(os.Args) <= index + 1) {
			UsageError()
		}
		var err error
		switch os.Args[index] {
		case "-port","--port":
			port,err = strconv.Atoi(os.Args[index + 1])
			if (err != nil) {
				UsageError()
			}
		case "-server","--server":
			server = os.Args[index + 1]
		default:
			UsageError()
		}
		index += 2
	}
	return index
}

func ExecuteCommand() {
	switch command {
	case "list":
		ListInterfaces()
	case "show":
		ShowInterface(commandArguments[0])
	}
}

func main() {
	var index = 1
	index = ParseGlobalOptions(index)
	index = ParseCommand(index)
	index = ParseCommandOptions(index)
	if index != len(os.Args) {
		UsageError()
	}
	addr = "http://" + server + ":" + strconv.Itoa(port) + "/service"
	ExecuteCommand()
	os.Exit(successExitCode)
}

func UsageError() {
	fmt.Fprintln(os.Stderr, "Usage")
	os.Exit(parametersExitCode)
}


func SupportedApiVersion(v string) bool {
	return v == version
}



func ExitOnBadStatus(resp *http.Response) {
	if (resp.StatusCode != http.StatusOK) {
		errorMsg := new(rest.ErrorMsg)
		err := json.NewDecoder(resp.Body).Decode(errorMsg)
		if (err != nil) {
			ProtocolError(err)
		}
		fmt.Fprintf(os.Stderr, "API error: %s\n", errorMsg.Error)
		os.Exit(apiExitCode)
	}
}

func ConnectionError(err error) {
	fmt.Fprintln(os.Stderr, "Connection error: " + err.Error())
	os.Exit(connectionExitCode)
}

func ProtocolError(err error) {
	fmt.Fprintln(os.Stderr, "Protocol mismatch: " + err.Error())
	os.Exit(apiExitCode)
}


func GetVersion() string {
	resp, err := client.Get(addr + "/version")
	if (err != nil) {
		ConnectionError(err)
	}
	defer resp.Body.Close()
	ExitOnBadStatus(resp)
	msg := new(rest.VersionMsg)
	err = json.NewDecoder(resp.Body).Decode(msg)
	if (err != nil) {
		ProtocolError(err)
	}
	if (!SupportedApiVersion(msg.Version)) {
		fmt.Fprintln(os.Stderr, "Unsupported api version: " + msg.Version)
		os.Exit(apiExitCode)
	}
	return msg.Version
}

func ListInterfaces() {
	v := GetVersion()
	resp, err := client.Get(addr + "/" + v + "/interfaces")
	if (err != nil) {
		ConnectionError(err)
	}
	defer resp.Body.Close()
	ExitOnBadStatus(resp)
	interfaceListMsg := new(rest.InterfaceListMsg)
	err = json.NewDecoder(resp.Body).Decode(interfaceListMsg)
	if (err != nil) {
		ProtocolError(err)
	}
	fmt.Println(strings.Join(interfaceListMsg.Interfaces, ", "))
}

func Ipv4MaskToDec(mask net.IPMask) (result string) {
	result = ""
	for i := 0; i < len(mask); i++ {
		result += strconv.Itoa(int(mask[i]))
		if (i != len(mask) - 1) {
			result += "."
		}
	}
	return
}

func PrintInterfaceInfo(info *rest.InterfaceInfoMsg) {
	fmt.Print(info.Name + ":\t")
	if (info.HwAddr != "") {
		fmt.Print("Hw_addr: " + info.HwAddr + "\n\t")
	}
	if (len(info.InetAddrs) != 0) {
		for i := 0; i < len(info.InetAddrs); i++ {
			ip, ipnet, err := net.ParseCIDR(info.InetAddrs[i])
			if err != nil {
				ProtocolError(err)
			}
			if (ip.To4() != nil) {
				fmt.Print("Ipv4: ", ip.String() + ", mask ", Ipv4MaskToDec(ipnet.Mask) + "\n\t")
			} else {
				fmt.Print("Ipv6: ", ipnet.String() + "\n\t")
			}
		}
	}
	fmt.Println("MTU:", info.MTU)
}

func ShowInterface(name string) {
	v := GetVersion()
	resp, err := client.Get(addr + "/" + v + "/interface/" + name)
	if (err != nil) {
		ConnectionError(err)
	}
	defer resp.Body.Close()
	ExitOnBadStatus(resp)
	interfaceInfoMsg := new(rest.InterfaceInfoMsg)
	err = json.NewDecoder(resp.Body).Decode(interfaceInfoMsg)
	if (err != nil) {
		ProtocolError(err)
	}
	PrintInterfaceInfo(interfaceInfoMsg)
}
