package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
)

type CommandUsage struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewMachineUsage() *CommandUsage {
	return &CommandUsage{}
}

func (m *CommandUsage) Start(respChan chan string, promptChan chan string, userChan chan string) {
	fmt.Println("starting cloud status...")
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getUsage
	cargo := CommandUsage{}
	stateMachine.Start(cargo)
}

func (m *CommandUsage) getUsage(cargo interface{}) statemachiner.StateFn {
	fmt.Printf("Final cmd:%+v\n", cargo)
	o := cloudsigma.NewUsage()
	args := o.NewGet()
	fmt.Println("Username:", config.Login().Username)
	args.Username = config.Login().Username
	args.Password = config.Login().Password
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	m.responseChan <- string(resp)
	return nil
}
