package main

import (
	"os"
	"sshfolio/app"
	"sshfolio/ui"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	// Loads .env and SSH_SERVER_ENABLED value
	err := godotenv.Load()
	ui.Check(err, "Loading .env in main... Make sure you have your .env file in the root directory of this program", true)
	// Starts either the TUI SSH session or TUI program depending on if the flag is true or false.
	SSHEnabled, err := strconv.ParseBool(os.Getenv("SSH_SERVER_ENABLED"))
	ui.Check(err, "Parsing .env SSH_SERVER_ENABLED bool in main", true)

	// Starts the app :D
	if SSHEnabled {
		app.RunSSHTUI()
	} else {
		app.RunTUI()
	}
}
