package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
)

type CommandCloudStatus struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewCloudStatus() *CommandCloudStatus {
	return &CommandCloudStatus{}
}

func (m *CommandCloudStatus) Start(respChan chan string, promptChan chan string, userChan chan string) {
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getCloudStatus
	cargo := CommandCloudStatus{}
	stateMachine.Start(cargo)
}

func (m *CommandCloudStatus) getCloudStatus(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewCloudStatus()
	args := o.NewGet()
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	m.responseChan <- string(resp)
	return nil
}
