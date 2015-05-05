package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"os"
)

var (
	config *cloudsigma.Config
)

func main() {
	config = cloudsigma.NewConfig()
	_, err := config.Load()
	if err != nil {
		fmt.Println("Unable to load config.", err)
		os.Exit(1)
	}

	// Create the repl, add command state machines, and start the repl.
	repl := replizer.NewRepl()
	repl.Name = "CloudSigma IaaS"
	// Create a statemachine per command available in the repl.
	repl.Add("cloud status", NewCloudStatus().Start)
	repl.Add("usage", NewUsage().Start)
	repl.Add("balance", NewBalance().Start)
	repl.Add("create server", NewCreateServer().Start)
	repl.Add("notification contacts", NewGetNotifyContacts().Start)
	repl.Add("notification preferences", NewGetNotifyPrefs().Start)
	repl.Start()
}
