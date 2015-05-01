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
	repl.InstrMachineMap["cloud status"] = NewMachineCloudStatus().Start
	repl.InstrMachineMap["usage"] = NewMachineUsage().Start
	repl.InstrMachineMap["balance"] = NewMachineBalance().Start
	repl.InstrMachineMap["create server"] = NewMachineCreateServer().Start
	repl.Start()
}
