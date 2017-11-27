package main

const (
	apiError = "API error"
	incorrectUsage = "Incorrect Usage"
	usage = "NAME:"
	incorrectCommand = "No help topic for"
	showWithoutArgument = "Empty argument list"
	version = "main version"
	connectionError = "Connection error"
	serverPath = "../../ifconfig-service/main/"
)

func TestMain(m *testing.M) {
	err := os.Chdir(serverPath)
	if (err != nil) {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}
	
	os.Exit()
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
		{"incorrect command with bad flag", []string{"no such command", "--grw", "trwrefwa", "gfdfdz"}, 1, incorrectUsage},
		{"incorrect flag", []string{"--nosuchflag"}, 0, incorrectUsage},
		{"incorrect flag with garbage", []string{"--nosuchflag", "wqew", "uytuytuj", "jhffgd"}, 0, incorrectUsage},
		{"version", []string{"--version"}, 0, version},
		{"version with garbage", []string{"--version", "wqew", "uytuytuj", "jhffgd"}, 0, version},
		{"version with incorrect flag", []string{"--version", "--wqew", "uytuytuj", "jhffgd"}, 0, incorrectUsage},
		{"list command with incorrect flag", []string{"list", "--wqew", "uytuytuj", "jhffgd"}, 0, incorrectUsage},
		{"list command", []string{"list"}, 0, ""},
		{"server flag with list command", []string{"--server", "localhost", "list"}, 0, ""},
		{"google server flag with list command", []string{"--server", "google.com", "list"}, 1, connectionError},
		{"valid server and port flags with list command", []string{"--server", "localhost", "--port", "55555", "list"}, 0, ""},
		{"valid port flag with list command",[]string{"--port", "55555", "list"}, 0, ""},
		{"missed server flag value with list command",[]string{"--server", "--port", "55555", "list"}, 1, incorrectCommand},
		{"invalid port value with list command",[]string{"--port", "12312321", "list"}, 1, connectionError},
		{"wrong port value format with list command",[]string{"--port", "dsadqw", "list"}, 0, incorrectUsage},
		{"show command", []string{"show","lo"}, 0, ""},
		{"missed show command argument", []string{"show","lo"}, 1, showWithoutArgument},
		{"show command with invalid interface", []string{"show","no such interface"}, 1, apiError},
		{"server flag with show command", []string{"--server", "localhost", "show", "lo"}, 0, ""},
		{"google server flag with show command", []string{"--server", "google.com", "show", "lo"}, 1, connectionError},
		{"valid server and port flags with show command", []string{"--server", "localhost", "--port", "55555", "show", "lo"}, 0, ""},
		{"valid port flag with show command",[]string{"--port", "55555", "show", "lo"}, 0, ""},
		{"missed server flag value with show command",[]string{"--server", "--port", "55555", "show", "lo"}, 1, incorrectCommand},
		{"invalid port value with show command",[]string{"--port", "12312321", "show", "lo"}, 1, connectionError},
		{"wrong port value format with show command",[]string{"--port", "dsadqw", "show", "lo"}, 0, incorrectUsage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
		})
	}
}
