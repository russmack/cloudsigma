package main

import (
	"github.com/russmack/replizer"
)

func main() {
	// Create the repl, add command state machines, and start the repl.
	repl := replizer.NewRepl()
	repl.Name = "CloudSigma IaaS"
	// Create a statemachine per command available in the repl.
	repl.InstrMachineMap["cloud status"] = NewMachineCloudStatus().Start
	repl.InstrMachineMap["create server"] = NewMachineCreateServer().Start
	repl.Start()
}
