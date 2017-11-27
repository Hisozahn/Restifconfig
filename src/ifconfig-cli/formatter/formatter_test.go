package formatter

import (
	"testing"
	"ifconfig-service/rest"
)

func TestPrintInterfaceInfoSuccess(t *testing.T) {
	msg := rest.InterfaceInfoMsg{Name: "lo", MTU: 65536}
	err := PrintInterfaceInfo(&msg)
	if (err != nil) {
		t.Error(err)
	}
}


func TestPrintInterfaceInfoBadIpAddr(t *testing.T) {
	var msgs = []rest.InterfaceInfoMsg{
		rest.InterfaceInfoMsg{Name: "lo", InetAddrs: []string{""}, MTU: 65536},
		rest.InterfaceInfoMsg{Name: "lo", InetAddrs: []string{"ds ad wq wq"}, MTU: 65536},
		rest.InterfaceInfoMsg{Name: "lo", InetAddrs: []string{"333.333.222.333"}, MTU: 65536},
	}
	for i := range msgs {
		err := PrintInterfaceInfo(&msgs[i])
		if (err == nil) {
			t.Errorf("Expected error, msg: %v", msgs[i])
		}
	}
}

