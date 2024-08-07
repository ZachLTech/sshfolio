package main

import (
	"sshfolio/app"
)

// Starts either the TUI SSH session or TUI program depending on if the flag is true or false
var SSHEnabled bool = true

const (
	host = "0.0.0.0"
	port = "23"
)

func main() {
	if SSHEnabled {
		app.RunSSHTUI(host, port)
	} else {
		app.RunTUI()
	}
}
