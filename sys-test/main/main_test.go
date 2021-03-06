package main

import (
	"path"
	"strings"
	"fmt"
	"os"
	"testing"
	"os/exec"
)

const (
	serverFlag = "localhost"
	portFlag = "1111"
	addr = serverFlag + ":" + portFlag

	apiError = "API error"
	incorrectUsage = "Incorrect Usage"
	usage = "NAME:"
	incorrectCommand = "No help topic for"
	showWithoutArgument = "Empty argument list"
	version = "main version"
	connectionError = "Connection error"

	cliBinary = "./cli"
	serverBinary = "./service"
	binaryPath = "src/github.com/Hisozahn/Restifconfig/bin/"
)

var	goPath = os.Getenv("GOPATH")


func TestMain(m *testing.M) {
	err := os.Chdir(path.Join(goPath, binaryPath))
	if (err != nil) {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}
	execServer := exec.Command(serverBinary, addr)
	err = execServer.Start()
	if (err != nil) {
		fmt.Printf("could not start server: %v", err)
	}

	code := m.Run()

	err = execServer.Process.Kill()
	if (err != nil) {
		fmt.Printf("could not terminate server: %v", err)
	}
	os.Exit(code)
}

func TestSystem(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exitCode int
		outputPrefix string
	}{
		{"no arguments", []string{}, 0, usage},
		{"empty argument", []string{""}, 0, usage},
		{"incorrect command", []string{"no such command"}, 1, incorrectCommand},
		{"incorrect command with garbage", []string{"no such command", "grw", "trwrefwa", "gfdfdz"}, 1, incorrectCommand},
		{"incorrect command with bad flag", []string{"no such command", "--grw", "trwrefwa", "gfdfdz"}, 1, incorrectCommand},
		{"incorrect flag", []string{"--nosuchflag"}, 0, incorrectUsage},
		{"incorrect flag with garbage", []string{"--nosuchflag", "wqew", "uytuytuj", "jhffgd"}, 0, incorrectUsage},
		{"version", []string{"--version"}, 0, version},
		{"version with garbage", []string{"--version", "wqew", "uytuytuj", "jhffgd"}, 0, version},
		{"version with incorrect flag", []string{"--version", "--wqew", "uytuytuj", "jhffgd"}, 0, incorrectUsage},
		{"list command with incorrect flag", []string{"list", "--wqew", "uytuytuj", "jhffgd"}, 0, incorrectUsage},
		{"list command", []string{"--port", portFlag, "--server", serverFlag, "list"}, 0, "lo,"},
		{"google server flag with list command", []string{"--server", "google.com", "--port", portFlag, "list"}, 1, connectionError},
		{"valid server and wrong port flags with list command", []string{"--server", serverFlag, "--port", "0", "list"}, 1, connectionError},
		{"missed server flag value with list command",[]string{"--server", "--port", portFlag, "list"}, 1, incorrectCommand},
		{"invalid port value with list command",[]string{"--server", serverFlag, "--port", "12312321", "list"}, 1, connectionError},
		{"wrong port value format with list command",[]string{"--server", serverFlag, "--port", "dsadqw", "list"}, 0, incorrectUsage},
		{"show command", []string{"--port", portFlag, "--server", serverFlag, "show","lo"}, 0, "lo:"},
		{"missed show command argument", []string{"--port", portFlag, "--server", serverFlag, "show"}, 1, showWithoutArgument},
		{"show command with invalid interface", []string{"--port", portFlag, "--server", serverFlag, "show","no such interface"}, 1, apiError},
		{"google server flag with show command", []string{"--port", portFlag, "--server", "google.com", "show", "lo"}, 1, connectionError},
		{"missed server flag value with show command",[]string{"--server", "--port", portFlag, "show", "lo"}, 1, incorrectCommand},
		{"invalid port value with show command",[]string{"--server", serverFlag, "--port", "12312321", "show", "lo"}, 1, connectionError},
		{"wrong port value format with show command",[]string{"--server", serverFlag, "--port", "dsadqw", "show", "lo"}, 0, incorrectUsage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(cliBinary, tt.args...)

			output, err := cmd.CombinedOutput()
			if (tt.exitCode == 0 && err != nil) {
				t.Fatalf("got: %v, expected 0 exit code", err)
			}
			if (tt.exitCode != 0 && err == nil) {
				t.Fatal("got nil error, expected non-zero exit code")
			}

			actual := string(output)

			if (tt.outputPrefix != "" && !strings.HasPrefix(actual, tt.outputPrefix)) {
				t.Fatalf("got: %s, expected prefix: %s", actual, tt.outputPrefix)
			}
		})
	}
}
