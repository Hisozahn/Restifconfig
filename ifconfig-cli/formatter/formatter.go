package formatter

import (
	"fmt"
	"strconv"
	"net"
	"github.com/Hisozahn/Restifconfig/ifconfig-service/rest"
	"errors"
	"bytes"
)

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

func PrintInterfaceInfo(info *rest.InterfaceInfoMsg) error {
	var buffer bytes.Buffer

	buffer.WriteString(info.Name + ":\t")
	if (info.HwAddr != "") {
		buffer.WriteString("Hw_addr: " + info.HwAddr + "\n\t")
	}
	if (len(info.InetAddrs) != 0) {
		for i := 0; i < len(info.InetAddrs); i++ {
			ip, ipnet, err := net.ParseCIDR(info.InetAddrs[i])
			if err != nil {
				return errors.New("Error parsing ip address: " + info.InetAddrs[i])
			}
			if (ip.To4() != nil) {
				buffer.WriteString("Ipv4: " + ip.String() + ", mask " + Ipv4MaskToDec(ipnet.Mask) + "\n\t")
			} else {
				buffer.WriteString("Ipv6: " + info.InetAddrs[i] + "\n\t")
			}
		}
	}
	buffer.WriteString("MTU:" + strconv.Itoa(info.MTU))
	fmt.Println(buffer.String())
	return nil
}
